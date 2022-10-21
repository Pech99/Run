package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/Pech99/Run/us32"
)

const confFile string = "config.txt"

func main() {

	hwnd := us32.GetForegroundWindow()
	title := us32.GetWindowText(us32.HWND(hwnd))

	if len(os.Args) < 2 {
		return
	}
	lable := os.Args[1]

	comand, err := getComand(title, lable)
	if err != nil {
		return
	}

	fmt.Println("com:", comand)

}

func getComand(title string, lable string) (string, error) {
	file, err := ioutil.ReadFile(confFile)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile("[\n\r]+")
	file = re.ReplaceAll(file, []byte("\n"))

	rows := strings.Split(string(file), "\n")
	var i int
	var row string
	for i, row = range rows {
		if len(row) < 1 || unicode.IsSpace(rune(row[0])) || row[0] == '#' || row[0] == ';' {
			continue
		} else if titleCheck(row, title) {
			break
		}
	}

	for _, row = range rows[i+1:] {
		if trow := strings.TrimSpace(row); len(trow) < 1 || trow[0] == '#' || trow[0] == ';' {
			continue

		} else if !unicode.IsSpace(rune(row[0])) {
			return "", errors.New("nessuna corrispondenza trovata")

		} else if action := strings.SplitN(trow, "=", 2); len(action) != 2 {
			continue

		} else if action[0] == lable {
			return strings.TrimSpace(action[1]), nil
		}
	}

	return "", nil
}

func titleCheck(row string, title string) bool {
	if strings.ToUpper(strings.TrimSpace(row)) == "DEFAULT" {
		return true
	}

	if row[0] == '+' {
		return regCheck(row[1:], title)
	}

	if row[0] == '\\' {
		title = title[1:]
	}
	return row == title
}

func regCheck(exp string, text string) bool {
	re := regexp.MustCompile(exp)
	return re.Match([]byte(text))
}
