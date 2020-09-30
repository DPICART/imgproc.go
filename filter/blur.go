package filter

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

type Blur struct {
}

func (f Blur) Name() string {
	return "blur"
}

func (f Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if nil != err {
		return err
	}
	img := imaging.Blur(src, 4)

	dstFile, err := os.Create(dstPath)
	if nil != err {
		return err
	}
	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}

	return jpeg.Encode(dstFile, img, &opts)
}
