package factorymethod

import (
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	db1 := factorymethod.DatabaseFactory("production")
	db1.putData("test", "this is mongodb")
	fmt.Println(db1.getData("test"))
	db2 := factorymethod.DatabaseFactory("development")
	db2.putData("test", "this is sqlite")
	fmt.Println(db2.getData("test"))
}
