package array

func GetQueueNext(index ,mod int) int {
	index++
	index = index % mod
	return index
}
