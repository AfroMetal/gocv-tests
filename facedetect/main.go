package main

import (
	"fmt"
	"image"
	"image/color"
	"gocv.io/x/gocv"
	"strings"
	"path/filepath"
	"os"
)

var (
	PathToHaar = "/images/haarcascade_frontalface_default.xml"
	PathToImages = "/images/facedetect"
)

func detect(path string, info os.FileInfo) {
	img := gocv.IMRead(path, gocv.IMReadColor)
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
	
	gocv.IMWrite(strings.Replace(path, info.Name(), strings.Join([]string{"out", info.Name()}, "_"), 1), img)
}

func main() {

	allowedExt := []string{".jpg", ".png", ".JPG", ".PNG"}
	
    err := filepath.Walk(PathToImages, func(path string, info os.FileInfo, err error) error {
    	if info.IsDir() {
			return nil
		}
		for _, e := range allowedExt {
			if filepath.Ext(path) == e {
				detect(path, info)
        		return nil
			}
		}
        return nil
    })
    if err != nil {
        panic(err)
    }
}