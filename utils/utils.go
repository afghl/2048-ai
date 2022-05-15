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
