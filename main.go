package main

import (
	"gemm123/grovego-api/routes"
)

func main() {
	router := routes.Routes()
	router.Run()
}
