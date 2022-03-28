package main

import "fmt"

func main() {
	//ch := make(chan int) //unbuffered channel
	//ch := make(chan int, 1) //buffered channel (size 1)
	ch := make(chan int, 2) //buffered channel (size 2)
	go writeData(ch)

	fmt.Println("[@main] attempting to read 10")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 10")

	fmt.Println("[@main] attempting to read 20")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 20")

	fmt.Println("[@main] attempting to read 30")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 30")

	fmt.Println("[@main] attempting to read 40")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 40")

	fmt.Println("[@main] attempting to read 50")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 50")

	fmt.Println("[@main] attempting to read 60")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 60")

	fmt.Println("[@main] attempting to read 70")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 70")

	fmt.Println("[@main] attempting to read 80")
	fmt.Println(<-ch)
	fmt.Println("[@main] completed reading 80")

}

func writeData(ch chan int) {
	fmt.Println("	[@writeData] attempting to write 10")
	ch <- 10
	fmt.Println("	[@writeData] completed writing 10")

	fmt.Println("	[@writeData] attempting to write 20")
	ch <- 20
	fmt.Println("	[@writeData] completed writing 20")

	fmt.Println("	[@writeData] attempting to write 30")
	ch <- 30
	fmt.Println("	[@writeData] completed writing 30")

	fmt.Println("	[@writeData] attempting to write 40")
	ch <- 40
	fmt.Println("	[@writeData] completed writing 40")

	fmt.Println("	[@writeData] attempting to write 50")
	ch <- 50
	fmt.Println("	[@writeData] completed writing 50")

	fmt.Println("	[@writeData] attempting to write 60")
	ch <- 60
	fmt.Println("	[@writeData] completed writing 60")

	fmt.Println("	[@writeData] attempting to write 70")
	ch <- 70
	fmt.Println("	[@writeData] completed writing 70")

	fmt.Println("	[@writeData] attempting to write 80")
	ch <- 80
	fmt.Println("	[@writeData] completed writing 80")
}
