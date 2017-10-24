package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//wolrdmap
var worldmap map[Placable][]Placable
var placename map[Placable]string

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
		return look()
	}

	if len(parsedCmd) == 2 && parsedCmd[0] == "идти" {
		return move(parsedCmd[1])
	}

	if len(parsedCmd) == 2 && parsedCmd[0] == "одеть" {
		return wear(parsedCmd[1])
	}

	if len(parsedCmd) == 2 && parsedCmd[0] == "взять" {
		return put(parsedCmd[1])
	}

	return "неизвестная команда"

}

func initGame() {
	var playername string
	fmt.Println("Назови свое имя и начнем:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		playername = scanner.Text()
	}

	kitchen = Kitchen{
		stuff: []string{"чай"},
	}

	room = Room{
		stuff: []string{"ключи", "конспекты"},
		wears: []string{"рюкзак"},
	}

	outdoor = Outdoor{
		stuff: []string{},
	}

	hallway = Hallway{
		stuff: []string{},
	}

	player = Person{
		name:  playername,
		place: &room,
	}

	worldmap = map[Placable][]Placable{
		&kitchen: []Placable{&hallway},
		&hallway: []Placable{&kitchen, &outdoor, &room},
		&room:    []Placable{&hallway},
		&outdoor: []Placable{&hallway},
	}

	placename = map[Placable]string{
		&kitchen: "кухня",
		&hallway: "коридор",
		&room:    "комната",
		&outdoor: "улица",
	}

	showHelp()
}

func showHelp() {
	fmt.Println("------------------------------------")
	fmt.Printf("Привет, %s! Давай начнем. Вот какие команды тебе доступны:\n", player.name)
	fmt.Println("осмотреться - что происходит вокруг тебя")
	fmt.Println("взять <предмет> - взять <предмет>")
	fmt.Println("одеть <вещь> - одеть на себя <вещь>")
	fmt.Println("идти <место> - пойти в <место>")
	fmt.Println("применить <предмет> <цель> - применить <предмет> на определенную <цель>")
}
