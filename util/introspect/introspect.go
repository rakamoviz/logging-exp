package introspect

import (
	"regexp"
	"runtime"
)

var (
	extractFnName = regexp.MustCompile(`^.*/(.*)$`)
)

//https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name/46289376#46289376
func GetFunctionName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return extractFnName.ReplaceAllString(frame.Function, "$1")
}
