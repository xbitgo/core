package tool_file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Exists .
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// WriteFile .
func WriteFile(name string, data []byte) error {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

// ReadListFile : 按行读文件
func ReadListFile(filename string) []string {
	result := make([]string, 0)
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
		if err == io.EOF {
			break
		}
	}
	return result
}

// WriteListFile : 按行存储文件
func WriteListFile(filename string, value []string) {
	fileHandler, err := os.Create(filename)
	defer fileHandler.Close()
	if err != nil {
		return
	}
	buf := bufio.NewWriter(fileHandler)
	for _, v := range value {
		_, _ = fmt.Fprintln(buf, v)
	}
	buf.Flush()
}
