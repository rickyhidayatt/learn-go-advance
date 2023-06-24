package main

import (
	"fmt"
	"sync"
)
func main(){

	data := make(chan int)
	err := make(chan string)
	
	wg := new(sync.WaitGroup)
	wg.Add(2)// menandakan ada 2 gorutine

	go func (w *sync.WaitGroup)  {
		defer w.Done()
		for i := 0 ; i < 10 ; i ++{
			data <- i
		}
		err <- "ada err"

	}(wg)

	go func (w *sync.WaitGroup)  {
		defer w.Done()
		for {
			select{

			case res := <- err:
				fmt.Println(res)
				return
				
			case datas := <- data:
				fmt.Println(datas)
				
			}
		}
	}(wg)
// tungguin prosesnya dengan menggunakan wg Wait
wg.Wait()

}


// func main() {

// 	data := make(chan string)

// 	go sendData(data)
// 	go getData(data)

// //  contoh pembuatan go routine langsung dari go routine main
// 	go func (in int)  {
// 		for i := 0; i < in ; i ++ {
// 			fmt.Println(i)
// 		}
// 	}(2)

// // kenapa ada time sleep ? supaya ga keburu mati program nya jadi nunggu dulu kalo dah beres baru selesai
// 	time.Sleep(time.Second)
// }

// untuk membuat go routine kita wajib menggunakan type data Chan
// karena sistem kerjanya gorutine kaya gini
// kirim data menggunakan go routine -> simpan di channel -> ambil data dari channel ke goroutine lain.



func sendData(send chan string) {
	send  <- "data from go routine" 
	// ( <- ) artinya kirim data gorutine type data string dan simpan di channel  
}



func getData(data chan string) {
get := <-data
// ( <- ) artinya simpan data dari gorutine type data string dan simpan data nya di variabel get   

fmt.Println(get)
}