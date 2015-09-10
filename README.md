BiPolar KCombinator KCombinator
===============================

Honestly? I won't even insult your intelligence by thinking anyone would need
this term or it's applications explained to them

Usage
=====

Create a function that takes an interface, and returns an interface

```go
func myKComb(obj interface{}) interface{} {
  // modify obj to your hearts content
}
```

Create an object of some kind that you care about maybe

```go
myObj := mypackage.MyStruct{}
```

Create a new BipolarKcombinatorKcombinator

```go
f := bipolarkcombinatorkcombinator.New(myObj, myKComb)
```

Invoke the resulting function to receive your KCombinator

```go
myK := f()
```

Try your luck, and see if you get back the original object, or the modified result
of the original object! It's a mystery, every time!

```go
resultObj := myK()
```

...Why?
=======

I was joking around about how HyperLogLog was almost a meaningless name, since it
doesn't do a great job telegraphing what it actually is (to me, anyways). I then
went and made up several joke data structures with absurd names, and then finally
made up this one.

I then thought to myself "You know... I could probably write that"

License
=======
Apache v2 - See LICENSE
