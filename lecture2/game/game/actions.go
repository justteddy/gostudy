package main

func look(player *Person) string {
	return player.place.look()
}

func take(player *Person, elm string) string {
	// res, ok := player.place.take(elm); !ok {
	// 	return res
	// }

	return player.place.take(elm)
}

// func move(player *Person, place *Placable) {

// }

// func wear(player *Person, wear *Wearable) {

// }

// func put(player *Person, stuff *Usable) {

// }

// func use(player *Person, stuff *Usable, env string) {

// }
