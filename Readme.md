# golang-test

- Here I am documenting my first learnings from golang.
- I am using gitbook.io's "learn go with tests" here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world



## Learnings
- test.sh shows command line interface to the go executer
- discipline of code cycle is as follows:
    - write a test under xxx_test.go
    - make the compiler pass
    - run the test and see if it fails w/ meaningful messages
    - write enough code to make the test pass
    - refactor