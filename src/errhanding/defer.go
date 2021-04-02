package errhanding

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jackSpanrrows/go-share-library/src/functions"
)

func WriteFile(filename string) {

	// 判断文件是否存在，不存在创建

	// 文件不存在要创建
	file, err := os.Create(filename)
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("%v -- 文件创建失败：%v", filename, err)
		return
	}
	file, err = os.OpenFile(filename, os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s", pathErr.Op, pathErr.Path, pathErr.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	f := functions.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, f())
	}

	return

}
