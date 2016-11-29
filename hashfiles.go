package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

var (
	ignoreFile  string
	ignoreDir   string
	input       string = "."
	output      string = "out.txt"
	groutineNum int    = runtime.NumCPU()

	IncorrectUsage = errors.New("Incorrect Usage.")
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		fmt.Println(tip)
		return
	}

	hr := NewHashRoom(input, output, ignoreDir, ignoreFile, groutineNum)
	hr.Run()
}

//解析传入的参数
func parseArgs(args []string) error {
	var length = len(args)
	//只允许一个参数有一个值，且参数不能为空。
	if length%2 == 0 {
		return IncorrectUsage
	}
	for i := 1; i < length-1; i += 2 {
		switch args[i][1:] {
		case "ignorefile", "if":
			ignoreFile = args[i+1]
		case "ignoredir", "id":
			ignoreDir = args[i+1]
		case "input", "i":
			input = args[i+1]
		case "output", "o":
			output = args[i+1]
		case "goroutine", "g":
			num, err := strconv.ParseInt(args[i+1], 10, 64)
			if err != nil {
				return IncorrectUsage
			}
			groutineNum = int(num)
		default:
			return IncorrectUsage
		}
	}
	return nil
}
