package main

import (
	"crypto/sha256"
	"fmt"
)

/*func badBadHash(data []type) uint64{
	var hash uint64 

	for _, b := range data  {
		hash += uint64(b)
		
	}
	return hash
}*/


func main (){
	 const input = "lemon"
	 hash := sha256.New()
	 hash.write([]byte(input))
	 fmt.Printf("%x\n", hash.sum(nil))
	// fmt.Printlf("Hash [\"%s\"]: %d\n", input)
	// input +=  "s" 
	// fmt.Printf("Hash [\"%s\"]: %d\n", input)
}