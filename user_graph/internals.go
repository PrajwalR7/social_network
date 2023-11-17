package user_graph

import "social/utils"

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
