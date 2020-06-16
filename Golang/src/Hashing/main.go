package main

import (
	"crypto/sha256"
	//	"crypto"
	"fmt"
)

func main() {
	input := "7A-15-91-01-BC-FU\n"
	uitkomst := sha256.Sum256([]byte(input))

	tussen := fmt.Sprintf("%x", uitkomst)
	fmt.Println("\n Je input was: ", input)
	fmt.Println("De uitkomst in de variabele is: ", tussen)
test()
}

func test(){
	start := "91-21-AB-80-CD-LP"
	 hash := Hash(start)
	fmt.Println("Je functie was: ",hash)
}

func Hash(macadres string) string {
	uitkomst := sha256.Sum256([]byte(macadres))

	tussen := fmt.Sprintf("%x", uitkomst)
	return tussen
}
