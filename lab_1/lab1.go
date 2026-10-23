package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strconv"
)

func main() {

	// scanner := bufio.NewScanner(os.Stdin)

	// scanner.Scan()

	// number, err := strconv.Atoi(scanner.Text())
	// if err != nil {
	// 	fmt.Println("Eroare la conversie")
	// }

	// if number < 2 {
	// 	fmt.Println("Nu este prim.")
	// } else {
	// 	for i := 2; i * i <= number; i++ {
	// 		if number % i == 0 {
	// 			fmt.Println("Nu este prim.")
	// 			return
	// 		}
	// 	}
	// 	fmt.Println("Este prim.")
	// }

	// 2

	// scanner := bufio.NewScanner(os.Stdin)

	// intSlice := []int{}
	// var number int;
	// var n int;
	// // var err error;
	// if scanner.Scan() {
	// 	n, _ = strconv.Atoi(scanner.Text())
	// }

	// for scanner.Scan() {
	// 	number, _ = strconv.Atoi(scanner.Text())
	// 	intSlice = append(intSlice, number)

	// }
	// // fmt.Println(n)
	// for i := 0; i < n - 1; i++ {
	// 	var min_i int = i

	// 	for j := i + 1; j < n; j++ {
	// 		if intSlice[j] < intSlice[min_i]{
	// 			min_i = j
	// 		}
	// 	}

	// 	c := intSlice[i]
	// 	intSlice[i] = intSlice[min_i]
	// 	intSlice[min_i] = c

	// }

	// for i := 0; i < n; i++{
	// 	fmt.Printf("%d ", intSlice[i])

	// }

	// Exercitiul 3

	// scanner := bufio.NewScanner(os.Stdin)
	// intSlice := []int{}
	// parSlice := []int{}

	// for scanner.Scan() {
	// 	number, _ := strconv.Atoi(scanner.Text())
	// 	intSlice = append(intSlice, number)
	// }

	// for _, val := range intSlice{
	// 	if val % 2 == 0{
	// 		parSlice = append(parSlice, val)

	// 	}

	// }

	// for _, val := range parSlice{
	// 	fmt.Printf("%d ", val)
	// }

	// Exercitiul 4

	// scanner := bufio.NewScanner(os.Stdin)
	// floatSlice := []float32{}
	// var index int
	// if scanner.Scan() {
	// 	index, _ = strconv.Atoi(scanner.Text())

	// }
	// for scanner.Scan(){
	// 	val, _ := strconv.ParseFloat(scanner.Text(), 32)

	// 	floatSlice = append(floatSlice, float32(val))

	// }

	// newSlice := floatSlice[: index]
	// newSlice = append(newSlice, floatSlice[index + 1: ]...)

	// for _, i := range newSlice{
	// 	fmt.Printf("%f ", i)

	// }

	// Exerctiul 5

	// scanner := bufio.NewScanner(os.Stdin)
	// slice := []string{}

	// for scanner.Scan() {
	// 	index := scanner.Text()
	// 	slice = append(slice, index)
	// }

	// n := len(slice)
	// for i := 0; i < len(slice) /2 ; i++{
	// 	c := slice[i]
	// 	slice[i] = slice[n - i - 1]
	// 	slice[n - i - 1] = c
	// }

	// for _, i := range slice{
	// 	fmt.Printf("%s ", i)
	// }

	// Exercitiul 6

	// var mymap map[string]int = make(map[string]int)
	// // var output map[int][]string = make(map[int][]string)

	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	str := scanner.Text()
	// 	if scanner.Scan() {
	// 		val, _ := strconv.Atoi(scanner.Text())
	// 		mymap[str] = val

	// 	}
	// }

	// var nume []string
	// var inaltimi []int

	// for i, j := range mymap {
	//     nume = append(nume, i)

	// 	inaltimi = append(inaltimi, j)
	// }

	// for i := 0; i < len(inaltimi); i++ {

	//     for j := i + 1; j < len(inaltimi); j++ {
	//         if inaltimi[i] < inaltimi[j] {
	//             c := inaltimi[i]
	// 			inaltimi[i] = inaltimi[j]
	// 			inaltimi[j] = c

	//             d := nume[i]
	// 			nume[i] = nume[j]
	// 			nume[j] = d

	//         }
	//     }
	// }

	// for _, v := range nume{
	// 	fmt.Printf("%s ", v)
	// }

	// Exercitiul 7

	// var mymap map[string]int = make(map[string]int)

	// file, err := os.Open("file.txt")
	// if err != nil{
	// 	fmt.Println(err)
	// }

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan(){
	// 	mymap[scanner.Text()]++

	// }

	// defer file.Close()

	// for i, j := range mymap{
	// 	fmt.Printf("%s: %d\n", i, j)
	// }

	// Exercitiul 8

	// var mymap map[string]int = make(map[string]int)
	// var output map[int][]string = make(map[int][]string)

	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	str := scanner.Text()
	// 	if scanner.Scan() {
	// 		val, _ := strconv.Atoi(scanner.Text())
	// 		mymap[str] = val

	// 	}
	// }

	// for i, j := range mymap {
	// 	output[j] = append(output[j], i)

	// }

	// for i, j := range output {
	// 	fmt.Printf("Deparatmentul %d: [", i)
	// 	for _, v := range j {
	// 		fmt.Printf("%s ", v)
	// 	}
	// 	fmt.Printf("]\n")
	// }

	// Exercitiul 9

	scanner := bufio.NewScanner(os.Stdin)
	var mymap map[int]int = make(map[int]int)
	// slice := []int{}
	var numar int
	if scanner.Scan() {
		numar, _ = strconv.Atoi(scanner.Text())
	}

	for scanner.Scan() {
		nr, _ := strconv.Atoi(scanner.Text())
		// slice = append(slice, nr)
		gasit := numar - nr
		i, ok := mymap[gasit]
		if ok {
			fmt.Printf("%d %d\n", i, nr)
		}
		mymap[nr] = 1

	}

}
