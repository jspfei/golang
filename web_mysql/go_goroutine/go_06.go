package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sw sync.WaitGroup
var itemChan chan *item

var resultChan chan *result

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item

	sum int64
}

//生产者
func producer(ch chan *item) {
	var id int64
	for {
		id++
		number := rand.Int63()

		tmp := &item{
			id:  id,
			num: number,
		}

		ch <- tmp
	}
	close(ch)
}

func calc(num int64) int64 {
	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}

	return sum
}

func consumer(ch chan *item, resultChan chan *result) {
	defer sw.Done()
	for {
		tmp := <-ch
		sum := calc(tmp.num)

		resObj := &result{
			item: tmp,
			sum:  sum,
		}

		resultChan <- resObj
	}
}

func startWorker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

func printResult(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id:%v,num:%v,sum:%v\n", ret.item.id, ret.item.num, ret.sum)
	}
	time.Sleep(time.Second)
}

func main() {

	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 100)

	go producer(itemChan)
	sw.Add(30)
	startWorker(30, itemChan, resultChan)
	sw.Wait()
	close(resultChan)
	printResult(resultChan)

}
