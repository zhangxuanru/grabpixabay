/*
@Time : 2020/6/11 11:54
@Author : zxr
@File : go_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"testing"
)

func Test_Go(t *testing.T) {
	c := make(chan bool)
	b := make(chan string)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("true")
			case <-b:
				fmt.Println("false")
			}
		}
	}()
	fmt.Println("...........")
}
