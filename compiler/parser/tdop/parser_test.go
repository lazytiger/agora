package tdop

import (
	"testing"

	"github.com/PuerkitoBio/goblin/compiler/scanner"
)

var (
	cases = []struct {
		src []byte
	}{
		0: {
			src: []byte(`return 5`),
		},
		1: {
			src: []byte(`aB := 5
return aB`),
		},
		2: {
			src: []byte(`
a := 7
b := 10
add := a + b
sub := a - b
mul := a * b
div := a / b // TODO : Should div return a float even for Int?
mod := b % a
not := !a
unm := -a
`),
		},
		3: {
			src: []byte(`
func Add(x, y) { // Essentially means var Add = func ...
  return x + y
}
return Add(4, "198")
`),
		},
		4: {
			src: []byte(`
Add := func(x, y) { // Essentially means var Add = func ...
  return x + y
}
return Add(4, "198")
`),
		},
		5: {
			src: []byte(`
func Fib(n) {
  if n < 2 {
    return 1
  }
  return Fib(n-1) + Fib(n-2)
}
return Fib(30)
`),
		},
		6: {
			src: []byte(`
import "fmt" // implicit fmt variable
fmt.Println("Hello ", "world")
`),
		},
	}

	isolateCase = 6
)

func TestParse(t *testing.T) {
	Scanner = new(scanner.Scanner)
	for i, c := range cases {
		if isolateCase >= 0 && i != isolateCase {
			continue
		}

		Parse("test", c.src)
	}
}