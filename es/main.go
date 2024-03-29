package main

import (
    "context"
    "fmt"

    "github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Married bool   `json:"married"`
}

func main() {
    client, err := elastic.NewClient(elastic.SetURL("https://elk-offline.wsecar.cn"), elastic.SetBasicAuth("wangzhipeng", "wangzhipeng"))
    if err != nil {
        // Handle error
        panic(err)
    }

    fmt.Println("connect to es success")
    p1 := Person{Name: "lmh", Age: 18, Married: false}
    put1, err := client.Index().
        Index("user").
        BodyJson(p1).
        Do(context.Background())
    if err != nil {
        // Handle error
        panic(err)
    }
    fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
