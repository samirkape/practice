// Compare two arrays generated by SHA256 from crypto package

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	m1 := sha256.Sum256([]byte("X"))
	m2 := sha256.Sum256([]byte("x"))

	fmt.Printf("%x\n%x\n%t\n%T\n%d\n", m1, m2, m1 == m2, m1, len(m1))

	m1h := fmt.Sprintf("%x", m1)
	m2h := fmt.Sprintf("%x", m2)

	diff := diff_counter(m1h, m2h)

	fmt.Println(diff)

}

func diff_counter(m1, m2 string) int {
	j, counter := 0, 0
	if m1 == m2 {
		return 0
	} else {
		for i := 0; i < len(m1); i++ {
			if m1[i] != m2[j] {
				counter++
				j++
			}
		}
	}
	return counter
}