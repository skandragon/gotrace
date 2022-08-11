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

// LambertianMaterial defines a matt material.
type LambertianMaterial struct {
	albedo Vector3
}

// NewLambertianMaterial returns a new material with the provided
// color.
func NewLambertianMaterial(color Vector3) LambertianMaterial {
	return LambertianMaterial{albedo: color}
}

// Scatter calculates how rays should scatter from this material.
func (m LambertianMaterial) Scatter(r Ray, hr *HitRecord) (bool, Ray, Vector3) {
	scatterDirection := hr.Normal.Add(RandomUnitSphere())
	if NearZeroVector(scatterDirection) {
		scatterDirection = hr.Normal
	}
	return true, Ray{hr.P, scatterDirection, r.Time}, m.albedo
}
