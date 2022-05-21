package main

import (
	"fmt"
)

type UsernameFound struct {
	userName string
}

func (e *UsernameFound) Error() string {
	return fmt.Sprintf("User not found %v", e.userName)
}

func myFunc() error {
	ok := false
	if ok != false {
		return nil
	}

	return &UsernameFound{"Mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

}
