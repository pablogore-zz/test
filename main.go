package main

import (
	"fmt"
	"strconv"
)


var MAX_CONCURRENT_CONNECTION = 15
var START_TIME_SCRIPT string
var END_TIME_SCRIPT string
var INPUT_FILE string
var OUTPUT_FILE string
var OUTPUT_FILE_FAILED string
var DEBUG_LEVEL_SHORT = 1
var DEBUG_LEVEL_LONG = 2


func main() {
	process()

}

/**
 * Core processing of the script
 */
func process() {
	var strs []string

	for i := 1990; i <= 2018; i++ {
		file := fmt.Sprintf("laucnty%s", strconv.Itoa(i)[2:])
		strs = append(strs, file)
		fmt.Println(strs)

	}

	fmt.Print(strs)

}
