package main

import (
	"math/rand"
	"time"
)

// Возможно не совсем правильно понял что здесь нужно сделать

const maxError = 30
const maxThread = 10

type Workers struct {
	error int
	f     func(in <-chan int, out chan<- int)
}

func worker(in <-chan int, out chan<- int) {
	var data int
	data = <-in
	ticker := time.NewTicker(2 * time.Second)
	println("Worker запущен!!!!", data)
	for {
		select {
		case <-ticker.C:
			{
				e := rand.Intn(10)
				println("do something", data, e)
				if e > 5 {
					out <- data
				} else {
					// Эмулируем ошибку
					out <- -1
				}

			}
		}
	}
}

// myFunc := []MyFunc{}

func main() {

	var wSlice []Workers
	wSlice = append(wSlice, Workers{10, worker}, Workers{15, worker}, Workers{5, worker})
	// wSlice := [][]Workers({10, worker})

	for i, x := range wSlice {
		countErr := 0
		println("Запуск воркера №", i+1, "Количество запусков", maxThread)
		rd := make(chan int)

		for j := 0; j < maxThread; j++ {
			wr := make(chan int)
			go x.f(wr, rd)
			wr <- i + 1
		}

		for countErr < x.error {
			select {
			case f := <-rd:
				{
					if f == -1 {
						countErr++
					}
					println("Получено", f, "обработчиком -", i, countErr, x.error)
				}
			}
		}
		println("Допущено ошибок при обработке", countErr, "обработчиком -", i+1, "Остановлен воркер ")
	}

}
