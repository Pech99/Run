package us32

import (
	"syscall"
	"unsafe"
)

var (
	us32                    = syscall.MustLoadDLL("user32.dll")
	procGetWindowText       = us32.MustFindProc("GetWindowTextW")
	procGetWindowTextLength = us32.MustFindProc("GetWindowTextLengthW")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

/*
func GetWindow(funcName string) uintptr {
	proc := us32.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}
*/

func GetWindowThreadProcessId(hwnd uintptr) (uintptr, error) {
	prc := us32.MustFindProc("GetWindowThreadProcessId")
	ret, _, err := prc.Call(hwnd, uintptr(unsafe.Pointer(&hwnd)))
	//fmt.Println("ProcessId: ", ret, " Error: ", err)
	return ret, err
}
