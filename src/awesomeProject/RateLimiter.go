package main

import (
	"fmt"
	"time"
)



func main() {
	print("starting rate limiter")
	bucket := make(chan time.Time, 10)
	//token  := make(chan int, 5)
	//client send to bucket
	go infinitReqFromBucketToServer(bucket)
	//close (bucket)
	limitRate(bucket, 1)

}

func limitRate(bucket chan time.Time, rateLimiter int) {

	tickStream := time.Tick(time.Duration(1000/rateLimiter) * time.Millisecond)
	println(cap(bucket))

	for i:=0; i<cap(bucket); i++{
		bucket <- time.Now()
	}
	for token := range tickStream {
		bucket <- token
		//bucket <- fmt.Sprintf("adding token to bucket: %s", token)
		fmt.Printf("\nlimiter %s", token)
		}

}

func infinitReqFromBucketToServer(bucket chan time.Time) {
	i := 0
	for {//infinite loop
		<- bucket
		fmt.Printf("\n req #%d\n", i)
		time.Sleep(time.Millisecond * 200)
		i++
	}

}

/*
go build
go get
go test
 */