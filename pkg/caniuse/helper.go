package caniuse

import (
	"path/filepath"
	"runtime"
)

func getDataPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	path, err := filepath.Abs(basepath + "/../../data.json")
	if err != nil {
		panic(err)
	}
	return path
}
