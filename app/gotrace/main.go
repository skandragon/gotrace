/*
 * Copyright 2022 Michael Graff.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func setPixel(im *image.RGBA, samples int, x int, y int, color Vector3) {
	pixelOffset := (im.Rect.Dy()-y-1)*im.Stride + x*4
	im.Pix[pixelOffset] = uint8(color.X)
	im.Pix[pixelOffset+1] = uint8(color.Y)
	im.Pix[pixelOffset+2] = uint8(color.Z)
	im.Pix[pixelOffset+3] = 0xff
}

var (
	materialGround = NewLamtertianMaterial(Vector3{0.8, 0.8, 0})
	materialCenter = NewLamtertianMaterial(Vector3{0.7, 0.3, 0.3})
	materialLeft   = NewReflectiveMaterial(Vector3{0.8, 0.8, 0.8}, 0.3)
	materialRight  = NewReflectiveMaterial(Vector3{0.8, 0.6, 0.2}, 1.0)
	materialGlass  = NewDielectricMaterial(1.5)

	r = math.Cos(math.Pi / 4)

	world = &World{
		TMin:     0.001,
		TMax:     math.MaxFloat64,
		MaxDepth: 50,
		Lights:   []Light{},
		Objects: []Object{
			//Sphere{Vector3{0, 0, -1}, 0.5, materialCenter},
			Sphere{Vector3{0, -100.5, -1}, 100, materialGround},
			Sphere{Vector3{-r, 0, -1}, r, materialLeft},
			//Sphere{Vector3{1, 0, -1}, 0.25, materialRight},
			Sphere{Vector3{r, 0, -1}, r, materialGlass},
		},
	}
)

type processedLine struct {
	y      int
	colors []Vector3
}

type workItem struct {
	y               int
	imageHeight     int
	imageWidth      int
	samplesPerPixel int
}

func absorbLines(im *image.RGBA, samples int, c chan processedLine) {
	for line := range c {
		for x, color := range line.colors {
			setPixel(im, samples, x, line.y, color)
		}
	}
}

func worker(workerID int, camera Camera, wg *sync.WaitGroup, w chan workItem, c chan processedLine) {
	defer wg.Done()
	log.Printf("Worker %d starting...", workerID)

	for work := range w {
		renderLine(camera, work, c)
	}
	log.Printf("Worker %d ended.", workerID)
}

func renderLine(camera Camera, work workItem, c chan processedLine) {
	colors := make([]Vector3, 0, work.imageWidth)
	for i := 0; i < work.imageWidth; i++ {
		rgb := &Vector3{}
		for s := 0; s < work.samplesPerPixel; s++ {
			v := (float64(work.y) + rand.Float64()) / float64(work.imageHeight-1)
			u := (float64(i) + rand.Float64()) / float64(work.imageWidth-1)
			ray := camera.GetRay(u, v)
			rgb.AddAccum(world.Cast(ray, world.MaxDepth))
		}
		pixelColor := rgb.
			DivideScalar(float64(work.samplesPerPixel)).
			Gamma2().
			Clamp(0, 0.999).
			MultiplyScalar(256)

		colors = append(colors, pixelColor)
	}
	c <- processedLine{work.y, colors}
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samplesPerPixel := 50

	camera := NewCamera(90, aspectRatio)

	im := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	resultChan := make(chan processedLine)
	workChan := make(chan workItem, imageHeight)
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go worker(i, camera, &wg, workChan, resultChan)
	}

	go absorbLines(im, samplesPerPixel, resultChan)
	for j := 0; j < imageHeight; j++ {
		workChan <- workItem{j, imageHeight, imageWidth, samplesPerPixel}
	}
	close(workChan)
	log.Printf("Waiting for workers to complete...")
	wg.Wait()
	close(resultChan)

	out, err := os.Create("out.png")
	check(err, "Error writing to file: %v\n")
	defer out.Close()
	png.Encode(out, im)
}
