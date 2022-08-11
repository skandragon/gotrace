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
	"math"
	"math/rand"
)

// Camera defines how we see the world.
type Camera struct {
	ViewportHeight  float64
	ViewportWidth   float64
	FocalLength     float64
	LensRadius      float64
	Time0           float64
	Time1           float64
	origin          Vector3
	lowerLeftCorner Vector3
	horizontal      Vector3
	vertical        Vector3
	u, v, w         Vector3
}

// NewCamera returns a new Camera with the given aspect ratio.
// Other things are hard-coded currently.
func NewCamera(lookFrom Vector3, lookAt Vector3, vup Vector3, fieldOfView float64, aspectRatio float64, aperture float64, focalLength float64, time0 float64, time1 float64) Camera {
	ret := Camera{
		FocalLength: focalLength,
		LensRadius:  aperture / 2,
		Time0:       time0,
		Time1:       time1,
		origin:      lookFrom,
	}

	theta := fieldOfView * math.Pi / 180.0
	h := math.Tan(theta / 2)
	ret.ViewportHeight = 2.0 * h
	ret.ViewportWidth = ret.ViewportHeight * aspectRatio

	ret.w = lookFrom.Subtract(lookAt).Normalize()
	ret.u = vup.Cross(ret.w).Normalize()
	ret.v = ret.w.Cross(ret.u)

	ret.horizontal = ret.u.MultiplyScalar(ret.ViewportWidth * ret.FocalLength)
	ret.vertical = ret.v.MultiplyScalar(ret.ViewportHeight * ret.FocalLength)
	ret.lowerLeftCorner = ret.origin.
		Subtract(ret.horizontal.DivideScalar(2)).
		Subtract(ret.vertical.DivideScalar(2)).
		Subtract(ret.w.MultiplyScalar(ret.FocalLength))
	return ret
}

func randomBetween(a, b float64) float64 {
	r := b - a
	return a + rand.Float64()*r
}

// GetRay returns a ray from the camera's origin, pointing in the
// specified direction calculated by u, v.
func (c Camera) GetRay(s float64, t float64) Ray {
	rd := RandomUnitDisk().MultiplyScalar(c.LensRadius)
	offset := c.u.MultiplyScalar(rd.X).Add(c.v.MultiplyScalar(rd.Y))
	direction := c.lowerLeftCorner.
		Add(c.horizontal.MultiplyScalar(s)).
		Add(c.vertical.MultiplyScalar(t)).
		Subtract(c.origin).
		Subtract(offset)
	return Ray{c.origin.Add(offset), direction, randomBetween(c.Time0, c.Time1)}
}
