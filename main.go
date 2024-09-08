package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// 第一引数の桁数のパスワードを生成したい
	var digitOption = flag.Int("d", -1, "set digit of password.")
	flag.Parse()

	fmt.Printf("digit: %d\n", *digitOption)
	if *digitOption == -1 {
		return
	}
	const NUM = "abcdefghijklmnopqrstuvwxyz0123456789"
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

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
