package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewFileReader(name string) FileReader {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return FileReader{
		file:    f,
		scanner: bufio.NewScanner(f),
	}
}

func (r *FileReader) NextInt() (int, error) {
	if r.scanner.Scan() {
		i, _ := strconv.ParseInt(r.scanner.Text(), 10, 0)
		return int(i), nil
	}
	return 0, fmt.Errorf("no more input")
}

func (r *FileReader) SplitNextIntLine(sep string) ([]int, error) {
	var res []int
	if r.scanner.Scan() {
		s := strings.Split(r.scanner.Text(), sep)
		for _, str := range s {
			i, _ := strconv.ParseInt(str, 10, 0)
			res = append(res, int(i))
		}
		return res, nil
	}
	return nil, fmt.Errorf("no more input")
}

func (r *FileReader) SplitNextLine(sep string) ([]string, error) {
	if r.scanner.Scan() {
		return strings.Split(r.scanner.Text(), sep), nil
	}
	return nil, fmt.Errorf("no more input")
}

func (r *FileReader) NextLine() (string, error) {
	if r.scanner.Scan() {
		return r.scanner.Text(), nil
	}
	return "", fmt.Errorf("no more input")
}

func (r *FileReader) NextLineSplit(sep string) ([]string, error) {
	if r.scanner.Scan() {
		return strings.Split(r.scanner.Text(), sep), nil
	}
	return nil, fmt.Errorf("no more input")
}

func (r *FileReader) Close() {
	r.file.Close()
}
