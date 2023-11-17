package user_graph

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"social/utils"
)

func Print(args ...interface{}) {
	fmt.Println(args...)
}

type UserGraph struct {
	UserList map[string][]string `json:"users"`
}

func NewUserGraph() UserGraph {
	f, err := os.OpenFile("user_graph_offline.txt", os.O_RDONLY, 0777)
	if err != nil {
		return UserGraph{
			UserList: make(map[string][]string),
		}
	}
	defer f.Close()

	file_contents, err := io.ReadAll(f)
	if err != nil {
		Print("Unable to load offline user data")
		os.Exit(1)
	}
	var user_graph_offline UserGraph
	err = json.Unmarshal(file_contents, &user_graph_offline)
	if err != nil {
		Print("Unable to parse offline user data")
	}
	return user_graph_offline
}

func (u *UserGraph) userExists(user string) bool {
	for userKey := range u.UserList {
		if userKey == user {
			return true
		}
	}
	return false
}

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
		Print("The requested friend user is not present, Please try another user")
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

func (u *UserGraph) displayOtherUsers(currUser string) {
	Print("Please select another user from the following list")
	for userKey := range u.UserList {
		if userKey != currUser {
			Print(userKey)
		}
	}
}

func (u *UserGraph) ConvertGraphToBytes() []byte {
	bytes, err := json.Marshal(u)
	if err != nil {
		Print("Unable to marhal user list")
		os.Exit(1)
	}
	return bytes
}
