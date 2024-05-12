package main

import (
	"fmt"
)

/*
init関数は特殊な関数で、パッケージの初期化に使われます。
mainパッケージに書くとmain関数より先に実行されます。
*/
func init() {
	fmt.Printf("start")
}

func main() {
}

// todo  wire、main、docker、config、db、transaction
