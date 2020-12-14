package main

import (
	"fmt"
	"time"
	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {
	input := ""
	for {
		fmt.Println("Pomodoro started at " + time.Now().Format("15:04"))

		start(25)

		fmt.Println("Work done! 5 minutes break")

		notify = notificator.New(notificator.Options{
			DefaultIcon: "",
			AppName:     "Pomodoro",
		})

		notify.Push("Work done", "Go rest for 5 minutes", "", notificator.UR_CRITICAL)

		start(5)

		fmt.Print("5 minutes passed. Press enter to start timer. Type exit to stop. ")
		notify.Push("Break finished", "Time to start a new pomodoro", "", notificator.UR_CRITICAL)

		fmt.Scanln(&input)
		if input == "exit" {
			return
		}
	}
}

func start(timerLength int) {
	ticker := time.NewTicker(1 * time.Minute)
	done := make(chan bool)
	minutesDone := 0
	printProgressBar(timerLength, minutesDone)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				minutesDone++
				printProgressBar(timerLength, minutesDone)
			}
		}
	}()

	time.Sleep(time.Duration(timerLength) * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println("")
}

func printProgressBar(total int, minutesDone int) {
	fmt.Print("\r")
	fmt.Print("[")
	for i := 0; i < total; i++ {
		if i < minutesDone {
			fmt.Print("X")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Print("]")
}
