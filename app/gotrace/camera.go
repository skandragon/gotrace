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

// Camera defines how we see the world.
type Camera struct {
	ViewportHeight  float64
	ViewportWidth   float64
	FocalLength     float64
	origin          Vector3
	lowerLeftCorner Vector3
	horizontal      Vector3
	vertical        Vector3
}

// NewCamera returns a new Camera with the given aspect ratio.
// Other things are hard-coded currently.
func NewCamera(aspectRatio float64) Camera {
	ret := Camera{
		ViewportHeight: 2.0,
		FocalLength:    1.0,
		origin:         Vector3{0, 0, 0},
	}
	ret.ViewportWidth = ret.ViewportHeight * aspectRatio
	ret.horizontal = Vector3{ret.ViewportWidth, 0, 0}
	ret.vertical = Vector3{0, ret.ViewportHeight, 0}
	ret.lowerLeftCorner = ret.origin.
		Subtract(ret.horizontal.DivideScalar(2)).
		Subtract(ret.vertical.DivideScalar(2)).
		Subtract(Vector3{0, 0, ret.FocalLength})
	return ret
}

// GetRay returns a ray from the camera's origin, pointing in the
// specified direction calculated by u, v.
func (c Camera) GetRay(u, v float64) Ray {
	direction := c.lowerLeftCorner.
		Add(c.horizontal.MultiplyScalar(u)).
		Add(c.vertical.MultiplyScalar(v)).
		Subtract(c.origin)
	return Ray{Origin: c.origin, Direction: direction}
}
