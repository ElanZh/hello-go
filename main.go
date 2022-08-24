package main

import "hello-go/router"

func main() {
	r := router.Router

	router.SetRouter()

	err := r.Run(":1234")
	if err != nil {
		return
	}
}
