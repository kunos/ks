package ks

type IGameRunner interface {
	Step(float32)
	Update(float32)
	Render(float32)
	IsActive() bool
}

var gameRunners []IGameRunner

func init() {
}

func AddGameRunner(object IGameRunner) {
	gameRunners = append(gameRunners, object)
}

func RunGame() {

	for {
		dt := float32(0.003)
		Update(dt)
		Render(dt)
	}

}

func Render(dt float32) {
	for _, gr := range gameRunners {
		if gr.IsActive() {
			gr.Render(dt)
		}
	}
}

func Update(dt float32) {
	for _, gr := range gameRunners {
		if gr.IsActive() {
			gr.Update(dt)
		}

	}

}

func Step(dt float32) {
	for _, gr := range gameRunners {
		if gr.IsActive() {
			gr.Step(dt)
		}
	}
}
