package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	TimeInStart := 0

	filePath := "main.go"
	oldString := "TimeInStart := " + strconv.Itoa(TimeInStart)
	newString := "TimeInStart := " + strconv.Itoa(TimeInStart+1)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	newContent := strings.Replace(string(content), oldString, newString, -1)

	err = ioutil.WriteFile(filePath, []byte(newContent), 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Прошло минут:", TimeInStart)
}
