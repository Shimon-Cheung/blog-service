package main

import "github.com/Shimon-Cheung/blog-service/internal/routers"

func main() {
	router := routers.NewRouter()
	router.Run()
}
