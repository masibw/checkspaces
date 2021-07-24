package a

import(
 _ "embed"
	"fmt"
)

var (
	//go:embed testfile.txt
	file []byte

	// go:embed testfile.txt // want "There is a space between slash and go:embed"
	fileInvalid []byte
)

func f() {
	fmt.Println(file, fileInvalid)
}
