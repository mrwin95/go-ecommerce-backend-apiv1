package main

import "github.com/mrwin95/go-ecommerce-backend-api/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run(":8080")
}
