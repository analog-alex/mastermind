package failures

import (
	"fmt"
	"os"
)

func ToStderr(msg string, err error) {
	if err == nil {
		_, _ = fmt.Fprintln(os.Stderr, msg)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, msg, err.Error())
	}
}
