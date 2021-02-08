package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func main() {
	findProcess(8405)
}

/*
  findProcess makes sure the PID has a proc file. If it doesn't then
  it means the process doesn't exist anymore.
*/
func findProcess(_pid int) (string, error) {
	root := "/proc"
	pid := strconv.Itoa(_pid)

	pidPath := path.Join(root, pid)
	_, err := os.Stat(pidPath)
	if os.IsNotExist(err) {
		return "", err
	}
	return pidPath, nil
}

// parse /proc/<pid>/io

func getChildren(pidPath string) ([]string, error) {
	// get the PID so we can build the /proc/<pid>/task/<tid>/children path
	pid := filepath.Base(pidPath)
	taskPath := path.Join(pidPath, "task")
	tidPath := path.Join(taskPath, pid)
	childrenPath := path.Join(tidPath, "children")
	data, err := ioutil.ReadFile(childrenPath)
}
