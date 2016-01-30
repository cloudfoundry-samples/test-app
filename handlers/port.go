package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

type Port struct {
}

func (_ *Port) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ports := strings.Split(r.Host,":")
	if len(ports) < 2 {
		w.Write([]byte(fmt.Sprintf("0")))
		return
	}

	w.Write([]byte(fmt.Sprintf("%s", ports[1])))
}
