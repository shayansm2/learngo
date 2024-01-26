package main

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	result := (*numbers)[0]
	for _, val := range *numbers {
		if val < result {
			result = val
		}
	}
	return result
}

func ReverseSlice(numbers *[]int) {
	for i := 0; i < len(*numbers)/2; i++ {
		SwapElements(numbers, i, len(*numbers)-i-1)
	}
}

func SwapElements(numbers *[]int, i, j int) {
	if i < 0 || j < 0 {
		return
	}

	if i >= len(*numbers) || j >= len(*numbers) {
		return
	}

	(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
}
