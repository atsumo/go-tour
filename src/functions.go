/**
関数は、０個以上の引数を取ることができる
変数名の後ろに方を書く
 */
package main
import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
