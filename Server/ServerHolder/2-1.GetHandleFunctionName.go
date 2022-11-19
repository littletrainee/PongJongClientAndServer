package ServerHolder

import (
	"reflect"
	"runtime"
	"strings"
)

func GetHandleFunctionName(function any) string {
	temp := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	end := strings.Index(temp, "-")
	temp = temp[69:end] // Start index at 69 end index at minus -1
	return temp
}
