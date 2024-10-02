go mod init hello # create a new file called go.mod. seems to take every .go file that is not a test file and then specify that it exists for running

go test # runs testing. I guess any *_test.go will be ran with 

go test -bench=. # runs with benchmark what will pick an iteration for b.N