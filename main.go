package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
)

var (
	redisHost      = flag.String("rh", "127.0.0.1", "Redis Hostname")
	redisPort      = flag.String("rp", "6379", "Redis Port")
	maxConnections = flag.Int("mc", 1000, "Max connections to Redis")
)

var redisPool = redis.NewPool(func() (redis.Conn, error) {
	fmt.Println("Connect to redis://" + *redisHost + ":" + *redisPort + "...")

	c, err := redis.Dial("tcp", *redisHost+":"+*redisPort)

	return c, err
}, *maxConnections)

var rc = redisPool.Get()

var categories = [10]string{
	"top",
	"dom",
	"int",
	"eco",
	"ent",
	"spo",
	"52",
	"gourmet",
	"love",
	"trend",
}

func main() {
	flag.Parse()

	router := NewRouter()
	fmt.Println("Init news...")

	for _, c := range categories {
		fmt.Println(c)
	}

	fmt.Println("Binding localhots:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
