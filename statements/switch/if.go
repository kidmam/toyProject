package main

import "fmt"

func main() {
    flavors := []string{"chocolate", "vanilla", "strawberry", "banana"}

    for _, flav := range flavors {
        if flav == "strawberry" {
            fmt.Println(flav, "is my favorite!")
            continue
        }

        if flav == "vanilla" {
            fmt.Println(flav, "is great!")
            continue
        }

        if flav == "chocolate" {
            fmt.Println(flav, "is great!")
            continue
        }

        fmt.Println("I've never tried", flav, "before")
    }
}