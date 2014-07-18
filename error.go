// Package viewerr ...
package viewerr

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

// VERSION the version info.
const VERSION = "0.1"

// getFirstInt getting the firse args param.
func getFirstInt(args ...interface{}) int {
	var count = 0
	var firstArg interface{}
	for _, arg := range args {
		count++
		if 1 == count {
			firstArg = arg
			break
		}
	}
	var trackStack = 1
	if firstArg != nil {
		if _, ok := firstArg.(int); ok {
			trackStack = firstArg.(int)
		}
	}

	return trackStack
}

// WrapError getting the error line and file, echo error format such as,
// '[migo.go:74] ERROR error...'.
func WrapError(err error, args ...interface{}) (nerr error) {
	if err == nil {
		return
	}

	trackStack := getFirstInt(args...)
	_, file, line, _ := runtime.Caller(trackStack)
	file = filepath.Base(file)
	errMsg := err.Error()

	if !strings.HasPrefix(errMsg, "DEBUG") && !strings.HasPrefix(errMsg, "INFO") {
		errMsg = "ERROR " + errMsg
	}
	errMsg = fmt.Sprintf("[%s:%d] %s", file, line, errMsg)

	return errors.New(errMsg)
}

// AddrWrapError getting the http error line, file, addr and url,
// echo error format such as,
// '[migo.go:74] [127.0.0.1 /index/index] ERROR error...'
func AddrWrapError(r *http.Request, err error, args ...interface{}) (nerr error) {
	if err == nil {
		return
	}

	trackStack := getFirstInt(args...)
	_, file, line, _ := runtime.Caller(trackStack)
	file = filepath.Base(file)
	errMsg := err.Error()
	addr := strings.Split(r.RemoteAddr, ":")[0]
	url := r.URL.Path

	var prefix string
	if !strings.HasPrefix(errMsg, "DEBUG") && !strings.HasPrefix(errMsg, "INFO") {
		prefix = "ERROR"
	}
	errMsg = fmt.Sprintf("[%s:%d] [%s %s] %s %s", file, line, addr, url, prefix, errMsg)

	return errors.New(errMsg)
}

// WrapErrorf echo the error line and file with format.
func WrapErrorf(format string, args ...interface{}) (err error) {
	return WrapError(fmt.Errorf(format, args...), int(2))
}

// AddrWrapErrorf echo the http error line and file path with format.
func AddrWrapErrorf(r *http.Request, format string, args ...interface{}) (err error) {
	return AddrWrapError(r, fmt.Errorf(format, args...), int(2))
}

// DumpStack echo the error line and file path in io Writer.
func DumpStack(w io.Writer) {
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		w.Write([]byte(fmt.Sprintf("%s:%d\n", file, line)))
	}
}
