// Below code is the example of channel syncronisation
package main

import (
	"fmt"
	"sync"
)

func main() {

	// var ch1 chan string = make(chan string)
	// // var ch2 chan string  //trying this bu this is not correct

	// // make (chan string) likhna jaruri h kyunki ye help krta h memory allocate krne me
	// // <- this symbol is used to pass any value in any variable

	// go func (){
	// 	ch1<-"hello world 'YE WALA CHAL GYA'"
	// 	// ch2<-"Ye wala nhi chalega kyunki memory hi allocate nhi hui h"
		
	// }()
	// fmt.Printf("ch1 = %v\n ", <-ch1)
	// // fmt.Printf("ch1 = %v\n ch2 = %v\n", <-ch1, <-ch2)
	// // print statement me <- symbol value nikaalne ke liye use hota h
	// YE SAARA EK BASIC STRUCTURE THA CHANNEL KA




	// MAIN CODE STARTING FROM HERE

	// channel is used for synchronisation means after completing one task next task will be done
	var ch chan int=make(chan int)
	wg:= sync.WaitGroup{}
	wg.Add(1)
	go sendNumber(ch)
	go recieveNumber(ch, &wg)
	wg.Wait()



	// CHANNEL BUFFERING
	ch2:=make(chan int,5)
	// here 5 represents that it can store 5 values at a time
	go func(){
		ch2<-1
		ch2<-2
		ch2<-3
		ch2<-4
		ch2<-5
	}()

	for val:=range ch2{
	fmt.Printf("buffering channel wala chl rha h, val = %v\n", val)
	}

}

func sendNumber(ch chan <- int) {
	// in above line we use <- symbol so that we can only send value, we can not recieve value from this func now
	for i := 0; i < 11; i++ {
		ch<-i
        
	}
	close (ch) // close is used to close the channel
}

func recieveNumber(ph <- chan int, wg *sync.WaitGroup) {
	// in above line we use <- symbol so that we can only recieve value, we can not send value from this func now
	defer wg.Done()
	for val:= range ph {
		fmt.Println(val)
	}
	
}
