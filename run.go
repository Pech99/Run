package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/Pech99/Run/us32"
)

func main() {

	hwnd := us32.GetForegroundWindow()
	title := us32.GetWindowText(us32.HWND(hwnd))

	if len(os.Args) < 2 {
		return
	}
	lable := os.Args[1]

	com, err := getCom(title, lable)
	if err != nil {
		return
	}

	_ = com
	fmt.Print("hwnd: ", hwnd, "\t- window: ", title, "\n")
}

func getCom(FGtitle string, lable string) (string, error) {
	file, err := ioutil.ReadFile("config.conf")
	if err != nil {
		return "", err
	}

	rows := strings.Split(string(file), "\n")

	for _, row := range rows {
		if len(row) < 1 || row[0] == ';' {
			continue
		}

		col := strings.Split(row, ",")
		if col[0][0] == '+' {

			return row[strings.Index(row, ","):], nil
		}
		if col[0][0] == '-' && FGtitle == col[0][1:] {
			return row[strings.Index(row, ","):], nil
		}

	}

	return string(file), nil
}

func mechProcess() {

}

func mechLable() {

}

func find(text string, exp string) bool {
	re := regexp.MustCompile(exp)
	return re.Match([]byte(text))
}
