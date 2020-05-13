package hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Type TestCase implements ways of interpreting the input and output files and running code over them.

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

type Tests []Test

func NewTests(path string) *Tests {
	t := Tests(TestsReaders(path))
	return &t
}

func (t *Tests) Run(f func(Test)) {
	for k, v := range *t {
		f(v)
		if k < len(*t)-1 {
			fmt.Println()
		}
	}
}

type Test struct {
	Name string
	In   TestReader
	Out  TestReader
}

func TestsReaders(path string) []Test {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	tests := make([]Test, 0)

	for _, file := range files {
		inFileName := file.Name()
		extension := filepath.Ext(inFileName)
		if strings.HasPrefix(inFileName, "input") && extension == ".txt" {
			name := strings.TrimPrefix(inFileName, "input")
			name = strings.TrimSuffix(name, extension)

			inFilePath := filepath.Join(path, inFileName)
			outFilePath := filepath.Join(path, "output"+name+extension)

			tests = append(tests, Test{
				Name: name,
				In:   TestReader(fileReader(inFilePath)),
				Out:  TestReader(fileReader(outFilePath)),
			})

		}

	}

	return tests
}

func fileReader(filename string) bufio.Reader {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	return *bufio.NewReader(f)
}

type TestReader bufio.Reader

func (t *TestReader) NextLine() string {
	reader := bufio.Reader(*t)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		if err != io.EOF {
			log.Fatalln(err)
		}
	}

	s := string(bytes)
	if len(s) > 0 {
		if rune(s[len(s)-1]) == '\n' {
			s = s[:len(s)-1]
		}
	}

	*t = TestReader(reader)
	return s
}
