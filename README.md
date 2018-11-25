# Golang-TDD - Quotes from 
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces

## benchmarks
It's similar to writing tests but it gives more details. Instead of using *testing.T, benchmark uses *testing.B. Naming convention is Benchmark<Name_of_Function>. 

```go test -bench=.```

## Arrays and slices...
Slices do not encode the size of the collection

## Structs and Interfacse
This is quite different to interfaces in most other programming languages. Normally you have to write code to say My type Foo implements interface Bar.

But in our case

* Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface

* Circle has a method called Area that returns a float64 so it satisfies the Shape interface

* string does not have such a method, so it doesn't satisfy the interface

> In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

## [Table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) 
​Table driven tests are useful when you want to build a list of test cases that can be tested in the same manner.

