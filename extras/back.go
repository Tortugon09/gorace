package scenes

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type BackgroundAnimation struct {
	images          []fyne.Resource
	currentImage    int
	animationDelay  time.Duration
	backgroundImage *canvas.Image
}

func NewBackgroundAnimation(images []fyne.Resource, animationDelay time.Duration) *BackgroundAnimation {
	return &BackgroundAnimation{
		images:         images,
		currentImage:   0,
		animationDelay: animationDelay,
	}
}

func (ba *BackgroundAnimation) StartAnimation(window fyne.Window) {
	ba.backgroundImage = canvas.NewImageFromResource(ba.images[0])
	ba.backgroundImage.Resize(fyne.NewSize(800, 600))
	ba.backgroundImage.Move(fyne.NewPos(0, 0))

	go func() {
		for {
			time.Sleep(ba.animationDelay)
			ba.currentImage = (ba.currentImage + 1) % len(ba.images)
			window.Canvas().Refresh(ba.backgroundImage)
		}
	}()
}

func (ba *BackgroundAnimation) GetCurrentImage() *canvas.Image {
	return ba.backgroundImage
}

func (ba *BackgroundAnimation) StopAnimation() {
	// You can implement this method to stop the animation if needed.
}

func (ba *BackgroundAnimation) GetNextImage() fyne.Resource {
	return ba.images[ba.currentImage]
}
