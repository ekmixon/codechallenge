package solution

import "fmt"

// Returns greatest common divisor of a and b
func gcd(a, b, res int) int {

	if a == b {
		return res * a

	} else if (a%2 == 0) && (b%2 == 0) {
		return gcd(a/2, b/2, 2*res)

	} else if a%2 == 0 {
		return gcd(a/2, b, res)

	} else if b%2 == 0 {
		return gcd(a, b/2, res)

	} else if a > b {
		return gcd(a-b, b, res)

	} else {
		return gcd(a, b-a, res)
	}

}

func gcdNew(x, y int) int {
	
	for y >= 1 {
		
		fmt.Printf("%v : %v\n", x, y)

		if x < y {
			x, y = y, x
		}
		x, y = y, x % y

		
	}
	return x
}