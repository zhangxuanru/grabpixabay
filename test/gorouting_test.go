/*
@Time : 2020/6/4 17:30
@Author : zxr
@File : gorouting_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestGo(t *testing.T) {
	for i := 0; i < 10; i++ {
		ch := make(chan int)
		go func(ch chan int) {
			fmt.Println(<-ch)
		}(ch)
		ch <- i
	}
}

func TestGo1(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go printOdd(ch, &wg)
	go printEvent(ch, &wg)
	wg.Wait()
}

func printOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i += 2 {
		fmt.Println(i)
		ch <- i
		<-ch
	}
}

func printEvent(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println(i)
		ch <- i
	}
}
