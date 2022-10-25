module github.com/devdynam0507/dyworld-go-app

go 1.19

replace github.com/devdynam0507/dyworld-go-graphics => ../graphics

replace github.com/devdynam0507/dyword-go-games => ../games

require github.com/devdynam0507/dyword-go-games v0.0.0-00010101000000-000000000000

require (
	github.com/devdynam0507/dyworld-go-graphics v0.0.0-00010101000000-000000000000 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/nsf/termbox-go v1.1.1 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
)
