package main

import (
	"github.com/derv-dice/gstk/pkg/webpb"
	"time"
)

const (
	bar1MaxVal = 10
	bar2MaxVal = 100
)

func main() {
	exit := make(chan bool)

	var wpb = webpb.NewWebProgressBar(":8080", 1000)
	wpb.Run()
	defer wpb.Stop()

	bar1, err := wpb.AddNewProgressBar("ProgressBar 1", 0, bar1MaxVal)
	if err != nil {
		panic(err)
	}

	bar2, err := wpb.AddNewProgressBar("ProgressBar 2", 0, bar2MaxVal)
	if err != nil {
		panic(err)
	}

	go func() {
		for bar1.Val() < bar1MaxVal {
			time.Sleep(time.Second)
			bar1.Inc()
			wpb.AddNewEventf("bar 1 incremented [%d -> %d]", bar1.Val()-1, bar1.Val())
		}
	}()

	go func() {
		for bar2.Val() < bar2MaxVal {
			time.Sleep(time.Second)
			bar2.Add(15)
			wpb.AddNewEventf("bar 2 [%d -> %d]", bar2.Val()-1, bar2.Val())
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			if bar1.Val() == bar1MaxVal && bar2.Val() == bar2MaxVal {
				time.Sleep(time.Second)
				exit <- true
			}
		}
	}()

	<-exit
}
