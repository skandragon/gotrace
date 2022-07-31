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

// ReflectiveMaterial defines a matt material.
type ReflectiveMaterial struct {
	albedo Vector3
	fuzz   float64
}

// NewReflectiveMaterial returns a new material with the provided
// color.
func NewReflectiveMaterial(color Vector3, fuzz float64) ReflectiveMaterial {
	return ReflectiveMaterial{albedo: color, fuzz: fuzz}
}

// Scatter calculates how rays should scatter from this material.
func (m ReflectiveMaterial) Scatter(r Ray, hr *HitRecord) (bool, Ray, Vector3) {
	reflected := reflectRay(r.Direction.Normalize(), hr.Normal).
		Add(RandomUnitSphere().MultiplyScalar(m.fuzz))
	scattered := Ray{hr.P, reflected}
	return scattered.Direction.Dot(hr.Normal) > 0, scattered, m.albedo
}

func reflectRay(v Vector3, n Vector3) Vector3 {
	return v.Subtract(n.MultiplyScalar(v.Dot(n) * 2))
}
