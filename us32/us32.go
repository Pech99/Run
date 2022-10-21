package us32

import (
	"syscall"
	"unsafe"
)

/*
var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetWindowText       = user32.NewProc("GetWindowTextW")
	procGetWindowTextLength = user32.NewProc("GetWindowTextLengthW")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindoW")
	procMessageBox          = user32.NewProc("MessageBoxW")
)*/

var (
	user32                  = syscall.MustLoadDLL("user32.dll")
	procGetWindowText       = user32.MustFindProc("GetWindowTextW")
	procGetWindowTextLength = user32.MustFindProc("GetWindowTextLengthW")
	procGetForegroundWindow = user32.MustFindProc("GetForegroundWindow")
	procMessageBox          = user32.MustFindProc("MessageBoxW")
	//procGetWindowThreadProcessId = user32.MustFindProc("GetWindowThreadProcessId")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

// restituisce la lunghezza del testo della finiesra con id hwnd
func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

// restituisce il testo della finiesra con id hwnd
func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen),
	)

	return syscall.UTF16ToString(buf)
}

func GetForegroundWindow() uintptr {
	hwnd, _, _ := procGetForegroundWindow.Call()
	return hwnd
}

func MessageBox(hwnd HWND, title, caption string, flags uint) int {
	ret, _, _ := procMessageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(flags),
	)

	return int(ret)
}

/*
func GetWindowThreadProcessId(hwnd uintptr) uintptr {
	ret, _, _ := procGetWindowThreadProcessId.Call(hwnd, uintptr(unsafe.Pointer(&hwnd)))
	return ret
}
*/

/*
func GetWindow(funcName string) uintptr {
	proc := user32.MustFindProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}
*/
