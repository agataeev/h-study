package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// WorkerResult хранит промежуточную сумму, вычисленную горутиной.
type WorkerResult struct {
	sum float64
}

func calculateLeibnizPart(ctx context.Context, start int, step int, resultChan chan<- WorkerResult, wg *sync.WaitGroup) {
	defer wg.Done()

	var sum float64
	sign := 1.0
	for i := start; ; i += step {
		select {
		case <-ctx.Done():
			// Передаем результат перед завершением.
			resultChan <- WorkerResult{sum: sum}
			return
		default:
			// Используем формулу ряда Лейбница.
			sum += sign / (2.0*float64(i) + 1.0)
			sign = -sign // Чередование знака.
		}
	}
}

func main() {
	// Обрабатываем аргументы запуска.
	numWorkers := flag.Int("n", 1, "number of goroutines")
	flag.Parse()

	if *numWorkers <= 0 {
		fmt.Println("Number of goroutines must be greater than zero")
		return
	}

	// Настраиваем обработку сигналов.
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Канал для передачи промежуточных результатов.
	resultChan := make(chan WorkerResult, *numWorkers)

	// Запускаем горутины.
	var wg sync.WaitGroup
	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go calculateLeibnizPart(ctx, i, *numWorkers, resultChan, &wg)
	}

	// Обработка сигналов остановки.
	go func() {
		<-signalChan
		fmt.Println("\nReceived termination signal. Calculating final result...")
		cancel() // Уведомляем горутины о завершении.
	}()

	// Суммируем промежуточные результаты.
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var totalSum float64
	for result := range resultChan {
		totalSum += result.sum
	}

	// Умножаем итог на 4 для получения значения Пи.
	pi := 4 * totalSum
	fmt.Printf("Approximated value of Pi: %.15f\n", pi)
}
