package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	u := prototype.NewUser("email@gmail.com", "9876543")
	uu := u.WithEmail("email@hotmail.com").WithPhone("999999")  fmt.Println("Original User Object ", u)
	fmt.Println("User object copied during runtime ", uu)
}
