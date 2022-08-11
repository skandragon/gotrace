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

type movingSphere struct {
	Center0   Vector3
	Center1   Vector3
	Time0     float64
	Time1     float64
	midcenter Vector3
	Radius    float64
	Material  Material
}

// NewMovingSphere returns a well constructed MovingSphere with some small speed improvements.
func NewMovingSphere(c0 Vector3, c1 Vector3, t0 float64, t1 float64, r float64, mat Material) Hittable {
	return movingSphere{
		Center0:   c0,
		Center1:   c1,
		midcenter: c1.Subtract(c0),
		Time0:     t0,
		Time1:     t1,
		Radius:    r,
		Material:  mat,
	}
}

func (s movingSphere) center(t float64) Vector3 {
	tt := (t - s.Time0) / (s.Time1 - s.Time0)
	return s.Center0.Add(s.midcenter.MultiplyScalar(tt))
}

// Hit calculates the hit of a sphere.
func (s movingSphere) Hit(r Ray, tMin float64, tMax float64) *HitRecord {
	oct := s.center(r.Time)
	oc := r.Origin.Subtract(oct)
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
		if root < tMin || root > tMax {
			return nil
		}
	}

	hitPoint := r.Point(root)
	outwardNormal := hitPoint.Subtract(oct).DivideScalar(s.Radius)
	hr := &HitRecord{T: root, P: hitPoint, Material: s.Material}
	hr.SetFaceNormal(r, outwardNormal)
	return hr
}
