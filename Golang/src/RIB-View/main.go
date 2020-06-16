package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
	"time"
)

var input string

func main() {
	cmd := exec.Command("route", "print")
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	//Client interface start
	fmt.Println("RIB view gestart")
	fmt.Println("U gaat zometeen gevraagd worden voor een IP-adress. \nDe regels waarin dit IP adress voorkomt zullen met blauw gemarkeerd worden.")
	fmt.Println("Voor alstublieft het IP adress in. (format: [192.168.0.0])")
	inputscan := bufio.NewScanner(os.Stdin)
	inputscan.Scan()
	input = inputscan.Text()
	fmt.Println("Uw input was: ", input)
	fmt.Println("Programma start \n****************************")

// Start het programma
	cmd.Start()

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() { //Check pipe output loop.
		tekst := scanner.Text()

		if strings.Contains(tekst, input){
			color.Blue(tekst) //Print de line met de kleur blauw.
		} else {
			fmt.Println(tekst)
		}
		}


	cmd.Wait() //Wait until command finishes
fmt.Println("*******************\n Het programma sluit over 120 seconden.")
	duration := time.Duration(120)*time.Second
	time.Sleep(duration)

}