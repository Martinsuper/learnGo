package getData

import (
	"bufio"
	"fmt"
	"io"
	"mAlgorithm/aGenerateRand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadFile(path string) []int {
	fd, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(fd)
	}
	defer fd.Close()
	if err != nil {
		fmt.Println("read error:", err)
	}
	buff := bufio.NewReader(fd)
	slice := make([]int, 1000000)
	cnt := 0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		arr := strings.Split(string(data), " ")
		for i := 0; i < len(arr); i++ {
			if arr[i] == "" {
				continue
			}
			res, _ := strconv.Atoi(arr[i])
			slice[cnt] = res
			cnt++
		}
	}
	return slice[:cnt]
}

func SetData() {
	datapath, _ := filepath.Abs(filepath.Dir("./"))
	datapath += "/src/mAlgorithm/data/sortData2"
	file, error := os.Create(datapath)
	if error != nil {
		fmt.Println(error)
	}
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10; j++ {
			data := aGenerateRand.GenerateRand()
			data1 := strconv.Itoa(data) + " "
			_, _ = file.WriteString(data1)
		}
		_, _ = file.WriteString("\n")
	}
}
