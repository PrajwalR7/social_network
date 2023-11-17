package utils

type stringArr []string

// Polyfill function which adds the include method to []string type similar to what we have in JS
func (s stringArr) includes(key string) bool {
	for _, val := range s {
		if val == key {
			return true
		}
	}
	return false
}

func FilterUsers(usermap map[string][]string, filterList []string) []string {
	var totalList stringArr
	for keys := range usermap {
		totalList = append(totalList, keys)
	}
	filteredList := make([]string, 0)

	for _, filterItem := range filterList {
		if !totalList.includes(filterItem) {
			filteredList = append(filteredList, filterItem)
		}
	}

	return filteredList
}
