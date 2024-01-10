package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetProjectRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		dir = filepath.Dir(dir)
	}
}

func GetCallerLocationSkip(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	return file + ":" + fmt.Sprint(line)
}
