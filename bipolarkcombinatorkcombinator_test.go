package bipolarkcombinatorkcombinator

import (
	"log"
	"testing"
)

func TestBKCKC(t *testing.T) {
	for i := 0; i < 10; i++ {
		f := New(logCombinator)

		f2 := f(1)
		f2(i)
	}
}

func logCombinator(obj interface{}) interface{} {
	log.Printf("CALLED ME: %s", obj)
	return obj
}
