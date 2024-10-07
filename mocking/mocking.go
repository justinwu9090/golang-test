package main

const finalWord = "Go!"
const countdownStart = 3

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)

}

func main() {
	Countdown(os.Stdout)
}
