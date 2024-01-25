package shifter

import "reflect"

func NumberOfShift(arr, target []int) int {
	if len(arr) != len(target) {
		return -1
	}

	lenArr := len(arr)
	for shift := 0; shift < lenArr; shift++ {
		shiftedArr := make([]int, lenArr)
		copy(shiftedArr, arr[lenArr-shift:])
		copy(shiftedArr[shift:], arr[:lenArr-shift])

		if reflect.DeepEqual(shiftedArr, target) {
			return shift
		}
	}

	return -1
}
