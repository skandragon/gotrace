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
)

// Sphere defines a sphere.
type Sphere struct {
	Center   Vector3
	Radius   float64
	Material Material
}

// Hit calculates the hit of a sphere.
func (s Sphere) Hit(r Ray, tMin float64, tMax float64) *HitRecord {
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.LengthSquared()
	bHalf := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := bHalf*bHalf - a*c

	if discriminant < 0 {
		return nil
	}
	sqrtd := math.Sqrt(discriminant)

	root := (-bHalf - sqrtd) / a
	if root < tMin || root > tMax {
		root = (-bHalf + sqrtd) / a
	}
	if root < tMin || root > tMax {
		return nil
	}

	hitPoint := r.Point(root)
	outwardNormal := hitPoint.Subtract(s.Center).DivideScalar(s.Radius)
	hr := &HitRecord{T: root, P: hitPoint, Material: s.Material}
	hr.SetFaceNormal(r, outwardNormal)
	return hr
}
