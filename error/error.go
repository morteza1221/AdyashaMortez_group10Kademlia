package error

import (
	"fmt"
	"os"
)

func CheckFatalError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
