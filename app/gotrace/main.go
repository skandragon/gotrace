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
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sync"

	"github.com/pkg/profile"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

const (
	aspectRatio     = 16.0 / 9.0
	imageWidth      = 1200
	imageHeight     = int(float64(imageWidth) / aspectRatio)
	samplesPerPixel = 50
)

var (
	lookFrom = Vector3{13, 2, 3}
	lookAt   = Vector3{0, 0, 0}
	vup      = Vector3{0, 1, 0}

	r = math.Cos(math.Pi / 4)

	world = World{
		Camera:   NewCamera(lookFrom, lookAt, vup, 20, aspectRatio, 0.1, 10),
		TMin:     0.001,
		TMax:     math.MaxFloat64,
		MaxDepth: 500,
		Objects:  makeObjects(),
	}
)

func makeObjects() []Object {
	objects := []Object{}

	materialGround := NewLambertianMaterial(Vector3{0.5, 0.5, 0.5})
	objects = append(objects, Sphere{Vector3{0, -1000, 0}, 1000, materialGround})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := Vector3{
				float64(a) + 0.9*rand.Float64(),
				0.2,
				float64(b) + 0.9*rand.Float64(),
			}
			if center.Subtract(Vector3{4, 0.2, 0}).Length() > 0.9 {
				var sphereMaterial Material
				if chooseMat < 0.8 {
					albedo := RandomVector().Multiply(RandomVector())
					sphereMaterial = NewLambertianMaterial(albedo)
				} else if chooseMat < 0.95 {
					albedo := RandomVector().MultiplyScalar(0.5).AddScalar(0.5)
					fuzz := rand.Float64() * 0.5
					sphereMaterial = NewReflectiveMaterial(albedo, fuzz)
				} else {
					sphereMaterial = NewDielectricMaterial(1.5)
				}
				objects = append(objects, Sphere{center, 0.2, sphereMaterial})
			}
		}
	}

	material1 := NewDielectricMaterial(1.5)
	objects = append(objects, Sphere{Vector3{0, 1, 0}, 1.0, material1})

	material2 := NewLambertianMaterial(Vector3{0.4, 0.2, 0.1})
	objects = append(objects, Sphere{Vector3{-4, 1, 0}, 1.0, material2})

	material3 := NewReflectiveMaterial(Vector3{0.7, 0.6, 0.5}, 0.0)
	objects = append(objects, Sphere{Vector3{4, 1, 0}, 1.0, material3})

	return objects
}

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

func absorbLines(im *image.NRGBA, c chan processedLine) {
	for line := range c {
		log.Printf("Line %d of %d", line.y, imageHeight)
		y := im.Rect.Dy() - line.y - 1
		pixelOffset := (y-im.Rect.Min.Y)*im.Stride + (-im.Rect.Min.X)*4
		for _, color := range line.colors {
			im.Pix[pixelOffset] = uint8(color.X)
			pixelOffset++
			im.Pix[pixelOffset] = uint8(color.Y)
			pixelOffset++
			im.Pix[pixelOffset] = uint8(color.Z)
			pixelOffset++
			im.Pix[pixelOffset] = 0xff
			pixelOffset++
		}
	}
}

func worker(workerID int, world World, wg *sync.WaitGroup, w chan workItem, c chan processedLine) {
	defer wg.Done()
	log.Printf("Worker %d starting...", workerID)
	for work := range w {
		renderLine(world, work, c)
	}
	log.Printf("Worker %d ended.", workerID)
}

func renderLine(world World, work workItem, c chan processedLine) {
	colors := make([]Vector3, 0, work.imageWidth)
	for i := 0; i < work.imageWidth; i++ {
		rgb := Vector3{}
		for s := 0; s < work.samplesPerPixel; s++ {
			v := (float64(work.y) + rand.Float64()) / float64(work.imageHeight-1)
			u := (float64(i) + rand.Float64()) / float64(work.imageWidth-1)
			ray := world.Camera.GetRay(u, v)
			rgb = rgb.Add(world.Cast(ray, world.MaxDepth))
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

var (
	profileMemory = flag.Bool("profileMemory", false, "enable memory profiling")
	profileCPU    = flag.Bool("profileCPU", false, "enable CPU profiling")
	nCPU          = flag.Int("ncpu", runtime.NumCPU(), "Number of CPU cores to run on")
)

func main() {
	flag.Parse()

	if *profileCPU && *profileMemory {
		log.Fatal("Only one of -profileMemory or -profileCPU can be selected")
	}
	if *profileMemory {
		defer profile.Start(profile.ProfilePath("."), profile.MemProfile).Stop()
	}
	if *profileCPU {
		defer profile.Start(profile.ProfilePath(".")).Stop()
	}

	log.Println("NumCPU", *nCPU)

	im := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	resultChan := make(chan processedLine, imageHeight)
	workChan := make(chan workItem, imageHeight)
	wg := sync.WaitGroup{}
	for i := 0; i < *nCPU; i++ {
		wg.Add(1)
		go worker(i, world, &wg, workChan, resultChan)
	}

	go absorbLines(im, resultChan)
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
