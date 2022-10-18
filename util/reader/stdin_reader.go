package reader

import (
	"strconv"
	"strings"
)

type StdInReader struct {
	osArgs []string
	idx    int
}

func NewStdInReader(osArgs []string) StdInReader {
	return StdInReader{
		osArgs: osArgs,
		idx:    1,
	}
}

func (r *StdInReader) NextInt() int {
	i, _ := strconv.ParseInt(r.osArgs[r.idx], 10, 0)
	r.idx += 1
	return int(i)
}

func (r *StdInReader) NextIntLine(sep string) (res []int) {
	s := strings.Split(r.osArgs[r.idx], sep)
	for _, str := range s {
		i, _ := strconv.ParseInt(str, 10, 0)
		res = append(res, int(i))
	}
	r.idx += 1
	return
}

func (r *StdInReader) NextLine() (res string) {
	r.idx += 1
	return r.osArgs[r.idx-1]
}