package webpb

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/derv-dice/gstk/pkg/webpb/internal/model"
)

const debugMode = false

var ErrWebProgressBarNotInited = errors.New("WebProgressBar not initialized. It's important to use NewWebProgressBar()")

type WebProgressBar struct {
	mu     sync.Mutex
	server http.Server
	inited bool

	progressBars      map[string]model.ProgressBar
	progressBarsOrder []string

	eventLog model.EventLog
}

func NewWebProgressBar(addr string, eventLogLen int) *WebProgressBar {
	wpb := &WebProgressBar{
		progressBars:      map[string]model.ProgressBar{},
		progressBarsOrder: []string{},
		eventLog:          model.NewEventLog(eventLogLen),
	}

	handler := mux.NewRouter()
	handler.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		wpb.events(w, r)
	}).Methods(http.MethodGet)
	handler.HandleFunc("/updates", func(w http.ResponseWriter, r *http.Request) {
		wpb.render(w, r)
	}).Methods(http.MethodGet)
	handler.HandleFunc("/ui", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.New("ui").Parse(string(wpb.ui()))
		_ = tmpl.Execute(w, map[string]interface{}{"addr": addr})
	})

	if debugMode {
		handler.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	}

	wpb.server = http.Server{Addr: addr, Handler: handler}
	wpb.inited = true
	return wpb
}

func (b *WebProgressBar) Run() *WebProgressBar {
	if !b.isInited() {
		return nil
	}

	go func() {
		if err := b.server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	fmt.Printf("ProgressBar UI: http://localhost%s/ui\n", b.server.Addr)

	if debugMode {
		fmt.Printf("pprof: http://localhost%s/debug/pprof/\n", b.server.Addr)
	}

	return b
}

func (b *WebProgressBar) Stop() {
	_ = b.server.Close()
}

func (b *WebProgressBar) AddNewProgressBar(name string, val, max int) (bar model.ProgressBar, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.isInited() {
		err = ErrWebProgressBarNotInited
		return
	}

	if b.progressBars[name] != nil {
		err = fmt.Errorf("ProgressBar '%s' already exists", name)
		return
	}

	bar = model.NewProgressBar(val, max)
	b.progressBars[name] = bar
	b.progressBarsOrder = append(b.progressBarsOrder, name)
	return
}

func (b *WebProgressBar) GetProgressBarByName(name string) (bar model.ProgressBar, err error) {
	if b.progressBars[name] == nil {
		err = fmt.Errorf("ProgressBar with name '%s' not exists", name)
		return
	}

	bar = b.progressBars[name]
	return
}

func (b *WebProgressBar) AddNewEvent(event string) {
	b.eventLog.Push(event)
}

func (b *WebProgressBar) AddNewEventf(tmpl string, args ...any) {
	b.eventLog.Push(fmt.Sprintf(tmpl, args...))
}

func (b *WebProgressBar) isInited() bool {
	if b.progressBars == nil || b.eventLog == nil {
		return false
	}

	return true
}

func (b *WebProgressBar) events(w http.ResponseWriter, _ *http.Request) {
	view := new(View)
	view.EventLog = b.eventLog.Events()
	data, _ := json.Marshal(view)
	_, _ = w.Write(data)
}

func (b *WebProgressBar) render(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }, HandshakeTimeout: time.Millisecond * 100}

	var conn *websocket.Conn
	var err error
	if conn, err = u.Upgrade(w, r, nil); err != nil {
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	go func() {
		for {
			_, _, cErr := conn.ReadMessage()
			if cErr != nil {
				_ = conn.Close()
				return
			}
		}
	}()

	for {
		time.Sleep(time.Millisecond * 16) // Чуть больше 60Гц (~62,5)
		if err = b.sendJSON(conn); err != nil {
			return
		}
	}
}

func (b *WebProgressBar) sendJSON(conn *websocket.Conn) (err error) {
	view := new(View)

	var isUpdated bool
	isUpdated, view.EventLogUpdates = b.eventLog.IsUpdated()

	for k := range b.progressBars {
		if b.progressBars[k].IsUpdated() {
			isUpdated = true
		}
	}

	if !isUpdated {
		return
	}

	for i := range b.progressBarsOrder {
		view.ProgressBars = append(view.ProgressBars, &ProgressBarView{
			Name: b.progressBarsOrder[i],
			Val:  b.progressBars[b.progressBarsOrder[i]].Val(),
			Max:  b.progressBars[b.progressBarsOrder[i]].Len(),
		})
	}

	return conn.WriteJSON(view)
}

func (b *WebProgressBar) ui() []byte {
	return _uiTmpl
}

type View struct {
	ProgressBars    []*ProgressBarView `json:"progressBars"`
	EventLog        []string           `json:"eventLog,omitempty"`
	EventLogUpdates []string           `json:"eventLogUpdates,omitempty"`
}

type ProgressBarView struct {
	Name string `json:"Name"`
	Val  int    `json:"Val"`
	Max  int    `json:"Max"`
}
