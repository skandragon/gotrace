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

import "math"

type aabb struct {
	minimum Vector3
	maximum Vector3
}

// MakeAABB creates a new axis-aligned bounding box with the minimum and maximum points as corners.
func MakeAABB(min Vector3, max Vector3) Hittable {
	return aabb{
		minimum: min,
		maximum: max,
	}
}

func swap64(a, b float64) (float64, float64) {
	return b, a
}

func compare(vMin float64, vMax float64, origin float64, direction float64, tMin float64, tMax float64) bool {
	invD := 1.0 / direction
	t0 := (vMin - origin) * invD
	t1 := (vMax - origin) * invD
	if invD < 0.0 {
		t1, t0 = swap64(t0, t1)
	}
	t0 = math.Min(t0, tMin)
	t1 = math.Max(t1, tMax)
	return t1 > t0
}

var dummyHit = &HitRecord{}

func (s aabb) Hit(r Ray, tMin float64, tMax float64) *HitRecord {
	ret := compare(s.minimum.X, s.maximum.X, r.Origin.X, r.Direction.X, tMin, tMax) ||
		compare(s.minimum.Y, s.maximum.Y, r.Origin.Y, r.Direction.Y, tMin, tMax) ||
		compare(s.minimum.Z, s.maximum.Z, r.Origin.Z, r.Direction.Z, tMin, tMax)
	if !ret {
		return nil
	}
	return dummyHit
}
