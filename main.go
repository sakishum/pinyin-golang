package main

import (
	"fmt"

	"pinyin-golang/pinyin"
)

func main() {
	dict := pinyin.NewDict()

	s := ""
	// wǒ, hé shí néng bào fù?
	s = dict.Sentence(`我，何時能暴富？`).Unicode()
	fmt.Println(s)

}
