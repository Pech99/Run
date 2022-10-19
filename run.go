package main

import (
	"fmt"

	"github.com/Pech99/Run/us32"
)

func main() {
	hwndold := us32.GetWindow("GetForegroundWindow")
	for {
		if hwnd := us32.GetWindow("GetForegroundWindow"); hwnd != 0 && hwndold != hwnd {
			text := us32.GetWindowText(hwnd)
			pid, _ := us32.GetWindowThreadProcessId(hwnd)
			fmt.Print("hwnd: ", hwnd, "\t - pid: ", pid, "\t- window: ", text, "\n")
			hwndold = hwnd
		}
	}
}
