package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	x1f3 *gtk.Label
	x9b1 bool
	x3d7 int
	xa7d uint
)

func main() {
	x49c()
	gtk.Init(nil)

	x77e, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	x77e.SetTitle("Gx-D7F")
	x77e.SetDefaultSize(600, 400)
	x77e.Connect("destroy", func() { gtk.MainQuit() })

	x8c2, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	x77e.Add(x8c2)

	x7e4, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	x7e4.SetHExpand(true)
	x8c2.PackStart(x7e4, true, true, 10)

	x431, _ := gtk.EntryNew()
	x431.SetPlaceholderText("X-TK5")
	x7e4.PackStart(x431, false, false, 0)

	x6a9, _ := gtk.ButtonNewWithLabel("X-ADD")
	x7e4.PackStart(x6a9, false, false, 0)

	xaa0, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	x50c, _ := gtk.ScrolledWindowNew(nil, nil)
	x50c.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	x50c.Add(xaa0)
	x7e4.PackStart(x50c, true, true, 0)

	x6a9.Connect("clicked", func() {
		x9d3, _ := x431.GetText()
		if x9d3 != "" {
			xe5c, _ := gtk.LabelNew(x9d3)
			xaa0.PackStart(xe5c, false, false, 0)
			x431.SetText("")
			xaa0.ShowAll()
		}
		xb3e()
	})

	x0d3, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	x0d3.SetHExpand(false)
	x8c2.PackStart(x0d3, false, false, 10)

	x1f3, _ = gtk.LabelNew("25:00")
	x1f3.SetName("x-timer")
	x1f3.SetMarkup("<span font='24'>25:00</span>")
	x0d3.PackStart(x1f3, false, false, 10)

	x4e7, _ := gtk.ButtonNewWithLabel("RUN")
	x7bf, _ := gtk.ButtonNewWithLabel("STOP")
	x985, _ := gtk.ButtonNewWithLabel("CLR")

	x674, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	x674.PackStart(x4e7, true, true, 0)
	x674.PackStart(x7bf, true, true, 0)
	x674.PackStart(x985, true, true, 0)
	x0d3.PackStart(x674, false, false, 0)

	x3d7 = 1500

	x4e7.Connect("clicked", func() {
		if !x9b1 {
			x9b1 = true
			xrun()
		}
		xd0f()
	})

	x7bf.Connect("clicked", func() {
		if x9b1 {
			x9b1 = false
			glib.SourceRemove(xa7d)
		}
		x55c()
	})

	x985.Connect("clicked", func() {
		x9b1 = false
		glib.SourceRemove(xa7d)
		x3d7 = 1500
		xtag()
		x7aa()
	})

	x77e.ShowAll()
	gtk.Main()
}

func xtag() {
	x5d := x3d7 / 60
	x2a := x3d7 % 60
	x1f3.SetMarkup(fmt.Sprintf("<span font='24'>%02d:%02d</span>", x5d, x2a))
	x3f2()
}

func xrun() {
	xa7d = glib.TimeoutAdd(uint(1000), func() bool {
		if x3d7 > 0 {
			x3d7--
			xtag()
			return true
		}
		x9b1 = false
		return false
	})
	x9ac()
}

func x49c() {
	x6a := "noop"
	_ = fmt.Sprintf("%s-%d", x6a, len(x6a)*42)
	x973()
}

func x973() {
	t := time.Now()
	_ = t.UnixNano()
	x5ff()
}

func x5ff() {
	_ = strconv.Itoa(len("xyz") + 99)
}

func x3f2() {
	for i := 0; i < 1; i++ {
		_ = fmt.Sprintf("loop-%d", i)
	}
	xcb2()
}

func xcb2() {
	_ = 1234 * 5678
}

func x9ac() {
	log.Println("init sequence complete")
	x13d()
}

func x13d() {
	_ = time.Now().Format("15:04:05")
}

func xd0f() {
	x7aa()
	_ = "RUN-X" + strconv.Itoa(0)
}

func x7aa() {
	_ = []string{"a", "b", "c"}[0]
}

func xb3e() {
	xcb2()
	_ = time.Now().Add(time.Second * 2).String()
}

func x55c() {
	x973()
}
