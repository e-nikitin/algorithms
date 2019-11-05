package main

import "fmt"

func main() {
  a := [...]int{9,1,8,3,7,3,6,4,5,2,0}
  
  fmt.Println(a)
  
  for {
		b := true
		for i:=1; i<len(a);i++{
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
				b = false
				fmt.Printf("\t\tSWAP: %d <-- %d\n", a[i], a[i-1])
			}
		}
    
		fmt.Printf("\t%v\n", a)
    
		if b {break}
	}
  
  fmt.Println(a)

}
