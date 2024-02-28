package handler

import (
	"fmt"
	"sync"

	"github.com/CrossStack-Q/Go_Log_Stream_2024/config"
	"github.com/CrossStack-Q/Go_Log_Stream_2024/pkg/models"
)

func processEvent(i int, ch chan models.Event, wg *sync.WaitGroup) {
	for event := range ch {
		fmt.Println("Event Received ", event, "by handler ", i)
	}
	wg.Done()
}

func Init(ch chan models.Event, wg *sync.WaitGroup) {
	for i := 0; i < config.MAX_HANDLERS; i++ {
		wg.Add(1)
		go processEvent(i, ch, wg)
	}
}
