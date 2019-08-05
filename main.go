package main

	
import (
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/go-redis/redis"
)

func Hello(ctx *fasthttp.RequestCtx) {
	name := (ctx.UserValue("name")).(string)
	increaseCount()
	fmt.Fprintf(ctx, "hello, %s!\n", name)
}

func View(ctx *fasthttp.RequestCtx){
	redisClient := createClient()
	defer redisClient.Close()

	count, err := redisClient.Get("count").Result()
	if err != nil {
		count = "0"
		fmt.Println(err)
	}
	fmt.Fprintf(ctx, "view count: %s!\n", count)
}

func main() {
    router := fasthttprouter.New()
	router.GET("/hello/:name", Hello)
	router.GET("/view", View)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

func createClient() *redis.Client {
	return redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
	})
}

func increaseCount()  {
	redisClient := createClient()
	defer redisClient.Close()

	newCount := redisClient.Incr("count")
	if newCount != nil{
		fmt.Println(newCount)
	}
}