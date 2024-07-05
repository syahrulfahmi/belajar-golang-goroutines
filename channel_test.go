package belajargolanggoroutine_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Ini data dari channel"
		fmt.Println("Selesai mengirim data ke channel")
		time.Sleep(2 * time.Second)
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	channel <- "Ini data dari channel parameter"
	fmt.Println("Selesai mengirim data ke channel")
	time.Sleep(2 * time.Second)
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	channel <- "Ini data dari send only"
	fmt.Println("Selesai mengirim data ke channel")
	time.Sleep(2 * time.Second)
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // channel value must be less than buffer
	channel <- "ini value channel"
	channel <- "ini value channel"
	defer close(channel)

	fmt.Println(len(channel))
	fmt.Println(cap(channel))
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Ini iterasi ke" + strconv.Itoa(i)
		}
		fmt.Println("selesai")
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari Channel 1: ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari Channel 1: ", data)
			counter++
		default:
			println("Menunggu Data...") // default channel
		}
		if counter == 2 {
			break
		}
	}
}
