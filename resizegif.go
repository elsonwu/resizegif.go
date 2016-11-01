package resizegif

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"

	"github.com/nfnt/resize"
)

// Resize the gif to another thumbnail gif
func Resize(srcFile string, width int, height int) (*gif.GIF, error) {

	f, err := os.Open(srcFile)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	im, err := gif.DecodeAll(f)

	if err != nil {
		return nil, err
	}

	if width == 0 {
		width = int(im.Config.Width * height / im.Config.Width)
	} else if height == 0 {
		height = int(width * im.Config.Height / im.Config.Width)
	}

	// reset the gif width and height
	im.Config.Width = width
	im.Config.Height = height

	firstFrame := im.Image[0].Bounds()
	img := image.NewRGBA(image.Rect(0, 0, firstFrame.Dx(), firstFrame.Dy()))

	// resize frame by frame
	for index, frame := range im.Image {
		b := frame.Bounds()
		draw.Draw(img, b, frame, b.Min, draw.Over)
		im.Image[index] = ImageToPaletted(resize.Resize(uint(width), uint(height), img, resize.NearestNeighbor))
	}

	return im, nil
}

// Save gif file
func Save(gifImg *gif.GIF, desFile string) error {
	f, err := os.Create(desFile)
	defer f.Close()
	if err != nil {
		return err
	}

	return gif.EncodeAll(f, gifImg)
}

func ImageToPaletted(img image.Image) *image.Paletted {
	b := img.Bounds()
	pm := image.NewPaletted(b, palette.Plan9)
	draw.FloydSteinberg.Draw(pm, b, img, image.ZP)
	return pm
}
