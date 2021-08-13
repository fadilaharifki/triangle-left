// result
//     *
//    **
//   ***
//  ****
// *****

package main

import "fmt"

func main() {
	result := ""

	for i := 0; i < 5; i++ {
		for j := 4 - i; j > 0; j-- {
			result += " "
		}
		for k := 5 - i; k < 6; k++ {
			result += "*"
		}
		fmt.Println(result)
		result = ""
	}
}
