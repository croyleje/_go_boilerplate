package main

import (
	"fmt"
	"log"
	"os"
)

// Boilerplate selectable item list.
// TODO: Code is not fully implemented meant only as a reference.
type Item struct {
	Id       int
	Text     string
	Selected bool
}

type List struct {
	CursorPos int
	Items     []*Item
}

func InitializeList() *List {
	return &List{
		Items: make([]*Item, 0),
	}

}

func (l *List) AppendListItem(id int, text string) *Item {
	Item := &Item{
		Id:       id,
		Text:     text,
		Selected: false,
	}

	l.Items = append(l.Items, Item)

	return Item

}

func (l *List) ToggleSelection() {
	Item := l.Items[l.CursorPos]
	Item.Selected = !Item.Selected
	l.RenderItems(true)

}

func (l *List) CursorDown() {
	if (l.CursorPos + 1) >= len(l.Items) {
		l.CursorPos = l.CursorPos
	} else {
		l.CursorPos = (l.CursorPos + 1)
	}
	l.RenderItems(true)

}

func (l *List) CursorUp() {

	if (l.CursorPos + len(l.Items) - 1) < len(l.Items) {
		l.CursorPos = l.CursorPos
	} else {
		l.CursorPos = (l.CursorPos + len(l.Items) - 1) % len(l.Items)
	}
	l.RenderItems(true)

}

func (l *List) SelectAll() {
	for _, i := range l.Items {
		if !i.Selected {
			i.Selected = true
		}
	}
	l.RenderItems(true)

}

func (l *List) DeselectAll() {
	for _, i := range l.Items {
		i.Selected = false
	}
	l.RenderItems(true)

}

func (l *List) FirstEntry() {
	l.CursorPos = 0
	l.RenderItems(true)

}

func (l *List) LastEntry() {
	l.CursorPos = len(l.Items) - 1
	l.RenderItems(true)

}

func (l *List) RenderItems(redraw bool) {
	if redraw {
		// Move cursor up by number of menu items.  Assumes each menu item is ONE line.
		// Use either native escape codes to move cursor or atomicgo cursor package.
		// https://github.com/atomicgo/cursor
		cursor.Up(len(m.MenuItems) - 1)
	}

	for index, Item := range l.Items {
		var newline = "\n"
		if index == len(l.Items)-1 {
			newline = ""
		}

		// prefix := "  "
		var prefix string
		if Item.Selected {
			prefix = "\u25c9"
			// prefix = "[X] "
		} else {
			prefix = "\u25ef"
			// prefix = "[ ] "
		}

		// CursorPos rendering.
		var cursor string
		if index == l.CursorPos {
			cursor = "> "
		} else {
			cursor = "  "
			fmt.Printf("\r%s%s %s%s", cursor, prefix, Item.Text, newline)
		}
	}

}

func (l *List) Render() {

	l.RenderItems(false)

	// Hide cursor.
	cursor.Hide()
}

func (l *List) Init() ([]string, bool) {
	selectedItems, escape := l.CallProcess()
	return selectedItems, escape
}

func (l *List) CallProcess() (results []string, escape bool) {
	defer func() {
		// Show cursor when CallProcess returns.
		cursor.Show()
	}()

	l.Render()

	for {
		// Use either getChar() function or atomicgo keyboard package.
		// https://github.com/atomicgo/keyboard
		ascii, keyCode, err := getChar()

		if (ascii == 3 || ascii == 27) || err != nil {
			log.Fatal(err)
			return []string{""}, true
		}

		// Handel keypress events.

	}
}

func _main() {
	list := InitializeList()
	files, err := os.ReadDir("/full/path/to/files")
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range files {
		list.AppendListItem(k, v.Name())
	}

	selectedItems, escaped := list.Init()
	if escaped {
		os.Exit(0)
	}

	fmt.Println(selectedItems)

}
