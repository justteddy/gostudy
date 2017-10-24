package main

import (
	"fmt"
	"log"
	"strings"
)

func accessableWays(place Placable) string {
	if _, ok := worldmap[place]; !ok {
		log.Fatal("Shit happens!")
	}

	ways := []string{}
	for _, variant := range worldmap[place] {
		ways = append(ways, placename[variant])
	}

	return strings.Join(ways, ", ")
}

func removeFromPlace(stuff *[]string, thing string) bool {
	slice := *stuff
	for i, st := range *stuff {
		if st == thing {
			*stuff = append(slice[:i], slice[i+1:]...)
			return true
		}
	}

	return false
}

type Room struct {
	stuff []string
	wears []string
}

func (r *Room) put(thing string) string {
	if !removeFromPlace(&r.stuff, thing) {
		return "нет такого"
	}

	return fmt.Sprintf("предмет добавлен в инвентарь: %s", thing)
}

func (r *Room) look() string {
	var stuff string
	var wear string

	if len(r.stuff) == 0 {
		stuff = "пустая комната"
	} else {
		stuff = fmt.Sprintf("на столе: %s", strings.Join(r.stuff, ", "))
	}

	if len(r.wears) == 0 {
		wear = ""
	} else {
		wear = fmt.Sprintf(", на стуле - %s", strings.Join(r.wears, ", "))
	}

	return fmt.Sprintf("%s%s. можно пройти - %s", stuff, wear, accessableWays(r))
}

func (r *Room) oncome() string {
	return fmt.Sprintf("ты в своей комнате. можно пройти - %s", accessableWays(r))
}

type Kitchen struct {
	stuff []string
}

func (k *Kitchen) put(thing string) string {
	return "not realized"
}

func (k *Kitchen) look() string {
	var stuff string
	if len(k.stuff) == 0 {
		stuff = "ничего нет"
	} else {
		stuff = strings.Join(k.stuff, ", ")
	}

	return fmt.Sprintf("ты находишься на кухне, на столе %s, надо идти в универ. можно пройти - %s", stuff, accessableWays(k))
}

func (k *Kitchen) oncome() string {
	return fmt.Sprintf("кухня, ничего интересного. можно пройти - %s", accessableWays(k))
}

type Hallway struct {
	stuff []string
}

func (h *Hallway) put(thing string) string {
	return "not realized"
}

func (h *Hallway) look() string {
	return fmt.Sprintf("ты в коридоре. можно пройти - %s", accessableWays(h))
}

func (h *Hallway) oncome() string {
	return fmt.Sprintf("ничего интересного. можно пройти - %s", accessableWays(h))
}

type Outdoor struct {
	stuff []string
}

func (o *Outdoor) put(thing string) string {
	return "not realized"
}

func (o *Outdoor) look() string {
	return fmt.Sprintf("ты на улице. можно пройти - %s", accessableWays(o))
}

func (o *Outdoor) oncome() string {
	return "на улице весна. можно пройти - домой"
}
