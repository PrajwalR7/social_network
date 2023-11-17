package user_graph

import (
	"encoding/json"
	"os"
	"social/utils"
)

func (u *UserGraph) AddNewUser(inputUser string) {
	if userPresent := u.userExists(inputUser); !userPresent {
		u.UserList[inputUser] = make([]string, 0)
	}
}

func (u *UserGraph) AddFriend(currUser string, friend string) {
	if currUserExists := u.userExists(currUser); !currUserExists {
		Print("You are not present in our records, you need to first create a new user-account")
		os.Exit(0)
	}
	if friendExists := u.userExists(friend); !friendExists {
		Print("The requested friend is not present.")
		u.displayOtherUsers(currUser)
	} else {
		u.UserList[currUser] = utils.Add(u.UserList[currUser], friend)
		u.UserList[friend] = utils.Add(u.UserList[friend], currUser)
	}
}

func (u *UserGraph) DisplayFriends(currUser string) []string {
	if currUserExists := u.userExists(currUser); !currUserExists {
		Print("You are not present in our records, re-type your username")
		os.Exit(0)
	}

	return u.UserList[currUser]

}

// This utility is used to convert the user graph into bytes, required to store the user data offline in a file
func (u *UserGraph) ConvertGraphToBytes() []byte {
	bytes, err := json.Marshal(u)
	if err != nil {
		Print("Unable to marhal user list")
		os.Exit(1)
	}
	return bytes
}
