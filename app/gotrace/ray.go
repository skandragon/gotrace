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

// Ray defines a ray.
type Ray struct {
	Origin    Vector3
	Direction Vector3
}

// Point returns the point the ray points to.
func (r Ray) Point(t float64) Vector3 {
	b := r.Direction.MultiplyScalar(t)
	a := r.Origin
	return a.Add(b)
}

var world = World{
	Lights: []Light{},
	Objects: []Object{
		Sphere{Center: Vector3{0, 0, -1}, Radius: 0.5},
	},
}

// Cast returns the color of a point, using the vector to define
// where it is cast into the scene.
func (r Ray) Cast(tMin float64, tMax float64) Vector3 {

	var closestHit *HitRecord
	smallestDistance := tMax
	for _, obj := range world.Objects {
		if hitRecord := obj.Hit(r, tMin, smallestDistance); hitRecord != nil {
			if closestHit == nil || closestHit.T > hitRecord.T {
				smallestDistance = hitRecord.T
				closestHit = hitRecord
			}
		}
	}
	if closestHit != nil {
		return Vector3{1, 0.2, 0.1}
	}

	// make unit vector so y is between -1.0 and 1.0
	unitDirection := r.Direction.Normalize()

	// scale t to be between 0.0 and 1.0
	t := 0.5 * (unitDirection.Y + 1.0)

	white := Vector3{1.0, 1.0, 1.0}
	blue := Vector3{0.5, 0.7, 1.0}

	return white.Lerp(blue, t)
}
