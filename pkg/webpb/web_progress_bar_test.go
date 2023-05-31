package webpb

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/derv-dice/gstk/pkg/webpb/internal/model"
)

var webPb = NewWebProgressBar(":8080", 1000)
var bar, bar2 model.ProgressBar
var stop = make(chan bool)

func TestNewWebProgressBar(t *testing.T) {
	webPb.Run()
	var err error
	bar, err = webPb.AddNewProgressBar("test", 0, 1000)
	assert.NoError(t, err)
	bar2, err = webPb.AddNewProgressBar("test 2", 0, 5000)

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			bar.Inc()
			webPb.AddNewEventf("count %d", bar.Val())
		}
	}()

	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			bar2.Inc()
		}
	}()

	started := make(chan struct{}, 1)
	go func() {
		startHttpServer(started)
	}()
	<-started

	<-stop
}

// startHttpServer - Запуск http сервера, который умеет увеличивать счетчик у webPb
func startHttpServer(started chan struct{}) {
	http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inc")
		bar.Inc()
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("add")
		bar.Add(10)
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		stop <- true
	})

	go func() {
		err := http.ListenAndServe(":7070", nil)
		if err != nil {
			panic(err)
		}
	}()
	started <- struct{}{}
	return
}
