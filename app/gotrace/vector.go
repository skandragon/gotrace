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

	"golang.org/x/exp/constraints"
)

// Vector3 holds the X, Y, and Z components of a 3-dimensional
// vector.
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func clamp[T constraints.Ordered](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// Clamp will ensure that the vector components are within the
// provided bounds.
func (v Vector3) Clamp(min, max float64) Vector3 {
	return Vector3{
		X: clamp(v.X, min, max),
		Y: clamp(v.Y, min, max),
		Z: clamp(v.Z, min, max),
	}
}

// Length returns the length of the vector.
func (v Vector3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the length*length of the vector.
func (v Vector3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Dot will calculate the dot product of this vector to "o"
func (v Vector3) Dot(o Vector3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Normalize returns the unit vector with the same direction
// as this vector.
func (v Vector3) Normalize() Vector3 {
	l := 1 / v.Length()
	return Vector3{v.X * l, v.Y * l, v.Z * l}
}

// Add will add two vectors together.  It will not modify
// either the source or "o" vector.
func (v Vector3) Add(o Vector3) Vector3 {
	return Vector3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// AddAccum will add directly into "v"
func (v *Vector3) AddAccum(o Vector3) *Vector3 {
	v.X += o.X
	v.Y += o.Y
	v.Z += o.Z
	return v
}

// Subtract will subtrace "o" from this vector, and return
// a new vector.
func (v Vector3) Subtract(o Vector3) Vector3 {
	return Vector3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

// Multiply will add two vectors together.  It will not modify
// either the source or "o" vector.
func (v Vector3) Multiply(o Vector3) Vector3 {
	return Vector3{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
}

// AddScalar adds the scalar component to each element of the
// vector, and returns a new vector.
// This is in effect a translation.
func (v Vector3) AddScalar(t float64) Vector3 {
	return Vector3{v.X + t, v.Y + t, v.Z + t}
}

// SubtractScalar subtracts the scalar component from each element
// of the vector, and returns a new vector.
// This is in effect a translation.
func (v Vector3) SubtractScalar(t float64) Vector3 {
	return Vector3{v.X - t, v.Y - t, v.Z - t}
}

// MultiplyScalar scales the vector, and returns a new vector.
func (v Vector3) MultiplyScalar(t float64) Vector3 {
	return Vector3{v.X * t, v.Y * t, v.Z * t}
}

// DivideScalar scales the vector, and returns a new vector.
func (v Vector3) DivideScalar(t float64) Vector3 {
	scale := 1 / t
	return Vector3{v.X * scale, v.Y * scale, v.Z * scale}
}

// Lerp interprolates between vectors "v" and "o" based on "t", which
// should be between [0.0, 1.0] inclusive.
func (v Vector3) Lerp(o Vector3, t float64) Vector3 {
	return Vector3{
		X: v.X*(1.0-t) + o.X*t,
		Y: v.Y*(1.0-t) + o.Y*t,
		Z: v.Z*(1.0-t) + o.Z*t,
	}
}

// Gamma2 applies a sqrt mod to the vector, which is assumed to
// be a color.
func (v Vector3) Gamma2() Vector3 {
	return Vector3{math.Sqrt(v.X), math.Sqrt(v.Y), math.Sqrt(v.Z)}
}

// RandomUnitSphere returns a normalized, randomly created unit vector.
func RandomUnitSphere() Vector3 {
	for {
		p := Vector3{
			X: rand.Float64()*2 - 1.0,
			Y: rand.Float64()*2 - 1.0,
			Z: rand.Float64()*2 - 1.0,
		}
		if p.LengthSquared() < 1 {
			return p.Normalize()
		}
	}
}

// NearZeroVector returns true if all elements are almost zero.
func NearZeroVector(v Vector3) bool {
	const s = 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}
