package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)

	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}

	fileWg.Add(2)
	go source("data1.dat", dataCh, fileWg)
	go source("data2.dat", dataCh, fileWg)

	processWg.Add(4)
	go splitter(dataCh, evenCh, oddCh, processWg)
	go sum(evenCh, evenSumCh, processWg)
	go sum(oddCh, oddSumCh, processWg)
	go merger("result.dat", evenSumCh, oddSumCh, processWg)

	fileWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Done")
}

func source(filename string, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if val, err := strconv.Atoi(text); err == nil {
			dataCh <- val
		}
	}
}

func splitter(dataCh, evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(evenCh)
	defer close(oddCh)
	for val := range dataCh {
		if val%2 == 0 {
			evenCh <- val
		} else {
			oddCh <- val
		}
	}
}

func sum(valCh chan int, sumCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for val := range valCh {
		sum += val
	}
	sumCh <- sum
}

func merger(filename string, evenSumCh, oddSumCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for i := 0; i < 2; i++ {
		select {
		case evenSum := <-evenSumCh:
			file.WriteString(fmt.Sprintf("Even total : %d\n", evenSum))
		case oddSum := <-oddSumCh:
			file.WriteString(fmt.Sprintf("Odd total : %d\n", oddSum))
		}
	}
}
