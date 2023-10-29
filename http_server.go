package main

import (
	"fmt"
	"net/http"
)

func setHttpServer(port int) {
	http.Handle("/", http.FileServer(http.Dir("./template")))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
