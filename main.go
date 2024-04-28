package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
	fmt.Println(convert("PAYPALISHIRING", 1000) == "PAYPALISHIRING")
	fmt.Println(convert("A", 1) == "A")
	fmt.Println(convert("ABCDE", 5) == "ABCDE")
}

// 0  2  4  6  8  10  12
// 1  3  5  7  9  11  13

// 0     4      8      12
// 1  3  5  7   9  11  13
// 2     6     10      14

// 0        6          12
// 1     5  7      11  13
// 2  4     8  10      14
// 3        9          15

// 0            8              16
// 1        7   9          15  17
// 2     6     10      14      18
// 3  5        11  13          19
// 4           12              20

func convert(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	} else {
		var result string
		for i := 0; i < numRows; i++ {
			for j := 0; j < len(s); j++ {
				if j%(numRows*2-2) == i || j%(numRows*2-2) == numRows*2-2-i {
					//fmt.Printf("%c ", s[j])
					result += string(s[j])
				}
			}
		}
		return result
	}
}
