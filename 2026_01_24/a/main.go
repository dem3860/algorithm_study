package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	sc.Scan()
	S := sc.Text()

	ans := 0

	for _, char := range S {
		if char == 'i' || char == 'j' {
			ans++
		}
	}

	fmt.Fprintln(wr, ans)
}
