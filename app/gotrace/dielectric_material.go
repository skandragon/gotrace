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

// DielectricMaterial defines a matt material.
type DielectricMaterial struct {
	indexOfRefraction float64
}

// NewDielectricMaterial returns a new material.
func NewDielectricMaterial(indexOfRefraction float64) DielectricMaterial {
	return DielectricMaterial{indexOfRefraction: indexOfRefraction}
}

func refract(uv Vector3, n Vector3, etaiOverEtat float64, cosTheta float64) Vector3 {
	rOutPerp := uv.Add(n.MultiplyScalar(cosTheta)).MultiplyScalar(etaiOverEtat)
	rOutParallel := n.MultiplyScalar(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}

func reflectance(cosine float64, refIndex float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - refIndex) / (1 + refIndex)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

// Scatter calculates how rays should scatter from this material.
func (m DielectricMaterial) Scatter(r Ray, hr *HitRecord) (bool, Ray, Vector3) {
	refractionRatio := m.indexOfRefraction
	if hr.FrontFace {
		refractionRatio = 1.0 / m.indexOfRefraction
	}
	unitDirection := r.Direction.Normalize()
	cosTheta := math.Min(unitDirection.Neg().Dot(hr.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1.0
	var direction Vector3
	if cannotRefract || reflectance(cosTheta, refractionRatio) > rand.Float64() {
		direction = reflectRay(unitDirection, hr.Normal)
	} else {
		direction = refract(unitDirection, hr.Normal, refractionRatio, cosTheta)
	}
	return true, Ray{hr.P, direction, r.Time}, Vector3{1, 1, 1}
}
