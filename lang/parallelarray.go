package lang

import (
	"math/rand"
	"runtime"
	"sync"
)

type ParallelArray struct {
	Data    []interface{}
	Lock    sync.Mutex
	NThread int
}
type consumer interface {
	Apply(k int, v interface{}, r *rand.Rand)
}
type producer interface {
	Apply(k int, v interface{}, r *rand.Rand) interface{}
}

func (p ParallelArray) Len() int {
	return len(p.Data)
}

/* [start,end) : end is excluded
 * REMARK : this function does not handle lock
 */
func _for(p ParallelArray, f consumer, withRandom bool, start, end int) {
	N := end - start
	n := runtime.GOMAXPROCS(0)
	if p.NThread > 0 {
		n = p.NThread
	}
	if N < n {
		sem := make(Semaphore, N)
		for ; start < end; start++ {
			go func(i int) {
				var r *rand.Rand
				if withRandom {
					r = rand.New(rand.NewSource(int64(i)))
				}
				f.Apply(i, p.Data[i], r)
				sem.Signal()
			}(start)
		}
		sem.Wait(N)
	} else {
		s := N / n
		sem := make(Semaphore, n)
		for i := 0; i < n; i++ {
			go func(i int) {
				var r *rand.Rand
				if withRandom {
					r = rand.New(rand.NewSource(int64(i)))
				}
				for j := i*s + start; j < (i+1)*s+start; j++ {
					f.Apply(j, p.Data[j], r)
				}
				sem.Signal()
			}(i)
		}
		sem.Wait(n)
		if n*s != N {
			_for(p, f, withRandom, n*s+start, end)
		}
	}
}
func _update(p *ParallelArray, f producer, withRandom bool, start, end int) {
	N := end - start
	n := runtime.GOMAXPROCS(0)
	if p.NThread > 0 {
		n = p.NThread
	}
	if N < n {
		sem := make(Semaphore, N)
		for ; start < end; start++ {
			go func(i int) {
				var r *rand.Rand
				if withRandom {
					r = rand.New(rand.NewSource(int64(i)))
				}
				p.Data[i] = f.Apply(i, p.Data[i], r)
				sem.Signal()
			}(start)
		}
		sem.Wait(N)
	} else {
		s := N / n
		sem := make(Semaphore, n)
		for i := 0; i < n; i++ {
			go func(i int) {
				var r *rand.Rand
				if withRandom {
					r = rand.New(rand.NewSource(int64(i)))
				}
				for j := i*s + start; j < (i+1)*s+start; j++ {
					p.Data[j] = f.Apply(j, p.Data[j], r)
				}
				sem.Signal()
			}(i)
		}
		sem.Wait(n)
		if n*s != N {
			_update(p, f, withRandom, n*s+start, end)
		}
	}
}
func (p ParallelArray) For(f consumer, withRandom bool) {
	p.Lock.Lock()
	_for(p, f, withRandom, 0, len(p.Data))
	p.Lock.Unlock()
}

/* in-place map and replace */
func (p *ParallelArray) Update(f producer, withRandom bool) {
	p.Lock.Lock()
	_update(p, f, withRandom, 0, len(p.Data))
	p.Lock.Unlock()
}

//TODO map, reduce, fold
