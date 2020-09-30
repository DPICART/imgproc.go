package filter

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

type Grayscale struct {
}

func (f Grayscale) Name() string {
	return "grayscale"
}

func (f Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if nil != err {
		return err
	}
	img := imaging.Grayscale(src)

	dstFile, err := os.Create(dstPath)
	if nil != err {
		return err
	}
	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}

	return jpeg.Encode(dstFile, img, &opts)
}
