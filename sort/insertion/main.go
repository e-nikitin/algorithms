package main

import "fmt"

func main() {
  d := [...]int{2,0,9,5,3,1,6,8,4,7}
	//d := [...]int{2,1}
	//d := [...]int{0}
  
	fmt.Println(d)
  
	// From right to left
	for i := len(d)-2; i >= 0; i-- {
		for j := i; j < len(d)-1 && d[j] > d[j+1]; j++ {
			d[j],d[j+1] = d[j+1], d[j]
			fmt.Printf("SWAP:%d <-- %d\n", d[j], d[j+1])
		}
	}
  
	fmt.Println(d)
}
