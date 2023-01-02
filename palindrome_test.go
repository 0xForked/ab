package lorem_test

import (
	"math"
	"testing"
)

// BenchmarkPalindrome-10    	198891538	         5.995 ns/op
// BenchmarkPalindrome-10    	196673727	         5.987 ns/op
func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome1(2579752)
	}
}

// BenchmarkPalindrome2-10    	86398633	        13.91 ns/op
// BenchmarkPalindrome2-10    	81584090	        13.97 ns/op
func BenchmarkPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome0(2579752)
	}
}

func isPalindrome0(num int) bool {
	d := 0
	for i := num; i > 0; i /= 10 {
		d++
	}

	for i := 0; i < d/2; i++ {
		if (num/int(math.Pow10(i)))%10 != (num/int(math.Pow10(d-i-1)))%10 {
			return false
		}
	}

	return true
}

func isPalindrome1(num int) bool {
	d := 0
	for i := num; i > 0; i /= 10 {
		d++
	}

	for i := 0; i < d/2; i++ {
		powerOfTen := 1
		for j := 0; j < i; j++ {
			powerOfTen *= 10
		}
		if (num/powerOfTen)%10 != (num/(powerOfTen*10))%10 {
			return false
		}
	}

	return true
}
