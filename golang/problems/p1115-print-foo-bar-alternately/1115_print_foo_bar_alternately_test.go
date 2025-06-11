package printfoobaralternately

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFooBar(t *testing.T) {
	n := NewFooBar(100)

	ch := make(chan string, 20)

	var fooPrinter = func() {
		ch <- "foo"
	}

	var barPrinter = func() {
		ch <- "bar"
	}

	go func() {
		var prev string
		for v := range ch {
			if prev == "" || prev == "bar" {
				assert.Equal(t, "foo", v, "expected foo, got %s", v)
			} else if prev == "foo" {
				assert.Equal(t, "bar", v, "expected bar, got %s", v)
			}
			prev = v
            fmt.Print(v)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		n.Bar(barPrinter)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		n.Foo(fooPrinter)
	}()

	wg.Wait()

}
