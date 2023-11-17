package main

import (
	"fmt"
	"os"
	"social/user_graph"
)

func Print(args ...interface{}) {
	fmt.Println(args...)
}

func Validate(args ...interface{}) {
	for _, arg := range args {
		if arg == nil {
			fmt.Println("Invalid arguments provided please check your arguments again")
			os.Exit(0)
		}
	}
}

func main() {
	args := os.Args[1:]

	command := args[0]

	user_graph := user_graph.NewUserGraph()

	switch command {
	case "add_friend":
		{
			user := args[1]
			friend := args[2]

			Validate(user, friend)
			user_graph.AddFriend(user, friend)
		}
	case "new_user":
		{
			user := args[1]
			Validate(user)

			user_graph.AddNewUser(user)
		}
	case "display_friends":
		{
			user := args[1]
			Validate(user)

			friendList := user_graph.DisplayFriends(user)

			Print("Your friends are:")
			for _, friend := range friendList {
				Print(friend)
			}
		}
	}

	f, err := os.OpenFile("user_graph_offline.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		Print("Unable to store the curr user graph offline")
		os.Exit(1)
	}
	defer f.Close()

	graph_in_bytes := user_graph.ConvertGraphToBytes()
	_, err = f.Write(graph_in_bytes)
	if err != nil {
		Print("Unable to store the curr user graph offline")
		os.Exit(1)
	}
}
