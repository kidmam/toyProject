package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f Foo) TableName() string {
	return "foo"
}

func main() {
	f := &Foo{
		ID:        1,
		FirstName: "Foo",
		LastName:  "Bar",
		Age:       30,
	}

	fmt.Println(funk.Contains([]string{"foo", "bar"}, "bar")) // true
	funk.Contains([]*Foo{f}, f)                               // true
	funk.Contains([]*Foo{f}, nil)                             // false
}
