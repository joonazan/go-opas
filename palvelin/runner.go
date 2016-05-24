package main

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strconv"
	"sync/atomic"
	"time"
)

func Run(code string) (string, error) {
	id := strconv.Itoa(int(getUniqueNumber()))
	sourcePath := path.Join(os.TempDir(), "tmpforrun"+id+".go")

	sourceFile, err := os.Create(sourcePath)
	if err != nil {
		return "", errors.New("creating file failed: " + err.Error())
	}
	defer sourceFile.Close()
	defer os.Remove(sourcePath)

	sourceFile.WriteString(code)

	binaryPath := path.Join(os.TempDir(), "binaryforrun"+id)

	compileOutput, err := exec.Command("go", "build", "-o", binaryPath, sourcePath).CombinedOutput()
	if err != nil {
		return "", errors.New("Virhe kääntäessä:\n" + string(compileOutput))
	}
	defer os.Remove(binaryPath)

	cmd := exec.Command(binaryPath)
	c := make(chan res)

	go func() {
		output, err := cmd.CombinedOutput()
		c <- res{err, string(output)}
	}()

	select {
	case s := <-c:
		if s.err != nil {
			return "", errors.New("Ajonaikainen virhe:\n" + s.output)
		} else {
			return s.output, nil
		}
	case <-time.After(time.Second):
		cmd.Process.Kill()
		return "", errors.New("Ohjelmalla kesti yli sekunnin.")
	}
}

var counter int64 = 0

func getUniqueNumber() int64 {
	return atomic.AddInt64(&counter, 1)
}

type res struct {
	err    error
	output string
}
