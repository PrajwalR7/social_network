package utils

import "slices"

func FilterUsers(usermap map[string][]string, filterList []string) []string {
	var totalList []string
	for keys := range usermap {
		totalList = append(totalList, keys)
	}
	totalList = totalList[:]
	filteredList := make([]string, 0)

	for _, filterItem := range filterList {
		if !slices.Contains(totalList, filterItem) {
			filteredList = append(filteredList, filterItem)
		}
	}

	return filteredList
}
