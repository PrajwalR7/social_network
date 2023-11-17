package utils

func Add(list []string, item string) []string {
	alreadyPresent := false
	for _, val := range list {
		if val == item {
			alreadyPresent = true
			break
		}
	}
	if !alreadyPresent {
		list = append(list, item)
	}
	return list
}
