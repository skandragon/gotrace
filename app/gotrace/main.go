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
	"math"
	"math/rand"
	"os"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

func setPixel(im *image.RGBA, samples int, x int, y int, color Vector3) {
	pixelColor := color.
		DivideScalar(float64(samples)).
		Gamma2().
		Clamp(0, 0.999).
		MultiplyScalar(256)

	pixelOffset := (im.Rect.Dy()-y-1)*im.Stride + x*4
	im.Pix[pixelOffset] = uint8(pixelColor.X)
	im.Pix[pixelOffset+1] = uint8(pixelColor.Y)
	im.Pix[pixelOffset+2] = uint8(pixelColor.Z)
	im.Pix[pixelOffset+3] = 0xff
}

var world = &World{
	TMin:     0.001,
	TMax:     math.MaxFloat64,
	MaxDepth: 50,
	Lights:   []Light{},
	Objects: []Object{
		Sphere{Center: Vector3{0, 0, -1}, Radius: 0.5, Color: Vector3{1, 1, 1}},
		Sphere{Center: Vector3{0, -100.5, -1}, Radius: 100, Color: Vector3{1, 1, 1}},
	},
}

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samplesPerPixel := 100

	camera := NewCamera(aspectRatio)

	im := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			rgb := &Vector3{}
			for s := 0; s < samplesPerPixel; s++ {
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)
				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				ray := camera.GetRay(u, v)
				rgb.AddAccum(world.Cast(ray, world.MaxDepth))
			}
			setPixel(im, samplesPerPixel, i, j, *rgb)
		}
	}

	out, err := os.Create("out.png")
	check(err, "Error writing to file: %v\n")
	defer out.Close()
	png.Encode(out, im)
}
