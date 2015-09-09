// Package bipolarkcombinatorkcombinator is a package that exposes a BipolarKcombinatorKcombinator
// This is a function which when called returns a function that you can call which will
// either return the original object after the application of a user provided function,
// or the unmodified copy of the same object.
package bipolarkcombinatorkcombinator

import "math/rand"

type kCombinator func(interface{}) interface{}

// New is the way to create a Bipolar KCombinator KCombinator
func New(kComb kCombinator) func(interface{}) kCombinator {
	return func(obj interface{}) kCombinator {
		randy := rand.Int() % 2
		if randy > 0 {
			return kComb
		}
		return func(obj interface{}) interface{} {
			return obj
		}
	}
}
