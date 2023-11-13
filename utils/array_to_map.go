package utils

func ArrayToMap(arr []string) map[string]string {
	result := make(map[string]string)

	for i := 0; i < len(arr)-1; i += 2 {
		result[arr[i]] = arr[i+1]
	}

	return result
}
