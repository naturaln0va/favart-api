package utility

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"

	// used to read png files
	_ "image/png"

	"github.com/nfnt/resize"
)

// CreateThumbnail returns a thumbnail for the provided image file
func CreateThumbnail(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	imgSize := img.Bounds().Size()

	cr := cropRect(imgSize)
	cropped, err := cropAndFillWithCopy(img, cr)
	if err != nil {
		return err
	}

	thumb := resize.Resize(120, 0, cropped, resize.Lanczos3)

	options := jpeg.Options{Quality: 75}
	return jpeg.Encode(w, thumb, &options)
}

func cropRect(size image.Point) image.Rectangle {
	w := size.X
	h := size.Y

	if w > h {
		s := (w - h) / 2
		return image.Rect(s, 0, s+h, h)
	}

	s := (h - w) / 2
	return image.Rect(0, s, w, s+w)
}

func cropAndFillWithCopy(img image.Image, cr image.Rectangle) (image.Image, error) {
	result := image.NewRGBA(cr)
	draw.Draw(result, cr, &image.Uniform{color.White}, cr.Min, draw.Src) // fill with white
	draw.Draw(result, cr, img, cr.Min, draw.Over)
	return result, nil
}
