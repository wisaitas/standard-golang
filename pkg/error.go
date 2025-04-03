package pkg

import (
	"log"
	"runtime"
)

func Error(err error) error {
	if err == nil {
		return nil
	}

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}

	fn := runtime.FuncForPC(pc)
	funcName := ""
	if fn != nil {
		funcName = fn.Name()
	}

	log.Printf("error: {\"message\": \"%s\", \"file\": \"%s:%d\", \"func\": \"%s\"}\n", err.Error(), file, line, funcName)

	return err
}
