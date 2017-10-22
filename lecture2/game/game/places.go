package main

import (
	"fmt"
	"strings"
)

type Room struct {
	stuffs []string
	wears  []string
}

func (r *Room) look() {

}

type Kitchen struct {
	stuffs []string
}

func (k *Kitchen) look() string {
	var stuffHere string
	if len(k.stuffs) == 0 {
		stuffHere = "ничего нет"
	} else {
		stuffHere = strings.Join(k.stuffs, ",")
	}

	return fmt.Sprintf("ты находишься на кухне, на столе %s, надо идти в универ. можно пройти - коридор", stuffHere)
}

func (k *Kitchen) take(elm string) string {
	for i, val := range k.stuffs {
		if val == elm {
			k.stuffs = append(k.stuffs[:i], k.stuffs[i+1:]...)
			return fmt.Sprintf("предмет добавлен в инвентарь:%s", elm)
		}
	}

	return "нет такого"
}

type Hallway struct {
	stuffs []string
}

func (h Hallway) look() {

}

type Outdoor struct {
	stuffs []string
}

func (o Outdoor) look() {

}
