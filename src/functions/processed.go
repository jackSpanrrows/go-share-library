package functions

import "errors"

type Traveser func(ele interface{})

type Processes struct {
	IntSlice   []int
	SortByAsc  string
	FloatSlice []float32
}

/**
  Process:封装公共切片数组操作
*/
func Process(array interface{}, traveser Traveser) error {
	if array == nil {
		return errors.New("nil pointer")
	}
	var length int // 数组的长度
	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))
	case []float32:
		length = len(array.([]float32))
	default:
		return errors.New("error type")
	}
	if length == 0 {
		return errors.New("len is zero.")
	}
	traveser(array)
	return nil
}

/**
具体操作：升序排序数组元素
*/
func SortByAscending(ele interface{}) {
	intSlice, ok := ele.([]int)
	if !ok {
		return
	}
	length := len(intSlice)

	for i := 0; i < length; i++ {
		isChange := false
		for j := 0; j < length-1; j++ {

			if intSlice[j] > intSlice[j+1] {
				isChange = true
				intSlice[j], intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
		if isChange == false {
			return
		}
	}
}

/**
具体操作:降序排序数组元素
*/
func SortByDescending(ele interface{}) {
	intSlice, ok := ele.([]int)
	if !ok {
		return
	}
	length := len(intSlice)
	for i := 0; i < length; i++ {
		isChange := false
		for j := 0; j < length-1; j++ {
			if intSlice[j] < intSlice[j+1] {
				isChange = true
				intSlice[j], intSlice[j+1] = intSlice[j+1], intSlice[j]
			}
		}
		if isChange == false {
			return
		}
	}
}
