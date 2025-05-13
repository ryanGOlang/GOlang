package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
)

func x1(y1 int) {
	z1, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	z1.SetTitle(fmt.Sprintf("W #%d", y1))
	z1.SetDefaultSize(200, 100)
	a1, _ := gtk.LabelNew(fmt.Sprintf("Hello from W %d!", y1))
	z1.Add(a1)
	z1.ShowAll()
	z1.Connect("destroy", func() {
		z1.Destroy()
	})
}

func b1() {
	for i := 0; i < 10; i++ {
		_ = rand.Intn(100)
	}
}

func c1() {
	for j := 0; j < 5; j++ {
		fmt.Println("Hello World:", j)
	}
}

func d1() {
	time.Sleep(1 * time.Second)
}

func e1() int {
	return rand.Intn(50) + 1
}

func f1() {
	for k := 0; k < 3; k++ {
		fmt.Println("Hello World:", e1())
		d1()
	}
}

func main() {
	runtime.LockOSThread()
	gtk.Init(nil)
	e2 := 0
	glib.TimeoutAdd(uint(500), func() bool {
		e2++
		go x1(e2)
		b1()
		c1()
		f1()
		return true
	})
	gtk.Main()
}
