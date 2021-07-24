package specification

import(
	_ "embed"
	"fmt"
)

var (
	//any:something
	any string

	// any:something // want "There is a space between slash and the directive: any"
	any2 []byte
)

func f() {
	fmt.Println(any, any2)
}
