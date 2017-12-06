package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/str/controllers"
)

func main() {
	r := httprouter.New()
	sc := controllers.NewStrController()
	r.POST("/", sc.Str)
	http.ListenAndServe(":8080", r)
}
