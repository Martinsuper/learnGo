package getData

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []int{
	fd, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(fd)
	}
	defer fd.Close()
	if err != nil {
		fmt.Println("read error:", err)
	}
	buff := bufio.NewReader(fd)
	slice := make([]int,100000)
	cnt:=0
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}
		arr := strings.Split(string(data), " ")
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
			if arr[i] == "" {
				continue
			}
			res, _ := strconv.Atoi(arr[i])
			slice[cnt] = res;cnt++
		}
	}
	return slice[:cnt]
}
