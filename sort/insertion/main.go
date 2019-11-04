package main

import "fmt"

func main() {
  
	fmt.Println("From right to left")

	d := [...]int{2,0,9,5,3,1,6,8,4,7}
	fmt.Println(d)

	for i := len(d)-2; i >= 0; i-- {
		for j := i; j < len(d)-1 && d[j] > d[j+1]; j++ {
			d[j],d[j+1] = d[j+1], d[j]
			fmt.Printf("\t\tSWAP:%d <-- %d\n", d[j], d[j+1])
		}
		fmt.Printf("\t%v\n", d)
	}

	fmt.Println("--------------------------------------------------")

	fmt.Println("From left to right")

	a := [...]int{9,5,1,2,0,3,6,7,8}
	fmt.Println(a)

	for i:=1; i<len(a);i++  {
		for j:=i-1; j >= 0 && a[j] > a[j+1]; j-- {
			a[j], a[j+1] = a[j+1], a[j]
			fmt.Printf("\t\tSWAP:%d <-- %d\n", a[j+1], a[j])
		}
		fmt.Printf("\t%v\n", a)
	}


	fmt.Println(a)
}
