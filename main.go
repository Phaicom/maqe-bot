package main

import (
	"fmt"
	"log"
	"os"

	"github.com/phaicom/maqe-bot/models"
)

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		bot := new(models.Bot)
		result, err := bot.SetCommand(command)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
