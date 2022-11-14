package main

import (
	"fmt"
	"image"
	"math/rand"
	"time"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
)

func main() {
	originalImage, _ := mergi.Import(impexp.NewFileImporter("example.png"))
	watermarkImage, _ := mergi.Import(impexp.NewFileImporter("mask.png"))
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(originalImage.Bounds().Max.X - 100)
	y := rand.Intn(originalImage.Bounds().Max.Y - 100)
	fmt.Println(x)
	fmt.Println(y)
	res, _ := mergi.Watermark(watermarkImage, originalImage, image.Pt(x, y))
	mergi.Export(impexp.NewFileExporter(res, "watermark.png"))
	resCrop, _ := mergi.Crop(originalImage, image.Pt(x, y), image.Pt(100, 100))
	mergi.Export(impexp.NewFileExporter(resCrop, "crop.png"))
}
