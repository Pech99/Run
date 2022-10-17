package main

import (
	"fmt"
	"os"

	sys "golang.org/x/sys/windows"
)

func main() {

	for {
		process, err := os.FindProcess(int(sys.GetForegroundWindow()))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(process)
	}
}

func selectRoutin() []string {

	return nil

}

/*
C:\Users\Vitto\go\pkg\src\golang.org\x\sys
C:\Users\Vitto\go\src\golang.org\x\sys
C:\Program Files\Go\src\golang.org\x\sys
*/
