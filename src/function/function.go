package function

func GetMapKeys(mp map[string]interface{}) []string {
	keys := make([]string, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	return keys
}

func InterfaceArrayToStringArray(arr []interface{}) []string {
	narr := make([]string, len(arr))
	for i, v := range arr {
		narr[i] = v.(string)
	}
	return narr
}

func RemoveStringSliceRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
