package base

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"html/template"
)
type tools struct {

}
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

func GetMaxInt64(x int64, y int64) int64 {
	if x >= y {
		return x
	}else {
		return y
	}
}

func ReadAll(dir string) []byte {
	file,_ := os.Open(dir)
	defer file.Close()
	result,_ := ioutil.ReadAll(file)
	return result
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func ReadCodeFromFile(dir string) string {
	fmt.Println("ReadCode path: " + dir)
	file, _ := os.Open(dir)
	defer file.Close()
	reader := bufio.NewReader(file)
	var code string

	for {
		str, err := reader.ReadString('\n')
		code = code + str
		if err == io.EOF || err != nil {
			if str == "" {
				break
			}
		}
	}
	fmt.Println(code)
	return code
}


func LoadTemplates() interface{} {
	var Template interface{}
	var templates[] string
	fn := func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Printf("failed, load template:%v", err)
			return nil
		}

		if !f.IsDir() && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}

		return nil
	}

	filepath.Walk("views", fn)
	if len(templates) > 0 {
		Template = template.Must(template.ParseFiles(templates...))
	} else {
		Template = template.New("")
	}
	return Template
}