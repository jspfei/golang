package main

import (
	"sync"
)

type Dragon struct {
	name     string
	property string
}

func loadDragon(n, p string) Dragon {
	dragon := Dragon{
		name:     n,
		property: p,
	}

	return dragon
}

var cards map[string]Dragon

func loadCards() {
	cards = map[string]Dragon{
		"red":   loadDragon("红龙", "火"),
		"blue":  loadDragon("蓝色", "冰"),
		"white": loadDragon("白色", "光"),
		"black": loadDragon("黑龙", "暗"),
	}
}

var onlyOne sync.Once

func card(name string) Dragon {
	if cards == nil {
		onlyOne.Do(loadCards)
	}

	return cards[name]
}
