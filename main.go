package main

import (
	"fmt"
	"time"
)

func main() {

	data := make(chan string)
	
	go sendData(data)
	go getData(data)

//  contoh pembuatan go routine langsung dari go routine main
	go func (in int)  {
		for i := 0; i < in ; i ++ {
			fmt.Println(i)
		} 
	}(2)


// kenapa ada time sleep ? supaya ga keburu mati program nya jadi nunggu dulu kalo dah beres baru selesai
	time.Sleep(time.Second)
}



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