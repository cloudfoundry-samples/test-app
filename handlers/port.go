package handlers

import (
	"fmt"
	"net/http"
)

type Port struct {
	Ports []string
}

func (p *Port) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%s", p.Ports)))
}
