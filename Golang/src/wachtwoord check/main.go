package main
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)
func main(){
	var Pass string = "P@ssw0rd"
	var score int =0
	//Controle parameters
	var minLength int = 10
	var minUpper int =0
	var minNumb int =0
	var DefaultMode bool = true
	//Default mode zorgt dat er minimaal: 1 Upper, 1 getal, minLength 7
	var input string =""

	var uitkomst string = ""
	fmt.Println("Beste gebruiker, welk wachtwoord wilt u testen?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Pass = scanner.Text()
	fmt.Println("Uw input was: ", Pass)
	fmt.Println("Wilt u de standaard controle waardes gebruiken? (Y/n) default: y")
	//scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
if input != "n" && input != "N" {
	DefaultMode = true
	fmt.Println("U hebt gekozen voor de standaard controle waardes.")
} else {
	DefaultMode = false
	fmt.Println("U hebt gekozen voor eigen controle waardes.")
	fmt.Println("U wordt zometeen verzoekt om waardes in te vullen, \ndeze waardes hebben een invloed op de eind evaluatie ")
	fmt.Println("Wat is het minimum aantal tekens dat u wilt toestaan? (integer) default: 7")
	scanner.Scan()
	input = scanner.Text()
	if input != ""{
		a, err := strconv.Atoi(input)
		if err !=nil {
			panic(err)
		}
		minLength=a
		fmt.Println("U hebt gekozen voor een minimum aantal tekens van ", input)

} else {
minLength = 7
fmt.Println("U hebt gekozen voor een minimum aantal tekens van 7")
}
	fmt.Println("Wat is het minimum aantal cijfers dat u wilt toestaan? (integer) default: 1")
	scanner.Scan()
	input = scanner.Text()
	if input != ""{
		b, err := strconv.Atoi(input)
		if err !=nil{
			panic(err)
		}
		minNumb=b
		fmt.Println("U hebt gekozen voor een minimum aantal cijfers van ", input)

	} else {
		minNumb = 1
		fmt.Println("U hebt gekozen voor een minimum aantal tekens van 1")
	}
	fmt.Println("Wat is het minimum aantal Uppercase letters dat u wilt toestaan? (integer) default: 1")
	scanner.Scan()
	input = scanner.Text()
	if input != ""{
		c, err := strconv.Atoi(input)
		if err !=nil{
			panic(err)
		}
		minUpper=c
		fmt.Println("U hebt gekozen voor een minimum aantal Uppercase letters van ", input)

	} else {
		minUpper = 1
		fmt.Println("U hebt gekozen voor een minimum aantal Uppercase letters van 1")
	}

}
	fmt.Println("Uw input was: ", Pass)

    if DefaultMode {
    	minLength := 7
    	minNumb :=1
    	minUpper :=1
		uitkomst, score= passControle(Pass, minLength, minUpper, minNumb)

	} else {
		uitkomst, score = passControle(Pass, minLength, minUpper, minNumb)
		///Sufficient = passControle(Pass, minLength, minSpec, minUpper, minNumb)

	}
	fmt.Println("Het ingevoerde wachtwoord is beoordeeld met een score van: ", score,".")
	fmt.Println("Volgens de ingestelde parameters was dit wachtwoord: ", uitkomst, ".")


}

func passControle(Pass string, minLength, minUpper, minNumb int) (string, int) {
	///func passControle(Pass string, minLength, minSpec, minUpper, minNumb int) (bool) {
	//Check hoeveel van welk type tekens er in Pass ziten

	uitkomst :=""
var waarde int=0
bonus :=0.0
//rating := ""
	e := 0
	f := 0
var multiplier float64 =0
	g := 0
	//s := []string{Pass}
	for i, r := range Pass {
		fmt.Sprint(i, r)
		if unicode.IsDigit(r) {
			e++
		} else if unicode.IsUpper(r) {
			f++
		} else if unicode.IsLower(r) {
			g++
		}
	}
	///if (e >= minNumb && f >= minUpper && g >= minSpec && len(Pass)>= minLength) {
	 if e >= 1 {
		 multiplier +=1.0
		 //Aangezien er 10 characters bij komen
	 }
	 if f >= 1{
	 	multiplier +=2.6
	 	//Aangezien er 26 character opties zijn toegevoegd voor een bruteforce.
	 }
	 if g >=1{
	 	multiplier +=2.6
		 //Aangezien er 26 character opties zijn toegevoegd voor een bruteforce.
	 }


	if e >= minNumb && f >= minUpper && len(Pass)>= minLength {
		bonus =1.5

		uitkomst = "voldoende"
		/// } else if !(e >= minNumb && f >= minUpper && g >= minSpec && len(Pass)>= minLength) {
	} else if !(e >= minNumb && f >= minUpper && len(Pass)>= minLength) {
		uitkomst = "onvoldoende"
		bonus = 1
	}

	c := (math.Pow(float64(len(Pass)), multiplier) * bonus)
	waarde = int(math.Round(math.Log10(c)))
//Ik heb hier gebruik gemaakt van een log functie aangezien ik een exponentieel groter wordende waarde moest indammen.
	return uitkomst, waarde
}
