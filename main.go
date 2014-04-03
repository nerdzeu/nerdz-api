package main

import (
	"fmt"
	. "github.com/nerdzeu/nerdz-api/orm"
)

func main() {
	var user User
	err := user.New(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	info := user.GetInfo()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", info.Gravatar)
}
