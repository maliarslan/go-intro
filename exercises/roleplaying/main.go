// Starting code for the exercise
package main

import (
	"fmt"
	"slices"
)

type ItemType int

const (
	Potion ItemType = iota + 1
	Mage
)

// Item struct
type Item struct {
	Name string
	Type ItemType
}

// Player struct
type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) ItemExist(itemName string) bool {
	return slices.ContainsFunc(p.Inventory, func(item Item) bool {
		return item.Name == itemName
	})
}

// Pick up an item (modifies the player's inventory)
func (p *Player) PickUpItem(item Item) {
	if p.ItemExist(item.Name) {
		fmt.Printf("%v has already %v in their inventory!\n", p.Name, item.Name)
		return
	}

	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%v is succesfully added to %v's inventory\n", item.Name, p.Name)
	fmt.Printf("New inventory: %+v\n", p.Inventory)

}

// Drop an item (removes from inventory)
func (p *Player) DropItem(item Item) {
	if p.ItemExist(item.Name) {
		newInventory := []Item{}
		for _, value := range p.Inventory {
			if value.Name == item.Name {
				continue
			}
			newInventory = append(newInventory, value)
		}
		p.Inventory = newInventory
		fmt.Printf("%v has dropped %v\n", p.Name, item.Name)
		fmt.Printf("New inventory: %+v\n", p.Inventory)
		return
	}

	fmt.Printf("%v doesn't have %v\n", p.Name, item.Name)
}

// Use an item (if potion, remove it after use)
func (p *Player) UseItem(item Item) {
	if p.ItemExist(item.Name) {
		if item.Type == Potion {
			newInventory := []Item{}
			for _, value := range p.Inventory {
				if value.Name == item.Name {
					continue
				}
				newInventory = append(newInventory, value)
			}
			p.Inventory = newInventory
			fmt.Printf("Potion is removed from %v's inventory: %+v\n", p.Name, p.Inventory)
		}

		fmt.Printf("%v has used item %v\n", p.Name, item.Name)
		return
	}

	fmt.Printf("%v doesn't have %v\n", p.Name, item.Name)
}

func main() {
	liandry := Item{
		Name: "Liandry's Torment",
		Type: Mage,
	}
	rylai := Item{
		Name: "Rylai's Crystal Scepter",
		Type: Mage,
	}
	hp := Item{
		Name: "Health Potion",
		Type: Potion,
	}
	mp := Item{
		Name: "Mana Potion",
		Type: Potion,
	}
	player := Player{
		Name: "Stronk",
		Inventory: []Item{
			liandry,
			rylai,
			hp,
		},
	}

	fmt.Printf("New player joined! Welcome %v\n", player.Name)
	fmt.Printf("Players inventory: %+v\n", player.Inventory)

	player.PickUpItem(liandry)
	player.PickUpItem(mp)
	player.DropItem(rylai)
	player.UseItem(hp)
}
