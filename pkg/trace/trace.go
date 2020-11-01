package trace

import "runtime"
// Info stores caller's current line trace information
// Emulate __FILE__ __LINE__ and etc
type Info struct {
	File string
	Line int
	FunctionName string
}

// Trace returns caller's current line TraceInfo
func Trace() Info {
	pc := make([]uintptr, 10)
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    file, line := f.FileLine(pc[0])
	return Info{
		File: file,
		Line: line,
		FunctionName: f.Name(),
	}
}