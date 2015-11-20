package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

type Port struct {
}

func (_ *Port) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	port := strings.Split(r.Host,":")[1]

	w.Write([]byte(fmt.Sprintf("%s", port)))
}
