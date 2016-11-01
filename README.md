# resizegif.go

Resize animated GIF file without destroying animation Inspired by [go-scratch](https://github.com/dpup/go-scratch/blob/master/gif-resize)

# Methods

    // Resize the gif to another thumbnail gif
	func Resize(srcFile string, width int, height int) (*gif.GIF, error)
	
	// Save gif to file
	func Save(gifImg *gif.GIF, desFile string) error

# Usage

	// To generate a width 100px gif from the given demo.gif
	img, err := resizegif.Resize("demo.gif", 100, 0)
	resizegif.Save(img, "demo.w100.gif")
	
