package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestHashRoom(t *testing.T) {
	defer func() {
		os.Remove("testout")
	}()

	testDatas := []struct {
		dir          string
		output       string
		ignoreDir    string
		ignoreFile   string
		goroutineNum int
		expect       string
	}{
		{`testdata\a`, "testout", "", "", 1, "testdata\\a\\b.txt,E9D71F5EE7C92D6DC9E92FFDAD17B8BD49418F98,1\r\n"},
		{`testdata\b`, "testout", "[0-9]", "", 1, "testdata\\b\\a\\e.txt,58E6B3A414A1E090DFC6029ADD0F3555CCBA127F,1\r\n"},
		{`testdata\c`, "testout", "", "[0-9]", 1, "testdata\\c\\a.txt,86F7E437FAA5A7FCE15D1DDCB9EAEAEA377667B8,1\r\n"},
		{`testdata\1`, "testout", "[a-z]", "[0-9]", 1, "testdata\\1\\1\\a.txt,86F7E437FAA5A7FCE15D1DDCB9EAEAEA377667B8,1\r\n"},
	}

	for _, testData := range testDatas {
		rm := NewHashRoom(testData.dir, testData.output, testData.ignoreDir, testData.ignoreFile, testData.goroutineNum)
		rm.Run()
		bs, _ := ioutil.ReadFile(testData.output)
		if string(bs) != testData.expect {
			t.Errorf("TestHashRoom error expect with data: %+v\n", testData)
		}
	}
}
