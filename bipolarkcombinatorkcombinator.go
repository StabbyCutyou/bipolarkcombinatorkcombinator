// Package bipolarkcombinatorkcombinator is a package that exposes a BipolarKcombinatorKcombinator
// This is a function which when called returns a function that you can call which will
// either return the original object after the application of a user provided function,
// or the unmodified copy of the same object.
package bipolarkcombinatorkcombinator

import (
	"crypto/rand"
	"math/big"
)

// Version is the current version of the library, in semver
const Version = "0.0.1"

// KCombinator represents a function which returns an object, considered "self"
type KCombinator func() interface{}

// KCombinatorKCombinator represents a function which returns a function of type KCombinator
type KCombinatorKCombinator func() KCombinator

// Take in an object O, and a function that accepts and returns a value of type object - this is K0
// Return a function that when called, returns another function - this is K1
// K1 is a function that, when called, will randomly choose to return either K0 or K2
// K2 is a function that returns O un modified
// K0 is the user supplied function that will, presumably, modify 0 and return it

// New is the way to create a Bipolar KCombinator KCombinator
func New(self interface{}, combK func(interface{}) interface{}) KCombinatorKCombinator {
	return func() KCombinator {
		randy, _ := rand.Int(rand.Reader, big.NewInt(2))
		if randy.Int64() > 0 {
			return func() interface{} {
				return combK(self)
			}
		}
		return func() interface{} {
			return self
		}
	}
}
