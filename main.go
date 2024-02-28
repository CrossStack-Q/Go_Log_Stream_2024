package main

import (
	"sync"

	"github.com/CrossStack-Q/Go_Log_Stream_2024/pkg/handler"
	"github.com/CrossStack-Q/Go_Log_Stream_2024/pkg/models"
)

func main() {
	ch := make(chan models.Event)
	var wg sync.WaitGroup

	handler.Init(ch, &wg)
	ch <- models.SystemLog{
		Log:      models.Log{ID: 1, Source: "Sytem1", Body: "System1 is there"},
		Severity: "Severoty info",
	}
	close(ch)

	wg.Wait()
}
