package test

import (
	. "github.com/aabbcc1241/goutils/lang"
	"math/rand"
	"testing"
)

var c = 0

type for_s struct {
}

func (p for_s) Apply(k int, v Empty, r *rand.Rand) {
	c++
	//t.Log(k)
}
func TestParallelArray(t *testing.T) {
	arr := ParallelArray{Data: make([]Empty, 1048576), NThread: 8}
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
