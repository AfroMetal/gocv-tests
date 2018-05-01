package main

import (
	"fmt"
	"image"
	"image/color"
	"gocv.io/x/gocv"
	"strings"
	"io/ioutil"
	"path/filepath"
)

var (
	PathToHaar = "/images/haarcascade_frontalface_default.xml"
	PathToImages = "/images/facedetect"
)

func main() {

	// open webcam
	files, err := ioutil.ReadDir(PathToImages)
    if err != nil {
        panic(err)
    }
	
	for _, fileInfo := range files {
		file := filepath.Join(PathToImages, fileInfo.Name())
		img := gocv.IMRead(file, gocv.IMReadColor)
		defer img.Close()
		
		// color for the rect when faces detected
		blue := color.RGBA{0, 0, 255, 0}
		
		// load classifier to recognize faces
		classifier := gocv.NewCascadeClassifier()
		defer classifier.Close()
		
		if !classifier.Load(PathToHaar) {
			fmt.Printf("Error reading cascade file: %v\n", PathToHaar)
			return
		}
		
		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))
		
		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
			
			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
            pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
            gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}
		gocv.IMWrite(strings.Replace(file, ".jpg", "_result.jpg", 1), img)
	}
}