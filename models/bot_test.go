package models

import (
	"log"
	"testing"
)

type BotTest struct {
	Command string
	Want    Bot
}

var botTests = []BotTest{
	{"W5RW5RW2RW1R", Bot{X: 4, Y: 3, Direction: North}},
	{"RRW11RLLW19RRW12LW1", Bot{X: 7, Y: -12, Direction: South}},
	{"LLW100W50RW200W10", Bot{X: -210, Y: -150, Direction: West}},
	{"LLLLLW99RRRRRW88LLLRL", Bot{X: -99, Y: 88, Direction: East}},
	{"W55555RW555555W444444W1", Bot{X: 1000000, Y: 55555, Direction: East}},
}

func TestBot(t *testing.T) {
	for _, test := range botTests {
		bot := new(Bot)
		_, err := bot.SetCommand(test.Command)
		if err != nil {
			log.Fatal(err)
		}
		if *bot != test.Want {
			t.Errorf("SetCommand(%s) = %+v; want %+v", test.Command, *bot, test.Want)
		}
	}
}
