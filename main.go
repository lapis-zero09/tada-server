package main

import "github.com/lapis-zero09/tada-server/api/router"

func main() {
	r := router.NewRouter()
	r.Logger.Fatal(r.Start(":8080"))
}
