/**
 * author	: masakokh
 * year		: 2022
 * note		: start up function
 */
package main

//
import (
	"fmt"
)

// /**
//  *
//  */
// func strToBinary(s string, base int) []byte {

//     var b []byte

//     for _, c := range s {
//         b = strconv.AppendInt(b, int64(c), base)
//     }

//     return b
// }

func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b\n", res, c)
	}
	return res
}

/**
 *
 */
func main() {
	// fmt.Println("Hello, World!")

	// strVar := "100"

	// for i, character := range strVar {
	// 	// fmt.Println("ascii:", character, int(character))
	// 	fmt.Printf("Start Index: %d Value:%s  ascii:%d\n ", i, string(character), int(character), byte(int(character)))
	// }
	fmt.Print(binary("kun"))
}
