package main

import (
	"fmt"
	"time"
)

// or принимает один или более каналов и возвращает канал,
// который закроется, когда закроется любой из входных каналов.
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если нет входных каналов, возвращаем закрытый канал.
		closedChan := make(chan interface{})
		close(closedChan)
		return closedChan
	case 1:
		// Если только один канал, возвращаем его.
		return channels[0]
	}

	// Рекурсивно объединяем каналы.
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		select {
		case <-channels[0]:
		case <-or(append(channels[1:])...): // Рекурсивный вызов
		}
	}()
	return orDone
}

func main() {
	// Тестовая функция для создания канала, который закроется через заданное время.
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// Тестирование функции or
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start)) // ~1 second
}
