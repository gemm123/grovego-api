package main

import (
	"gemm123/treetrek/routes"
)

func main() {
	router := routes.Routes()
	router.Run()
}
