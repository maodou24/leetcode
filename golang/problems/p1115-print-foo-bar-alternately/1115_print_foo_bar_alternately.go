package printfoobaralternately
/*
 * @lc app=leetcode id=1115 lang=golang
 *
 * [1115] Print FooBar Alternately
 */

// @lc code=start
type FooBar struct {
	n int

    barChan chan int
    fooChan chan int
}

func NewFooBar(n int) *FooBar {
	f := &FooBar{n: n}
    f.barChan = make(chan int, 1)
    f.fooChan = make(chan int, 1)

    f.fooChan <- 1
    return f
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
        <- fb.fooChan
		// printFoo() outputs "foo". Do not change or remove this line.
        printFoo()
        fb.barChan <- 1
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
        <- fb.barChan
		// printBar() outputs "bar". Do not change or remove this line.
        printBar()
        fb.fooChan <- 1
	}
}
// @lc code=end

