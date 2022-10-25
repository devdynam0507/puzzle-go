package main

import (
	"github.com/devdynam0507/dyword-go-games"
)

func main() {
	env := games.GameEnvironment { 
		Cpu: 2,
		FrameRate: 25.0, // 25 FPS
		Difficulty: 90,
	}
	games.Initialize(env)
	games.Run()
}