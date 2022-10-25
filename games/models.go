package games

type GameEnvironment struct {
	Cpu int
	FrameRate float64
	Difficulty int32
}

type Player struct {
	symbol string
	isJump bool
	x int
	jumpTime, y float32
}

type Obstacles struct {
	symbol string
	x int
}