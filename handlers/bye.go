package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (h *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Bye World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Bye %s\n", d)
}
