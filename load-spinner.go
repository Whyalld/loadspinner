package loadspinner

import (
	"fmt"
	"time"
)

func ShowLoading(stopChan chan bool) {
	symbols := `-\|/`
	delay := 100 * time.Millisecond

	for {
		for _, symb := range symbols {
			select {
			case <- stopChan:
				fmt.Println("\033[1m\rDone!           \033[0m")
				return
			default:
				fmt.Printf("\033[1m\r%c Loading...\033[0m", symb)
				time.Sleep(delay)
			}
		}
	}
}