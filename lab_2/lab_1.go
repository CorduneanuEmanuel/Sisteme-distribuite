package main

import (
	"fmt"
	"strconv"
	"sync"
	// "sync"
)

func exercitiul1(ch chan string) {
	ch <- "salut"

}

func exercitiul2(ch chan string) {
	fmt.Println(<-ch)

}

func send(ch chan<- string) {
	ch <- "Exercitiul 5"
}

func receive(ch <-chan string, valoare *string, wg *sync.WaitGroup) {
	*valoare = <-ch
	wg.Done()
}

func main() {
	// Exercitiul 1

	// var ch = make(chan string)

	// go exercitiul1(ch)

	// fmt.Println(<-ch)

	// Exercitiul 2

	// var ch = make(chan string)
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func(ch chan string, wg *sync.WaitGroup){
	// 	ch <- "Exercitiul 2"
	// 	wg.Done()
	// }(ch, &wg)

	// fmt.Println(<-ch)

	// wg.Wait()

	// Exercitiul3

	// 	var ch = make(chan string, 2)

	// 	go func(ch chan string){
	// 		ch <- "Prim mesaj"
	// 		ch <- "Doi mesaj"
	// 	}(ch)

	// 	fmt.Println(<-ch)
	// 	fmt.Println(<-ch)

	// Exercitiul 4

	// var ch = make(chan string)
	// var ok = make(chan bool)

	// go func(ch chan string){
	// 	fmt.Println(<- ch)
	// 	ok <- true
	// }(ch)

	// ch <- "Exercitiul 4"

	// <- ok

	// Exercitiul 5

	// var ch = make(chan string)
	// var val string
	// var wg sync.WaitGroup
	// // var ok = make(chan bool)
	// wg.Add(1)
	// go receive(ch, &val, &wg)
	// go send(ch)

	// wg.Wait()
	// fmt.Println(val)

	//Exercitiul 6

	// var ch1 = make(chan string)
	// var ch2 = make(chan string)

	// var wg sync.WaitGroup

	// // wg.Add(2)
	// go func(ch1 chan string, wg *sync.WaitGroup){

	// 	ch1 <- "Primul robot"
	// }(ch1, &wg)
	// go func(ch2 chan string, wg *sync.WaitGroup){
	// 	ch2 <- "Al doilea robot"
	// }(ch2, &wg)

	// select{
	// 	case msg := <- ch1:
	// 		fmt.Println(msg)
	// 	case msg := <- ch2:
	// 		fmt.Println(msg)
	// }

	// Exercitiul 7

	var ch1 = make(chan string, 30)
	var ch2 = make(chan string, 30)

	var wg sync.WaitGroup

	wg.Add(2)
	go func(ch1 chan string, ch2 chan string, wg *sync.WaitGroup) {

		go func() {
			for x := 0; x < 10; x++ {
				for y := 0; y < 10; y++ {
					ch1 <- "(" + strconv.Itoa(x*10) + ", " + strconv.Itoa(y*10) + ")"
				}
			}
			close(ch1)
			fmt.Println("Robotul 1 s-a inchis")

		}()

		for v := range ch2 {
			fmt.Println("Robotul 1 primeste: " + v)
		}

		wg.Done()
	}(ch1, ch2, &wg)

	go func(ch1 chan string, ch2 chan string, wg *sync.WaitGroup) {

		go func() {
			for x := 0; x < 10; x++ {
				for y := 0; y < 10; y++ {
					ch1 <- "(" + strconv.Itoa(200+x*10) + ", " + strconv.Itoa(200+y*10) + ")"
				}
			}
			close(ch1)
			fmt.Println("Robotul 2 s-a inchis")
		}()

		for pos := range ch2 {
			fmt.Println("Robotul 2 primeste: " + pos)
		}

		wg.Done()
	}(ch2, ch1, &wg)

	wg.Wait()

	fmt.Println("Ambii roboti au terminat")
}
