//go:build app1

package main

import (
	"fmt"
	"sync"
)




type Msg struct {
	Value int
	Done bool 
}
func node(id int , in <-chan Msg , out chan <- Msg , isLeader *bool , wg *sync.WaitGroup ) {
	defer wg.Done()
	
	
	out <- Msg{Value: id, Done: false}

	for{
		mesaj := <-in

		if mesaj.Done {
			*isLeader = (mesaj.Value == id)


			out <- mesaj
			return
		}

		if mesaj.Value == id {

			*isLeader = true
			out <- Msg{Value: id, Done: true}
			return
		}

		if mesaj.Value > id {
			out <- mesaj
		}

	}



}
func RingFlooding (ids [] int) int { 
	n := len(ids)
	var wg sync.WaitGroup
	canale := make([]chan Msg, n)
	lideri := make([]bool, n)


	for i := 0; i < n; i++{
		canale[i] = make(chan Msg, n + 1)
	}

	for i:=0; i < n; i++{
		wg.Add(1)
		go node(ids[i], canale[i], canale[(i + 1) % n], &lideri[i], &wg)
	}

	wg.Wait()

	for i := 0; i < n; i++ {
		if lideri[i] {
			return ids[i]
		}
	}
	return 0

}
func main(){
	
	
	ids := []int{3, 7, 2, 9, 5}
	fmt.Println(RingFlooding(ids))

		



	
}