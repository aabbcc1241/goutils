package lang

type Empty bool
type Semaphore chan Empty

func (s Semaphore) P(n int) {
	e := Empty(true)
	for i := 0; i < n; i++ {
		s <- e
	}
}
func (s Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
func (s Semaphore) Lock() {
	s.P(1)
}
func (s Semaphore) Unlock() {
	s.V(1)
}
func (s Semaphore) Signal() {
	s.V(1)
}
func (s Semaphore) Wait(n int) {
	s.P(n)
}
