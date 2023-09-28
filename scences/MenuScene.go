package scenes

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MenuScene struct {
	window         fyne.Window
	gameStarted    bool
	stopBackground chan struct{}
}

func NewMenuScene(fyneWindow fyne.Window) *MenuScene {
	scene := &MenuScene{window: fyneWindow}
	scene.Render()
	return scene
}

func (s *MenuScene) Render() {
	backgroundImage := canvas.NewImageFromFile("assets/0-0000.jpg")
	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 0))

	botonIniciar := widget.NewButton("Start Game", s.StartGame)
	botonIniciar.Resize(fyne.NewSize(150, 30))
	botonIniciar.Move(fyne.NewPos(325, 10))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, botonIniciar))
	s.stopBackground = make(chan struct{}) // Inicializa el canal

	// Goroutine para cambiar imágenes cada segundo
	go func() {
		imagePaths := []string{
			"assets/1-0000.jpg",
			"assets/2-0000.jpg",
			"assets/3-0000.jpg",
			"assets/4-0000.jpg",
			"assets/5-0000.jpg",
			"assets/6-0000.jpg",
			"assets/7-0000.jpg",
			"assets/8-0000.jpg",
			"assets/9-0000.jpg",
			"assets/10-0000.jpg",
			"assets/11-0000.jpg",
			"assets/12-0000.jpg",
			"assets/13-0000.jpg",
			"assets/14-0000.jpg",
			"assets/15-0000.jpg",
			"assets/16-0000.jpg",
			"assets/17-0000.jpg",
			"assets/18-0000.jpg",
			"assets/19-0000.jpg",
			"assets/20-0000.jpg",
			"assets/21-0000.jpg",
			"assets/22-0000.jpg",
			"assets/23-0000.jpg",
		}

		currentImageIndex := 0

		for {
			select {
			case <-s.stopBackground:
				// Se detiene la goroutine cuando se cierra el canal
				return
			default:
				time.Sleep(30 * time.Millisecond)

				// Cambia a la siguiente imagen en la lista
				currentImageIndex = (currentImageIndex + 1) % len(imagePaths)

				// Crea una nueva imagen de fondo con la nueva URI
				newBackgroundImage := canvas.NewImageFromFile(imagePaths[currentImageIndex])
				newBackgroundImage.Resize(fyne.NewSize(800, 600))
				newBackgroundImage.Move(fyne.NewPos(0, 0))

				// Reemplaza la imagen de fondo anterior
				s.window.SetContent(container.NewWithoutLayout(newBackgroundImage, botonIniciar))
			}
		}
	}()

}

func (s *MenuScene) StartGame() {
	close(s.stopBackground)
	NewGameScene(s.window)
}
