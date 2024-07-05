package belajargolanggoroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())
	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	timer := time.After(5 * time.Second)

	fmt.Println(time.Now())
	time := <-timer
	fmt.Println(time)
}
