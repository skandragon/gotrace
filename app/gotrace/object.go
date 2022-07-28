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

// HitRecord is returned when an object is hit.
type HitRecord struct {
	P         Vector3
	Normal    Vector3
	T         float64
	FrontFace bool
	Material  Material
}

// SetFaceNormal will calculate the proper values for Normal and
// FrontFace.
func (hr *HitRecord) SetFaceNormal(r Ray, outwardNormal Vector3) {
	hr.FrontFace = r.Direction.Dot(outwardNormal) < 0
	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.MultiplyScalar(-1)
	}
}

// Object defines a world object that we can throw a ray at, and
// find out what color it should be.
type Object interface {
	Hit(r Ray, tMin float64, tMax float64) *HitRecord
}
