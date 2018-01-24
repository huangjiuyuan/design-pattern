package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	lc := LoggerChain()
	lc.LogMessage(INFO, "This is an information.")
	lc.LogMessage(DEBUG, "This is an debug level information.")
	lc.LogMessage(ERROR, "This is an error information.")
}
