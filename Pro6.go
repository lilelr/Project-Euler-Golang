package main

func main() {
	suma:=0
	sumb :=0
	for i:=1;i<=100;i++{
		suma += (i*i)
		sumb += i
	}
	println(sumb)
	sumb = (sumb * sumb)
	println(sumb)
	println(suma)
	println(sumb - suma)
}
