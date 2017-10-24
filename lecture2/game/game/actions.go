package main

import (
	"fmt"
	"reflect"
)

func look() string {
	return player.place.look()
}

func move(place string) string {
	for destination, name := range placename {
		if name == place {
			for _, variant := range worldmap[player.place] {
				if variant == destination {
					player.place = variant
					return player.place.oncome()
				}
			}
			break
		}
	}

	return fmt.Sprintf("нет пути в %s", place)
}

func wear(wear string) string {
	if wear != "рюкзак" {
		return fmt.Sprintf("нельзя одеть %s", wear)
	}

	if reflect.TypeOf(player.place) != reflect.TypeOf(&Room{}) {
		return fmt.Sprintf("здесь нет %s", wear)
	}

	var t interface{} = player.place
	original, ok := t.(*Room)
	if !ok {
		return fmt.Sprintf("здесь нет %s", wear)
	}

	player.isBagWeared = true
	for i, thing := range original.wears {
		if thing == wear {
			original.wears = append(original.wears[:i], original.wears[i+1:]...)
			break
		}
	}

	return fmt.Sprintf("вы одели: %s", wear)
}

func put(thing string) string {
	if !player.isBagWeared {
		return "некуда класть"
	}

	return player.place.put(thing)
}

// func use(player *Person, stuff *Usable, env string) {

// }
