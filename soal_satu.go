package main

import "fmt"

func main() {
    var arr [6]int
    var n, biggest, smallest int
    arr = [6]int {-5, 12, 11, -25, 2, 12345}
    
    for _,rgbig := range arr {
        if rgbig > n {
          n = rgbig
          biggest = n
        }
    }
    fmt.Println("Angka Terbesar: ", biggest)
    
    for _,rgsmall := range arr {
        if rgsmall < n {          
          n = rgsmall
          smallest = n
        } 
    }
    
    fmt.Println("Angka terkecil", smallest)
}