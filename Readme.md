# golang-test

- Here I am documenting my first learnings from golang.
- I am using gitbook.io's "learn go with tests" here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world



## Learnings
- run.sh shows command line interface to the go executer
- discipline of code cycle is as follows:
    - write a test under xxx_test.go
    - make the compiler pass
    - run the test and see if it fails w/ meaningful messages
    - write enough code to make the test pass
    - refactor
- xxx_test.go has keywords for:
    - func TestXXX() will make tests that run
    - func ExampleXXX() {//Output: X} will create autodocced examples that will also run if output comment is provided
    - func BenchmarkXXX allows you to run b.N times and measures how long it takes...

- Advantages of go:
    - type-safety at compile-time
    - compiles to assembly w/out a JIT like JVM
    - simple yet robust test framework
    - auto documentation
-  other advantages:
    - Go has faster cold start benchmarks for serverless functions.
    - Go is compiled and type safe
    - Go has a testing framework built into the SDK
    - JS/TS has better convenience built into the SDK for Arrays and Maps, e.g arrays can behave like streams
    - JS/TS enables programmers to dip into both sides of the stack