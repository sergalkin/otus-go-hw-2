package main

import (
	"bufio"
	"fmt"
	"github.com/sergalkin/otus-go-hw-2/internal/stringtransformer"
	"os"
	"strings"
)

func main() {
	strBytes := getStringToTransform()
	transFormResult := stringtransformer.Transform(strBytes)
	if transFormResult.Len() == 0 {
		fmt.Println("Некорректная строка")
	} else {
		fmt.Println("Результат распаковки:", transFormResult.String())
	}
}

func getStringToTransform() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку для распаковки: ")
	str, err := reader.ReadString('\n')
	str = strings.Trim(str, " \n")

	if err != nil {
		fmt.Printf("%#v \n", err)
		getStringToTransform()
	}

	if len(str) < 1 {
		getStringToTransform()
	}
	return []byte(str)
}
