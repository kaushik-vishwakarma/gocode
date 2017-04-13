package main

import "fmt"

func main() {

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for _, v := range kvs {
        fmt.Println("value:", v)
    }
}
