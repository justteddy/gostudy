package main

import (
	"fmt"
)

func look(player *Person) string {
	return player.place.look()
}

func move(player *Person, place string) string {
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

// func take(player *Person, elm string) string {
// 	/* проверка на наличие рюкзака */
// 	// return player.place.take(elm)
// 	return "not realized yet"
// }

// func wear(player *Person, wear *Wearable) {

// }

// func put(player *Person, stuff *Usable) {

// }

// func use(player *Person, stuff *Usable, env string) {

// }
