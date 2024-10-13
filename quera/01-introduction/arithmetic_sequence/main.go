package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		nums := make([]int, len(inputs)-1)
		for i := 1; i < len(inputs); i++ {
			nums[i-1], _ = strconv.Atoi(inputs[i])
		}
		fmt.Println(inputs[0], calcSequenceCount(nums))
	}
}

func calcSequenceCount(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	result := 0
	dif := nums[1] - nums[0]
	counter := 2
	i := 2
	for i < len(nums) {
		if nums[i]-nums[i-1] == dif {
			counter++
			i++
		} else {
			if i == len(nums)-1 {
				break
			}
			dif = nums[i+1] - nums[i]
			counter = 2
			result += (counter - 2) * (counter - 1) / 2
			i += 2
		}
	}
	result += (counter - 2) * (counter - 1) / 2
	return result
}
