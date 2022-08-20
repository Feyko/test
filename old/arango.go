package main

import (
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
)

type ayo struct {
}

type aya ayo

func main() {
	conn, _ := http.NewConnection(http.ConnectionConfig{Endpoints: []string{"http://localhost:8529"}})
	client, _ := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("user", ""),
	})
	db, _ := client.Database(nil, "db")
	exists, err := db.CollectionExists(nil, "coll")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exists)
}
