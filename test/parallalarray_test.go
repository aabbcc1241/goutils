package test

import (
	. "github.com/beenotung/goutils/lang"
	"math/rand"
	"testing"
	"time"
)

/* 1 thread  : 10.002 s
 * 8 threads :  2.002 s
 */

var c = 0

type for_s struct {
}

func (p for_s) Apply(k int, v Empty, r *rand.Rand) {
	c++
	time.Sleep(1 * time.Second)
	//t.Log(k)
}
func TestParallelArray(t *testing.T) {
	arr := ParallelArray{Data: make([]Empty, 10), NThread: 8}
	t.Log("array len:", len(arr.Data))
	//t.Log("start parallel for loop")
	arr.For(for_s{}, false)
	t.Log("non-collision", c)
	t.Log("collision", len(arr.Data)-c)
	//t.Log("finish parallel for loop")
	if c == len(arr.Data) {
		t.Fail()
	}
}
