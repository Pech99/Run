package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode"

	"github.com/Pech99/Run/us32"
)

//complile: go build -ldflags "-H windowsgui"
//hiden console

const confFile string = "./config.conf"

func main() {

	hwnd := us32.GetForegroundWindow()
	title := us32.GetWindowText(us32.HWND(hwnd))
	class := us32.GetClassNameW(us32.HWND(hwnd))

	if len(os.Args) < 2 {
		return
	}
	lable := os.Args[1]

	fmt.Println("title:", title, "\nlable:", lable)

	comand, err := getComand(title, class, lable)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if strings.ToUpper(strings.TrimSpace(comand)) == "INFO" {
		info := fmt.Sprint("hwnd: ", hwnd, "\nclass: ", class, "\ntitle: ", title)
		us32.MessageBox(0, info, "Foreground Window Info", 0)
		return
	}

	out, err := exeComand(comand)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("out:", out)

}

func getComand(title string, class string, lable string) (string, error) {
	file, err := os.ReadFile(confFile)
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
		} else if corrCheck(row, title, class) {
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

func corrCheck(row string, title string, class string) bool {
	if strings.ToUpper(strings.TrimSpace(row)) == "DEFAULT" {
		return true
	}

	target := title
	if row[0] == '@' {
		target = class
		row = row[1:]
	}

	if row[0] == '+' {
		return regCheck(row[1:], target)
	}

	if row[0] == '\\' {
		target = target[1:]
	}
	return row == target
}

func regCheck(exp string, text string) bool {
	re := regexp.MustCompile(exp)
	return re.Match([]byte(text))
}

func exeComand(comand string) (string, error) {
	cmdargs := strings.Fields(comand)
	var out []byte
	var err error

	if len(cmdargs) > 1 {
		out, err = exec.Command(cmdargs[0], cmdargs[1:]...).Output()
	} else {
		out, err = exec.Command(cmdargs[0]).Output()
	}

	return string(out), err
}
