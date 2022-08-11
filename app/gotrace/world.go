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

// World defines our massive world.
type World struct {
	Camera   Camera
	Objects  []Hittable
	MaxDepth int
	TMin     float64
	TMax     float64
}

// Cast returns the color of a point, using the vector to define
// where it is cast into the scene.
func (w World) Cast(r Ray, depth int) Vector3 {
	depth--
	if depth < 0 {
		return Vector3{}
	}

	var closestHit *HitRecord
	smallestDistance := world.TMax
	for _, obj := range world.Objects {
		if hitRecord := obj.Hit(r, world.TMin, smallestDistance); hitRecord != nil {
			if closestHit == nil || closestHit.T > hitRecord.T {
				smallestDistance = hitRecord.T
				closestHit = hitRecord
			}
		}
	}
	if closestHit != nil {
		if propagate, scatteredRay, attentuation := closestHit.Material.Scatter(r, closestHit); propagate {
			return attentuation.Multiply(world.Cast(scatteredRay, depth-1))
		}
		return Vector3{}
	}

	// make unit vector so y is between -1.0 and 1.0
	unitDirection := r.Direction.Normalize()

	// scale t to be between 0.0 and 1.0
	t := 0.5 * (unitDirection.Y + 1.0)

	white := Vector3{1.0, 1.0, 1.0}
	blue := Vector3{0.5, 0.7, 1.0}

	return white.Lerp(blue, t)
}
