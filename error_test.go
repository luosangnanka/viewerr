// Package viewerr ...
package viewerr

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	err := errors.New("test error")
	if err != nil {
		fmt.Println(WrapError(err))
		return
	}

	t.Fatal(fmt.Errorf("err not nil"))
}
