/*
	Package deprecated provides a convenient way to mark code segments as deprecated.
*/

package deprecated

import (
	"fmt"
	"reflect"
	"runtime"
)

// ThisIs returns an error for a list of deprecated items.
func ThisIs(list ...interface{}) error {
	if len(list) == 0 {
		return fmt.Errorf("nothing provided, nothing deprecated")
	}
	lv := "is"
	if len(list) > 1 {
		lv = "are"
	}
	var out string
	for _, item := range list {
		out += fmt.Sprintf("%s ", runtime.FuncForPC(reflect.ValueOf(item).Pointer()).Name())
	}
	return fmt.Errorf("%s%s deprecated", out, lv)
}
