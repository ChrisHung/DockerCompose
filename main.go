package main

	
import (
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Hello(ctx *fasthttp.RequestCtx) {
	name := (ctx.UserValue("name")).(string)
	fmt.Fprintf(ctx, "hello, %s!\n", name)
}

func main() {
    router := fasthttprouter.New()
	router.GET("/hello/:name", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}