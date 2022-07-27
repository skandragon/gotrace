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
	"os"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

const colorMax = 255.99

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	viewportHeight := 2.0
	viewportWidth := viewportHeight * aspectRatio
	focalLength := 1.0

	origin := Vector3{0.0, 0.0, 0.0}
	horizontal := Vector3{viewportWidth, 0.0, 0.0}
	vertical := Vector3{0.0, viewportHeight, 0.0}
	lowerLeft := origin.
		Subtract(horizontal.DivideScalar(2)).
		Subtract(vertical.DivideScalar(2)).
		Subtract(Vector3{0, 0, focalLength})

	im := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		v := float64(j) / float64(imageHeight-1)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)

			position := horizontal.MultiplyScalar(u).Add(vertical.MultiplyScalar(v))

			// direction = lowerLeft + (u * horizontal) + (v * vertical)
			direction := lowerLeft.Add(position)

			rgb := Ray{origin, direction}.Cast()

			color := rgb.MultiplyScalar(colorMax)

			pixelOffset := (imageHeight-j-1)*im.Stride + i*4
			im.Pix[pixelOffset] = uint8(color.X)
			im.Pix[pixelOffset+1] = uint8(color.Y)
			im.Pix[pixelOffset+2] = uint8(color.Z)
			im.Pix[pixelOffset+3] = 0xff
		}
	}

	out, err := os.Create("out.png")
	check(err, "Error writing to file: %v\n")
	defer out.Close()
	png.Encode(out, im)
}
