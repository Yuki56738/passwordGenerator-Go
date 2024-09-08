package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 第一引数の桁数のパスワードを生成したい
	var digitOption = flag.Int("d", -1, "set digit of password.")
	flag.Parse()

	fmt.Printf("digit: %d\n", *digitOption)
	if *digitOption == -1 {
		fmt.Println(os.Args[0], "-d <digit>")
		return
	}
	const NUM = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	//for i = 0; i < len(NUM); i++ {
	//	//fmt.Println(string(NUM[i]))
	//}
	rand.Seed(time.Now().UnixNano())
	var i int
	for i = 0; i <= *digitOption; i++ {
		fmt.Printf("%c", NUM[rand.Intn(len(NUM))])
	}
	fmt.Printf("\n")

	//fmt.Printf("%s", string(NUM[rand.Intn(len(NUM))]))
}
