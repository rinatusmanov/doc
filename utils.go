package doc

import (
	"fmt"
	"math/rand"
	"time"
)

type Tag interface {
	Tag() string
}

func delimiter(space int) (result string) {
	for i := 0; i < space; i++ {
		result += "\t"
	}
	return
}

const RandomLength = 36

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func random() string {
	return fmt.Sprintf("UUID%v%v%v%v",
		r.Int63(),
		r.Int63(),
		r.Int63(),
		r.Int63(),
	)
}

type tWriter struct {
	str string
}

func (w *tWriter) Write(p []byte) (n int, err error) {
	w.str += string(p)
	return len(p), nil
}

func (w *tWriter) String() (str string) {
	return w.str
}
