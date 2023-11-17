package user_graph

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
