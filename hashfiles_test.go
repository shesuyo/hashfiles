package main

import (
	"runtime"
	"strconv"
	"testing"
)

func initArgs() {
	ignoreFile = ""
	ignoreDir = ""
	input = "."
	output = "out.txt"
	groutineNum = runtime.NumCPU()
}

func TestParseArgsSuccess(t *testing.T) {
	testDatas := [][]string{
		{"", "-i", "a", "-o", "b", "-if", "c", "-id", "d", "-g", "7"},
		{"", "-input", "a", "-output", "b", "-ignorefile", "c", "-ignoredir", "d", "-goroutine", "7"},
	}
	for _, testData := range testDatas {
		initArgs()
		if parseArgs(testData) != nil {
			t.Error("ParseArgs error with", testData)
		}
		if input != testData[2] {
			t.Error("ParseArgs parse input error with", testData)
		}
		if output != testData[4] {
			t.Error("ParseArgs parse output error with", testData)
		}
		if ignoreFile != testData[6] {
			t.Error("ParseArgs parse ignoreFile error with", testData)
		}
		if ignoreDir != testData[8] {
			t.Error("ParseArgs parse ignoreDir error with", testData)
		}
		if num, _ := strconv.ParseInt(testData[10], 10, 64); int(num) != groutineNum {
			t.Error("ParseArgs parse groutineNum error with", testData)
		}
	}
}

func TestParseArgsFail(t *testing.T) {
	testDatas := [][]string{
		{"hashfiles", "error"},
		{"hashfiles", "-i", "a", "error", "sad"},
	}
	for _, testData := range testDatas {
		initArgs()
		if parseArgs(testData) != IncorrectUsage {
			t.Error("ParseArgs error with", testData)
		}
	}
}
