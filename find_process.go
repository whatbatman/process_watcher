package main

import (
	"os"
	"path"
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
