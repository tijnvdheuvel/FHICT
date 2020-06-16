package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Quotes met Jan... packages: "De meeste wielen zijn al uitgevonden, je moet alleen weten in welke garage ze staan."
func main() {
	Pass := ""
	input:=""
	PassLength :=20
	soort := "woord" //tekst of woord
	b := rand.Intn(10)
// de startup sequence
	fmt.Println("Beste gebruiker, welkom bij de wachtwoord generator.")
	fmt.Println("U gaat nu enkele vragen krijgen t.b.v. het genereren van een wachtwoord.")
	fmt.Println("Hoeveel tekens wilt u dat uw wachtwoord lang is? (integer) default: 10")
	scanner:= bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	if input == ""{
		PassLength = 10
		fmt.Println("U hebt gekozen voor de standaard lengte van 10.")
	} else {
		a, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		PassLength = a
		fmt.Println("Uw Input was ", input, ".")
	}
		fmt.Println("Welke type wachtwoord generatie zou u willen? 1: woorddelen, 2: willekeurige tekens (1/2) default: 1")
		scanner.Scan()
		input = scanner.Text()
		if input != "2"{
			fmt.Println("U hebt gekozen voor type 1, woorddelen.")
			soort = "woord"
		} else {
			fmt.Println("U hebt gekozen voor type 2, willekeurige tekens.")
			soort = "tekst"
		}
		fmt.Println("Uw wachtwoord wordt nu gegenereerd.")

		fmt.Println("============================================")
// Generate loop.
	for len(Pass) < PassLength {

		b = rand.Intn(10)

		if soort == "woord" {
			Pass = string(woordGenerator(b, PassLength, Pass))
		} else if soort == "tekst" {
			Pass = string(tekstGenerator(b, Pass))
		}
	}
fmt.Println("Uw gegenereerde wachtwoord is:")
	fmt.Println(Pass)
}

func woordGenerator(b, PassLength int, Pass string) string {
	//Ik heb hiervoor gekozen om een generator woorddelen uit een array te laten pakken,
	// Hierdoor zouden wachtwoorden makkelijker onthoudbaar moeten zijn.
	b = rand.Intn(10)
	g :=newNumber(10,b)
	h :=newNumber(5, b)
		if len(Pass+ beginStuk(g)) <PassLength && len(Pass+ beginStuk(g)) < int(math.Round(float64(PassLength/2))){
		Pass = Pass + beginStuk(g)
	} else if len(Pass+ tussenStuk(h)) <PassLength{
		Pass = Pass + tussenStuk(h)
	} else {
		Pass =tekstGenerator(b, Pass)
		}
	return Pass
}

func tekstGenerator(b int, Pass string) string {

	c := newNumber(126, b)
	if c == 39 || c == 34 || c < 33{
		c = newNumber(126, b)
	} else {
		Pass += string(c)
	}
	return Pass

}

func newNumber(einde, b int) int {
	time.Sleep(time.Duration(b) * time.Millisecond)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	uitkomst := r1.Intn(einde)
	return uitkomst

}