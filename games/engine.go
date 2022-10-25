package games

import (
	"fmt"
	"github.com/devdynam0507/dyworld-go-graphics"
	"math"
	"runtime"
	"time"
	"sync"
	"math/rand"
)

var player Player;
var obstacles Obstacles;
var env GameEnvironment;
var wait sync.WaitGroup
var score float64 = 0.000

const (
	NOT_EXISTS = -1
	GROUND_HEIGHT = 10
	JUMP_SPEED = 3
)

func draw() {
	graphics.DrawCall(func () {
		graphics.DrawUI(int(math.Round(score)), env.FrameRate)
		graphics.DrawCell(player.x, int(player.y), graphics.Yellow, graphics.Default, player.symbol)
		if obstacles.x > 0 {
			graphics.DrawCell(obstacles.x, 10, graphics.Red, graphics.Default, obstacles.symbol)
		}
	})
}

func determineCreateObstacles() bool {
	pick := rand.Int31n(100)
	return pick <= env.Difficulty
}

func updateCharacter() {
	if player.isJump {
		h := (player.jumpTime * player.jumpTime * (-0.7) / 2.5) + (player.jumpTime + 0.1)
		player.y += -h
		player.jumpTime += 0.5
		if int(player.y) >= GROUND_HEIGHT {
			player.isJump = false
			player.jumpTime = 0.0
			player.y = GROUND_HEIGHT
		}
	}
}

func done() {
	graphics.Close()
	wait.Done()
}

func update() {
	// if not exists obstacles in field
	if obstacles.x <= 0 && determineCreateObstacles() {
		obstacles.x = 30
	} else if obstacles.x > 0 {
		obstacles.x -= 1
	}
	updateCharacter()

	score += 0.1
	env.FrameRate += 0.001
}

func checkCollision() bool {
	return int(player.x) == obstacles.x && player.y == GROUND_HEIGHT
}

func gameloop() {
	rate := time.Duration(1000 / int64(env.FrameRate))
	fmt.Println(rate)
	for range time.Tick(time.Millisecond * rate) {
		update()
		draw()
		// if collision, game over
		if checkCollision() {
			done()
		}
	}
}

func initializeKeyInput() {
	RegisterKeyboardHandler(func (key uint16) {
		switch key {
		case ESC:
			done()
		case SPACE:
			if !player.isJump {
				player.isJump = true				
			}
		}
	})
	go StartCaptureKeyboard()
}

func Initialize(gameEnv GameEnvironment) {
	env = gameEnv
	player.symbol = "ጿ"
	player.isJump = false
	player.x = 3
	player.y = GROUND_HEIGHT
	player.jumpTime = 0.0
	obstacles.x = NOT_EXISTS // -1 is not exists on field
	obstacles.symbol = "ψ"
	
	// Use 2 cpu cores (Game Loop, Keyboard I/O)
	runtime.GOMAXPROCS(env.Cpu)
	graphics.Initialize()
}

func Run() {
	// Wating 1 go routine (Waiting for gameloop...)
	wait.Add(1)
	initializeKeyInput()
	go gameloop()

	wait.Wait()
}