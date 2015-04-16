package examples

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

var (
	// Retrieves all the information about the user with ID 1
	// the second parameter, an error, is suppressed
	user, _ = nerdz.NewUser(1)
)

// prints all the friends information
func findFriends() {
	// user.Friends() returns a pointer to an array whose elements are User
	if friendsList := user.Friends(); friendsList != nil {
		fmt.Println("#### Friends ######")
		// Dereference the pointer
		for _, otherUser := range *friendsList {
			fmt.Printf("%+v", otherUser)
		}

		fmt.Println("##################")
	} else {
		fmt.Printf("User(%d) hasn't any friends", user.Counter)
	}

}
