package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func toInt(value string) int {
	num, _ := strconv.Atoi(value)
	return num
}

func splitWithInt(value string) []int {
	vals := strings.Split(value, " ")
	nums := make([]int, len(vals))
	for index, elem := range vals {
		nums[index] = toInt(elem)
	}
	return nums
}

func main() {
	ins := splitWithInt(gets())
	fmt.Println(ins[0]*ins[1], (ins[0]+ins[1])*2)
}
