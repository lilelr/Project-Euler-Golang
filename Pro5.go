package main

import "math"

func main() {

	//primes := []int{2,3,5,7,11,13,17,19}
	//result := 1
	//var divisor int
	//for i:=range primes{
	//	divisor = primes[i]
	//	for divisor <=20{
	//		divisor *= primes[i]
	//	}
	//	divisor /=primes[i]
	//	result *=divisor
	//}
	//println(result)

	primes := []int{2,3,5,7,11,13,17,19}
	var a [20]float64
	k := 20.0
	N := 1
	check := true
	limit := math.Sqrt(k)
	for i:= range primes{
		a[i] =1
		if check {
			if float64(primes[i]) <= limit{
				a[i] = math.Floor(math.Log(k) /math.Log(float64(primes[i])))
			} else {
				check = false
			}
 		}
		N = N* primes[i] ^ (int(a[i]))
	}
	println(N)
}
