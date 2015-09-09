package bipolarkcombinatorkcombinator

import (
	"log"
	"testing"
)

func TestBKCKC(t *testing.T) {
	for i := 0; i < 10; i++ {
		f := New(i, logCombinator)
		f2 := f()
		f2()
	}
}

func logCombinator(obj interface{}) interface{} {
	log.Printf("CALLED ME: %s", obj)
	return obj
}
