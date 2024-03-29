// Package avatar contains functions for creating user avatars.
package avatar // import "code.soquee.net/avatar"

import (
	"image"
	"image/color"
	"image/draw"

	colorHash "mellium.im/xmpp/color"
)

const (
	edgeLen = 512
)

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// New returns a new avatar that is a hash of name.
func New(name string) image.Image {
	c := colorHash.String(name, 187, colorHash.None)

	img := image.NewPaletted(
		image.Rect(0, 0, edgeLen, edgeLen),
		color.Palette{c, color.White},
	)
	src := image.NewPaletted(
		image.Rect(0, 0, edgeLen, edgeLen),
		color.Palette{color.White},
	)

	draw.DrawMask(img, img.Bounds(), src, image.ZP, &circle{image.Point{X: edgeLen / 2, Y: edgeLen / 2}, edgeLen * 0.25}, image.ZP, draw.Over)

	return img
}
