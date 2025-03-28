package main

import "api/src/person/infrastructure"

func main() {
    server := infrastructure.NewServer()
    server.Run(":8080")
}