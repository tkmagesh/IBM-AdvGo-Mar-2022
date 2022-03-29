package main

import (
	"fmt"
	"log"
	"math/rand"
	"pool-demo/db"
	"pool-demo/pool"
	"sync"
	"time"
)

func main() {

	//create an Pool instance (with the pool size & factory)
	/*
		//When a resource is 'acquired'
		//the pool will check if it has any resources
		//if yes, then return from the pool
		//else create a resource using the factory and return
	*/
	/*
		//When a resource is 'released'
		//the pool will check if it is full
		//if yes, then discard the resource (after 'closing' the resource)
		//else, keep the resource to serve future clients (maintain the resource in the pool)
	*/
	/*
		//When 'close()' of the 'pool' is called
		// prevent anymore acquisition of the resources
		// make sure all the resources are 'closed' and discarded
	*/
	/*
		//Important
		//When a resource is given to a client (acquired),
		//the same resource should be given to another client until it is released
	*/

	p, err := pool.New(db.DBConnectionFactory, 5 /* pool size */)

	if err != nil {
		log.Fatalln(err)
	}
	wg := &sync.WaitGroup{}
	clientCount := 10
	wg.Add(clientCount)
	for client := 0; client < clientCount; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()
	// 5 resources should have been discarded and 5 resources should have been maintained in the pool
	fmt.Println("Second batch of operations")
	var input string
	fmt.Scanln(&input)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(client int) {
			doWork(client, p) //Resources should be returned from the pool when acquired
			wg.Done()
		}(i)
	}
	wg.Wait()
	p.Close()
}

func doWork(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Fatalln(err)
	}
	defer p.Release(conn)
	fmt.Printf("Worker : %d, Acquired %d:\n", id, conn.(*db.DBConnection).ID)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	fmt.Printf("Worker Done : %d, Releasing %d:\n", id, conn.(*db.DBConnection).ID)
}
