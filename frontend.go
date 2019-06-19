package main

import (
	"fmt"
	"io/ioutil"
)

var code = `

`

func Write()  {
	fileName := "E:/usr/pushy/Main.java"
	var d = []byte(code)

	err := ioutil.WriteFile(fileName, d, 0666)
	if err != nil {
		fmt.Println("write fail")
		panic(err)
	} else {
		fmt.Println("write success")
	}
}

func main(){
	Write()
}