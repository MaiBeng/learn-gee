package main

import (
	"fmt"
	"net/http"
	"gee/gee"
)

func main() {
	gee := gee.New()
	gee.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("1111")
	})
	gee.Run(":9999")
}
