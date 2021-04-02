package functions

type FilterFunc func(elem interface{}) interface{}

func Data(arr interface{}, filterFunc FilterFunc) interface{} {
	slice := make([]int, 0)
	array, _ := arr.([]int)

	for _, value := range array {
		integer, ok := filterFunc(value).(int)
		if ok {
			slice = append(slice, integer)
		}
	}
	return slice
}

/**
奇数变偶数
*/
func EventFilter(ele interface{}) interface{} {
	integer, ok := ele.(int)
	if ok {
		if integer%2 == 1 {
			integer = integer + 1
		}
	}
	return integer
}

/**
偶数变奇数
*/
func OddFilter(ele interface{}) interface{} {
	integer, ok := ele.(int)
	if ok {
		if integer%2 != 1 {
			integer = integer + 1
		}
	}
	return integer
}
