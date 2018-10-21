package models

import (
	"fmt"
	"regexp"
	"strconv"
)

// Direction type
//go:generate stringer -type=Direction
type Direction uint

// Direction Enum
const (
	North Direction = iota
	East
	South
	West
)

// Bot model
type Bot struct {
	X         int64
	Y         int64
	Direction Direction
}

// ValidateCommand from input argument
func ValidateCommand(command string) (bool, error) {
	r, err := regexp.Compile(`^[A-Z0-9]*$`)
	if err != nil {
		return false, err
	}
	validate := r.MatchString(command)
	return validate, nil
}

// SetCommand to Bot
func (bot *Bot) SetCommand(command string) (string, error) {
	match, err := ValidateCommand(command)
	if err != nil || !match {
		return "Invalid command", err
	}

	r := regexp.MustCompile(`[A-Z][0-9]*`)
	cmds := r.FindAllString(command, -1)
	for _, cmd := range cmds {
		flag := cmd[:1]

		if len(cmd[1:]) > 0 {
			points, err := strconv.Atoi(cmd[1:])
			if err != nil {
				return "", err
			}
			bot.Walk(points)
		} else {
			bot.Rotate(flag)
		}
	}
	return fmt.Sprintf("X: %d Y: %d Direction: %s\n", bot.X, bot.Y, bot.Direction), nil
}

// Rotate 90 degree to N side
func (bot *Bot) Rotate(side string) {
	if side == "L" {
		bot.Direction--
	} else {
		bot.Direction++
	}
	bot.Direction += 4
	bot.Direction %= 4
}

// Walk straight for N points
func (bot *Bot) Walk(points int) {
	switch bot.Direction {
	case 0:
		bot.Y += int64(points)
	case 1:
		bot.X += int64(points)
	case 2:
		bot.Y -= int64(points)
	case 3:
		bot.X -= int64(points)
	}
}
