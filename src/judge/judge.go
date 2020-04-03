package judge

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func judgePE(inFileName string, outFileName string, userFileName string) int {
	f1, err := os.Open(outFileName)
	if err != nil {
		fmt.Println("read file fail", err)
		return 0
	}
	defer f1.Close()

	fd1, err := ioutil.ReadAll(f1)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return 0
	}
	str1 := string(fd1)

	f2, err := os.Open(userFileName)
	if err != nil {
		fmt.Println("read file fail", err)
		return 0
	}
	defer f2.Close()

	fd2, err := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return 0
	}
	str2 := string(fd2)

	if !bytes.Equal(fd1, fd2) {
		str1 = strings.Replace(str1, " ", "", -1)
		str1 = strings.Replace(str1, "\r\n", "", -1)

		str2 = strings.Replace(str2, " ", "", -1)
		str2 = strings.Replace(str2, "\r\n", "", -1)

		if strings.EqualFold(str1, str2) {
			return 1
		}
	}
	return 0
}

func judgeWA(inFileName string, outFileName string, userFileName string) int {
	f1, err := os.Open(outFileName)
	if err != nil {
		fmt.Println("read file fail", err)
		return 0
	}
	defer f1.Close()

	fd1, err := ioutil.ReadAll(f1)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return 0
	}

	f2, err := os.Open(userFileName)
	if err != nil {
		fmt.Println("read file fail", err)
		return 0
	}
	defer f2.Close()
	fd2, err := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return 0
	}

	if !bytes.Equal(fd1, fd2) {
		return 1
	}
	return 0
}
