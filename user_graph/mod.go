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

/*
* USER GRAPH CONSTRUCTOR -
* This constructor reads the offline user data first. If the user data is not present
* It will create a new user graph and return that to main.go file
 */
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

/*
--------------------------- INTERNAL FUNCTIONS ---------------------------
*/
// This function is used to display other users present in the graph to stdout
func (u *UserGraph) displayOtherUsers(currUser string) {
	filterList := append(u.UserList[currUser], currUser)
	suggestionList := utils.FilterUsers(u.UserList, filterList)
	if len(suggestionList) != 0 {
		Print("Please select another user from the following list")
		for suggestion := range suggestionList {
			Print(suggestion)
		}
	}
}

func (u *UserGraph) userExists(user string) bool {
	for userKey := range u.UserList {
		if userKey == user {
			return true
		}
	}
	return false
}

/*
--------------------------- UTILITY FUNCTIONS ---------------------------
*/
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
