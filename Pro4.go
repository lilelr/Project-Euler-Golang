package main

func reverse(n int) int {
	reversed := 0
	for n > 0 {
		reversed = 10*reversed + n%10
		n = n/10
	}
	return reversed
}

func isPalindrome(n int) bool{
	return n == reverse(n)
}

func main() {
	largestPalindrome := 0
	a := 999
        var b int
	var db int
	for a>=100{
		if a%11 ==0 {
			b =999
			db =1
		} else {
			b = 999
			db = 11
		}

		for b>=a {
			if a*b <= largestPalindrome {
				break
			}
			if(isPalindrome(a*b)){
				largestPalindrome = a*b
			}
			b = b-db
		}
		a = a-1
	}
	println(largestPalindrome)

}
