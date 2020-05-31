package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)


type Flags struct {
	Depot string
}

var flags Flags

func init() {
	flag.StringVar(&flags.Depot, "depot", "depot", "Name of the depot sprite to use (default: depot)")

	flag.Parse()
}

func main() {
	for _, a := range flag.Args() {
		doImages(a)
		doIndexedImages(a)
	}
}

func doImages(name string) {
	bounds := image.Rect(0, 0, 600, 32)

	// Create the output image
	img := image.NewNRGBA(bounds)
	draw.Draw(img, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Over)

	// Draw 4x tracks
	tracks := loadImage("gui/" + name + "_32bpp.png").(*image.NRGBA)
	setTransparent(img, 0, 0, 20, 20)
	blit(img, tracks, 0, 3, 0, 6, 20, 15)

	setTransparent(img, 25, 0, 20, 20)
	blit(img, tracks, 25, 3, 34, 6, 20, 11)

	setTransparent(img, 50, 0, 20, 20)
	blit(img, tracks, 50, 8, 68, 8, 19, 8)

	setTransparent(img, 75, 0, 20, 20)
	blit(img, tracks, 75, 3, 101, 6, 20, 11)

	// Draw 1x crossing
	crossing := loadImage("gui/" + name + "_crossing_32bpp.png").(*image.NRGBA)
	setTransparent(img, 100, 0, 20, 20)
	blit(img, crossing, 100, 1, 34, 4, 20, 15)

	// Draw 1x depot (depot overlaid on track)
	depot := loadImage("gui/" + flags.Depot + "_32bpp.png").(*image.NRGBA)
	setTransparent(img, 125, 0, 20, 20)
	blit(img, tracks, 125, 1, 101, 0, 20, 19)
	blit(img, depot, 125, 0, 139, 0, 20, 20)

	// Draw 1x tunnel (tunnel overlaid on track)
	tunnel := loadImage("files/tunnel_gui_32bpp.png").(*image.NRGBA)
	setTransparent(img, 150, 0, 20, 20)
	blit(img, tracks, 150, 1, 101, 0, 20, 19)
	blit(img, tunnel, 150, 0, 0, 0, 20, 20)

	// Draw 1x convert (arrow overlaid on track)
	convertArrow := loadImage("files/convert_icon_32bpp.png").(*image.NRGBA)
	setTransparent(img, 175, 0, 20, 20)
	blit(img, tracks, 175, 9, 68, 8, 19, 8)
	blit(img, convertArrow, 175, 1, 0, 0, 20, 20)

	// Vertical (|) track
	genericButton := loadImage("files/icon_general_32bpp.png").(*image.NRGBA)
	setTransparent(img, 200, 0, 32, 32)
	blit(img, genericButton, 200, 0, 0, 0, 32, 32)
	blit(img, tracks, 210, 14, 0, 6, 20, 15)

	// Diagonal (/) track
	setTransparent(img, 250, 0, 32, 32)
	blit(img, genericButton, 250, 0, 0, 0, 32, 32)
	blit(img, tracks, 260, 15, 34, 6, 20, 11)

	// Horizontal (-) track
	setTransparent(img, 300, 0, 32, 32)
	blit(img, genericButton, 300, 0, 0, 0, 32, 32)
	blit(img, tracks, 310, 16, 68, 8, 19, 8)

	// Diagonal (\) track
	setTransparent(img, 350, 0, 32, 32)
	blit(img, genericButton, 350, 0, 0, 0, 32, 32)
	blit(img, tracks, 360, 15, 101, 6, 20, 11)

	// Crossing track
	setTransparent(img, 400, 0, 32, 32)
	blit(img, genericButton, 400, 0, 0, 0, 32, 32)
	blit(img, crossing, 410, 12, 34, 4, 20, 15)

	// Depot
	setTransparent(img, 450, 0, 32, 32)
	blit(img, genericButton, 450, 0, 0, 0, 32, 32)
	blit(img, tracks, 460, 12, 101, 0, 20, 19)
	blit(img, depot, 460, 11, 139, 0, 20, 20)

	// Tunnel
	setTransparent(img, 500, 0, 32, 32)
	blit(img, genericButton, 500, 0, 0, 0, 32, 32)
	blit(img, tracks, 510, 12, 101, 0, 20, 19)
	blit(img, tunnel, 510, 11, 0, 0, 20, 20)

	// Convert
	setTransparent(img, 550, 0, 32, 32)
	blit(img, genericButton, 550, 0, 0, 0, 32, 32)
	blit(img, tracks, 560, 20, 68, 8, 19, 8)
	blit(img, convertArrow, 560, 12, 0, 0, 20, 20)

	// Save the output image
	encode(img, "gui/gui_" + name + "_32bpp.png")

}

func setTransparent(dst *image.NRGBA, x, y, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			dst.Set(i+x, j+y, color.Transparent)
		}
	}
}

func setZero(dst *image.Paletted, x, y, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			dst.SetColorIndex(i+x, j+y, 0)
		}
	}
}

func blit8bpp(dst, src *image.Paletted, x1, y1, x2, y2, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if src.ColorIndexAt(i+x2, j+y2) != 0 {
				dst.SetColorIndex(i+x1, j+y1, src.ColorIndexAt(i+x2, j+y2))
			}
		}
	}
}

func blit(dst, src *image.NRGBA, x1, y1, x2, y2, w, h int) {
	bounds := image.Rect(x1, y1, x1 + w, y1 + h)
	loc := image.Point{x2, y2}

	draw.Draw(dst, bounds, src, loc, draw.Over)
}

func loadImage(filename string) image.Image {
	r, err := os.Open(filename)
	defer r.Close()
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	return img
}

func doIndexedImages(name string) {
	bounds := image.Rect(0, 0, 600, 32)

	var pal color.Palette
	tracks := loadImage("gui/" + name + "_8bpp.png").(*image.Paletted)
	if p, ok := tracks.ColorModel().(color.Palette); ok {
		pal = p
	}
	
	// Create the output image
	img := image.NewPaletted(bounds, pal)
	draw.Draw(img, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Over)

	// Draw 4x tracks
	setZero(img, 0, 0, 20, 20)
	blit8bpp(img, tracks, 0, 3, 0, 6, 20, 15)

	setZero(img, 25, 0, 20, 20)
	blit8bpp(img, tracks, 25, 3, 34, 6, 20, 11)

	setZero(img, 50, 0, 20, 20)
	blit8bpp(img, tracks, 50, 8, 68, 8, 19, 8)

	setZero(img, 75, 0, 20, 20)
	blit8bpp(img, tracks, 75, 3, 101, 6, 20, 11)

	// Draw 1x crossing
	crossing := loadImage("gui/" + name + "_crossing_8bpp.png").(*image.Paletted)
	setZero(img, 100, 0, 20, 20)
	blit8bpp(img, crossing, 100, 1, 34, 4, 20, 15)

	// Draw 1x depot (depot overlaid on track)
	depot := loadImage("gui/" + flags.Depot + "_8bpp.png").(*image.Paletted)
	setZero(img, 125, 0, 20, 20)
	blit8bpp(img, tracks, 125, 1, 101, 0, 20, 19)
	blit8bpp(img, depot, 125, 0, 139, 0, 20, 20)

	// Draw 1x tunnel (tunnel overlaid on track)
	tunnel := loadImage("files/tunnel_gui_8bpp.png").(*image.Paletted)
	setZero(img, 150, 0, 20, 20)
	blit8bpp(img, tracks, 150, 1, 101, 0, 20, 19)
	blit8bpp(img, tunnel, 150, 0, 0, 0, 20, 20)

	// Draw 1x convert (arrow overlaid on track)
	convertArrow := loadImage("files/convert_icon_8bpp.png").(*image.Paletted)
	setZero(img, 175, 0, 20, 20)
	blit8bpp(img, tracks, 175, 9, 68, 8, 19, 8)
	blit8bpp(img, convertArrow, 175, 1, 0, 0, 20, 20)

	// Vertical (|) track
	genericButton := loadImage("files/icon_general_8bpp.png").(*image.Paletted)
	setZero(img, 200, 0, 32, 32)
	blit8bpp(img, genericButton, 200, 0, 0, 0, 32, 32)
	blit8bpp(img, tracks, 210, 14, 0, 6, 20, 15)

	// Diagonal (/) track
	setZero(img, 250, 0, 32, 32)
	blit8bpp(img, genericButton, 250, 0, 0, 0, 32, 32)
	blit8bpp(img, tracks, 260, 15, 34, 6, 20, 11)

	// Horizontal (-) track
	setZero(img, 300, 0, 32, 32)
	blit8bpp(img, genericButton, 300, 0, 0, 0, 32, 32)
	blit8bpp(img, tracks, 310, 16, 68, 8, 19, 8)

	// Diagonal (\) track
	setZero(img, 350, 0, 32, 32)
	blit8bpp(img, genericButton, 350, 0, 0, 0, 32, 32)
	blit8bpp(img, tracks, 360, 15, 101, 6, 20, 11)

	// Crossing track
	setZero(img, 400, 0, 32, 32)
	blit8bpp(img, genericButton, 400, 0, 0, 0, 32, 32)
	blit8bpp(img, crossing, 410, 12, 34, 4, 20, 15)

	// Depot
	setZero(img, 450, 0, 32, 32)
	blit8bpp(img, genericButton, 450, 0, 0, 0, 32, 32)
	blit8bpp(img, tracks, 460, 12, 101, 0, 20, 19)
	blit8bpp(img, depot, 460, 11, 139, 0, 20, 20)

	// Tunnel
	setZero(img, 500, 0, 32, 32)
	blit8bpp(img, genericButton, 500, 0, 0, 0, 32, 32)
	blit8bpp(img, tracks, 510, 12, 101, 0, 20, 19)
	blit8bpp(img, tunnel, 510, 11, 0, 0, 20, 20)

	// Convert
	setZero(img, 550, 0, 32, 32)
	blit8bpp(img, genericButton, 550, 0, 0, 0, 32, 32)
	blit8bpp(img, tracks, 560, 20, 68, 8, 19, 8)
	blit8bpp(img, convertArrow, 560, 12, 0, 0, 20, 20)

	// Save the output image
	encode(img, "gui/gui_" + name + "_8bpp.png")
}

func encode(img image.Image, filename string) {
	of, err := os.Create(filename)
	defer of.Close()
	if err != nil {
		panic(err)
	}

	if err := png.Encode(of, img); err != nil {
		panic(err)
	}
}