package main

import (
	"./routers"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	router := fasthttprouter.New()
	routers.InitRoutes(router)
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
