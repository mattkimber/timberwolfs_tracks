package main

import (
	"flag"
	"fmt"
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
		doImages(1, a)
		doIndexedImages(1, a)
		doImages(2, a)
		doIndexedImages(2, a)
	}
}

func doImages(scale int, name string) {
	bounds := image.Rect(0, 0, 600*scale, 32*scale)

	// Create the output image
	img := image.NewNRGBA(bounds)
	draw.Draw(img, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Over)

	// Draw 4x tracks
	tracks := loadImage(fmt.Sprintf("gui/%dx/%s_32bpp.png", scale, name)).(*image.NRGBA)
	setTransparent(scale, img, 0, 0, 20, 20)
	blit(scale, img, tracks, 0, 3, 0, 6, 20, 15)

	setTransparent(scale, img, 25, 0, 20, 20)
	blit(scale, img, tracks, 25, 3, 34, 6, 20, 11)

	setTransparent(scale, img, 50, 0, 20, 20)
	blit(scale, img, tracks, 50, 8, 68, 8, 19, 8)

	setTransparent(scale, img, 75, 0, 20, 20)
	blit(scale, img, tracks, 75, 3, 101, 6, 20, 11)

	// Draw 1x crossing
	crossing := loadImage(fmt.Sprintf("gui/%dx/%s_crossing_32bpp.png", scale, name)).(*image.NRGBA)
	setTransparent(scale, img, 100, 0, 20, 20)
	blit(scale, img, crossing, 100, 1, 34, 4, 20, 15)

	// Draw 1x depot (depot overlaid on track)
	depot := loadImage(fmt.Sprintf("gui/%dx/%s_32bpp.png", scale, flags.Depot)).(*image.NRGBA)
	setTransparent(scale, img, 125, 0, 20, 20)
	blit(scale, img, tracks, 125, 1, 101, 0, 20, 19)
	blit(scale, img, depot, 125, 0, 139, 0, 20, 20)

	// Draw 1x tunnel (tunnel overlaid on track)
	tunnelFG := loadImage(fmt.Sprintf("files/%dx/tunnel_gui_fg_32bpp.png", scale)).(*image.NRGBA)
	tunnelBG := loadImage(fmt.Sprintf("files/%dx/tunnel_gui_bg_32bpp.png", scale)).(*image.NRGBA)
	setTransparent(scale, img, 150, 0, 20, 20)
	blit(scale, img, tunnelBG, 150, 0, 0, 0, 20, 20)
	blit(scale, img, tracks, 150, 10, 101, 10, 20, 9)
	blit(scale, img, tunnelFG, 150, 0, 0, 0, 20, 20)

	// Draw 1x convert (arrow overlaid on track)
	convertArrow := loadImage(fmt.Sprintf("files/%dx/convert_icon_32bpp.png", scale)).(*image.NRGBA)
	setTransparent(scale, img, 175, 0, 20, 20)
	blit(scale, img, tracks, 175, 9, 68, 8, 19, 8)
	blit(scale, img, convertArrow, 175, 1, 0, 0, 20, 20)

	// Vertical (|) track
	genericButton := loadImage(fmt.Sprintf("files/%dx/icon_general_32bpp.png", scale)).(*image.NRGBA)
	setTransparent(scale, img, 200, 0, 32, 32)
	blit(scale, img, genericButton, 200, 0, 0, 0, 32, 32)
	blit(scale, img, tracks, 210, 14, 0, 6, 20, 15)

	// Diagonal (/) track
	setTransparent(scale, img, 250, 0, 32, 32)
	blit(scale, img, genericButton, 250, 0, 0, 0, 32, 32)
	blit(scale, img, tracks, 260, 15, 34, 6, 20, 11)

	// Horizontal (-) track
	setTransparent(scale, img, 300, 0, 32, 32)
	blit(scale, img, genericButton, 300, 0, 0, 0, 32, 32)
	blit(scale, img, tracks, 310, 16, 68, 8, 19, 8)

	// Diagonal (\) track
	setTransparent(scale, img, 350, 0, 32, 32)
	blit(scale, img, genericButton, 350, 0, 0, 0, 32, 32)
	blit(scale, img, tracks, 360, 15, 101, 6, 20, 11)

	// Crossing track
	setTransparent(scale, img, 400, 0, 32, 32)
	blit(scale, img, genericButton, 400, 0, 0, 0, 32, 32)
	blit(scale, img, crossing, 410, 12, 34, 4, 20, 15)

	// Depot
	setTransparent(scale, img, 450, 0, 32, 32)
	blit(scale, img, genericButton, 450, 0, 0, 0, 32, 32)
	blit(scale, img, tracks, 460, 12, 101, 0, 20, 19)
	blit(scale, img, depot, 460, 11, 139, 0, 20, 20)

	// Tunnel
	setTransparent(scale, img, 500, 0, 32, 32)
	blit(scale, img, genericButton, 500, 0, 0, 0, 32, 32)
	blit(scale, img, tunnelBG, 510, 11, 0, 0, 20, 20)
	blit(scale, img, tracks, 510, 22, 101, 10, 20, 9)
	blit(scale, img, tunnelFG, 510, 11, 0, 0, 20, 20)

	// Convert
	setTransparent(scale, img, 550, 0, 32, 32)
	blit(scale, img, genericButton, 550, 0, 0, 0, 32, 32)
	blit(scale, img, tracks, 560, 20, 68, 8, 19, 8)
	blit(scale, img, convertArrow, 560, 12, 0, 0, 20, 20)

	// Save the output image
	encode(img, fmt.Sprintf("gui/%dx/gui_%s_32bpp.png", scale, name))

}

func setTransparent(scale int, dst *image.NRGBA, x, y, w, h int) {
	for i := 0; i < w*scale; i++ {
		for j := 0; j < h*scale; j++ {
			dst.Set(i+(x*scale), j+(y*scale), color.Transparent)
		}
	}
}

func setZero(scale int, dst *image.Paletted, x, y, w, h int) {
	for i := 0; i < w*scale; i++ {
		for j := 0; j < h*scale; j++ {
			dst.SetColorIndex(i+(x*scale), j+(y*scale), 0)
		}
	}
}

func blit8bpp(scale int, dst, src *image.Paletted, x1, y1, x2, y2, w, h int) {
	for i := 0; i < w*scale; i++ {
		for j := 0; j < h*scale; j++ {
			if src.ColorIndexAt(i+(x2*scale), j+(y2*scale)) != 0 {
				dst.SetColorIndex(i+(x1*scale), j+(y1*scale), src.ColorIndexAt(i+(x2*scale), j+(y2*scale)))
			}
		}
	}
}

func blit(scale int, dst, src *image.NRGBA, x1, y1, x2, y2, w, h int) {
	bounds := image.Rect(x1*scale, y1*scale, (x1+w)*scale, (y1+h)*scale)
	loc := image.Point{X: x2 * scale, Y: y2 * scale}

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

func doIndexedImages(scale int, name string) {
	bounds := image.Rect(0, 0, 600*scale, 32*scale)

	var pal color.Palette
	tracks := loadImage(fmt.Sprintf("gui/%dx/%s_8bpp.png", scale, name)).(*image.Paletted)
	if p, ok := tracks.ColorModel().(color.Palette); ok {
		pal = p
	}

	// Create the output image
	img := image.NewPaletted(bounds, pal)
	draw.Draw(img, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Over)

	// Draw 4x tracks
	setZero(scale, img, 0, 0, 20, 20)
	blit8bpp(scale, img, tracks, 0, 3, 0, 6, 20, 15)

	setZero(scale, img, 25, 0, 20, 20)
	blit8bpp(scale, img, tracks, 25, 3, 34, 6, 20, 11)

	setZero(scale, img, 50, 0, 20, 20)
	blit8bpp(scale, img, tracks, 50, 8, 68, 8, 19, 8)

	setZero(scale, img, 75, 0, 20, 20)
	blit8bpp(scale, img, tracks, 75, 3, 101, 6, 20, 11)

	// Draw 1x crossing
	crossing := loadImage(fmt.Sprintf("gui/%dx/%s_crossing_8bpp.png", scale, name)).(*image.Paletted)
	setZero(scale, img, 100, 0, 20, 20)
	blit8bpp(scale, img, crossing, 100, 1, 34, 4, 20, 15)

	// Draw 1x depot (depot overlaid on track)
	depot := loadImage(fmt.Sprintf("gui/%dx/%s_8bpp.png", scale, flags.Depot)).(*image.Paletted)
	setZero(scale, img, 125, 0, 20, 20)
	blit8bpp(scale, img, tracks, 125, 1, 101, 0, 20, 19)
	blit8bpp(scale, img, depot, 125, 0, 139, 0, 20, 20)

	// Draw 1x tunnel (tunnel overlaid on track)
	tunnelFG := loadImage(fmt.Sprintf("files/%dx/tunnel_gui_fg_8bpp.png", scale)).(*image.Paletted)
	tunnelBG := loadImage(fmt.Sprintf("files/%dx/tunnel_gui_bg_8bpp.png", scale)).(*image.Paletted)
	setZero(scale, img, 150, 0, 20, 20)
	blit8bpp(scale, img, tunnelBG, 150, 0, 0, 0, 20, 20)
	blit8bpp(scale, img, tracks, 150, 10, 101, 10, 20, 9)
	blit8bpp(scale, img, tunnelFG, 150, 0, 0, 0, 20, 20)

	// Draw 1x convert (arrow overlaid on track)
	convertArrow := loadImage(fmt.Sprintf("files/%dx/convert_icon_8bpp.png", scale)).(*image.Paletted)
	setZero(scale, img, 175, 0, 20, 20)
	blit8bpp(scale, img, tracks, 175, 9, 68, 8, 19, 8)
	blit8bpp(scale, img, convertArrow, 175, 1, 0, 0, 20, 20)

	// Vertical (|) track
	genericButton := loadImage(fmt.Sprintf("files/%dx/icon_general_8bpp.png", scale)).(*image.Paletted)
	setZero(scale, img, 200, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 200, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, tracks, 210, 14, 0, 6, 20, 15)

	// Diagonal (/) track
	setZero(scale, img, 250, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 250, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, tracks, 260, 15, 34, 6, 20, 11)

	// Horizontal (-) track
	setZero(scale, img, 300, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 300, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, tracks, 310, 16, 68, 8, 19, 8)

	// Diagonal (\) track
	setZero(scale, img, 350, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 350, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, tracks, 360, 15, 101, 6, 20, 11)

	// Crossing track
	setZero(scale, img, 400, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 400, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, crossing, 410, 12, 34, 4, 20, 15)

	// Depot
	setZero(scale, img, 450, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 450, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, tracks, 460, 12, 101, 0, 20, 19)
	blit8bpp(scale, img, depot, 460, 11, 139, 0, 20, 20)

	// Tunnel
	setZero(scale, img, 500, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 500, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, tunnelBG, 510, 11, 0, 0, 20, 20)
	blit8bpp(scale, img, tracks, 510, 22, 101, 10, 20, 9)
	blit8bpp(scale, img, tunnelFG, 510, 11, 0, 0, 20, 20)

	// Convert
	setZero(scale, img, 550, 0, 32, 32)
	blit8bpp(scale, img, genericButton, 550, 0, 0, 0, 32, 32)
	blit8bpp(scale, img, tracks, 560, 20, 68, 8, 19, 8)
	blit8bpp(scale, img, convertArrow, 560, 12, 0, 0, 20, 20)

	// Save the output image
	encode(img, fmt.Sprintf("gui/%dx/gui_%s_8bpp.png", scale, name))
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
