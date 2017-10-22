package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// places
var kitchen Kitchen
var room Room
var outdoor Outdoor
var hallway Hallway

// player
var player Person

func main() {
	initGame()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(handleCommand(scanner.Text()))
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

func handleCommand(command string) string {
	parsedCmd := strings.Split(command, " ")
	if len(parsedCmd) == 0 {
		return "Напишите, что хотите сделать!"
	}

	if len(parsedCmd) == 1 && parsedCmd[0] == "осмотреться" {
		return look(&player)
	}

	if len(parsedCmd) == 2 && parsedCmd[0] == "взять" {
		return take(&player, parsedCmd[1])
	}

	return "123"

}

func initGame() {
	var playername string

	fmt.Println("Назови свое имя и начнем:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		playername = scanner.Text()
	}

	kitchen = Kitchen{
		stuffs: []string{"чай"},
	}

	room = Room{
		stuffs: []string{"конспекты", "ключи"},
		wears:  []string{"рюкзак"},
	}

	outdoor = Outdoor{
		stuffs: []string{},
	}

	hallway = Hallway{
		stuffs: []string{},
	}

	player = Person{
		name:  playername,
		place: &kitchen,
	}

	fmt.Println("------------------------------------")
	fmt.Printf("Привет, %s! Давай начнем. Вот какие команды тебе доступны:\n", player.name)
	fmt.Println("осмотреться - что происходит вокруг тебя")
	fmt.Println("взять <предмет> - взять <предмет>")
	fmt.Println("одеть <вещь> - одеть на себя <вещь>")
	fmt.Println("идти <место> - пойти в <место>")
	fmt.Println("применить <предмет> <цель> - применить <предмет> на определенную <цель>")
}
