package main

import (
	"errorCode"
	"fmt"
)

func main() {
	errorCode.SetLG(1)
	fmt.Println(errorCode.CodeMap[1])
}


