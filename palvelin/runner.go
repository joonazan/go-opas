package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

func Compile(code string) (string, error) {
	sourcePath := path.Join(os.TempDir(), "tmpforrun"+uniqueString()+".go")

	sourceFile, err := os.Create(sourcePath)
	if err != nil {
		return "", errors.New("creating file failed: " + err.Error())
	}
	defer sourceFile.Close()
	defer os.Remove(sourcePath)

	sourceFile.WriteString(code)

	return CompileFile(sourcePath)
}

func CompileFile(sourcePath string) (string, error) {
	binaryPath := path.Join(os.TempDir(), "binaryforrun"+uniqueString())

	compileOutput, err := exec.Command("go", "build", "-o", binaryPath, sourcePath).CombinedOutput()
	if err != nil {
		return "", errors.New("Virhe kääntäessä:\n" + err.Error() + "\n" + string(compileOutput))
	}

	return binaryPath, nil
}

func uniqueString() string {
	return strconv.Itoa(int(uniqueNumber()))
}

var counter int64 = 0

func uniqueNumber() int64 {
	return atomic.AddInt64(&counter, 1)
}

func RunBinary(binaryPath string, input string) (string, error) {

	cmd := exec.Command(binaryPath)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	go func() {
		io.Copy(stdin, strings.NewReader(input))
		stdin.Close()
	}()

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

type res struct {
	err    error
	output string
}
