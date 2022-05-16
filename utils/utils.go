package utils

func DeepCopy(arr [][]int) [][]int {
	rst := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		rstArr := make([]int, len(arr[i]))
		for j := 0; j < len(arr[i]); j++ {
			rstArr[j] = arr[i][j]
		}
		rst[i] = rstArr
	}
	return rst
}

func Equal(arr1, arr2 [][]int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if len(arr1[i]) != len(arr2[i]) {
			return false
		}
		for j := 0; j < len(arr1[i]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}
	return true
}
