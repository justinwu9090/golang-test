package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'", expected, sum)
	}
}

// example is auto generated in godoc and also ran with `go test -v`
// special comment Output: 6 will mean the example is actually executed
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
