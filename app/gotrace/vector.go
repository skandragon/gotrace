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

// Vector3 holds the X, Y, and Z components of a 3-dimensional
// vector.
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// Length returns the length of the vector.
func (v Vector3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Dot will calculate the dot product of this vector to "o"
func (v Vector3) Dot(o Vector3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Normalize returns the unit vector with the same direction
// as this vector.
func (v Vector3) Normalize() Vector3 {
	l := v.Length()
	return Vector3{v.X / l, v.Y / l, v.Z / l}
}

// Add will add two vectors together.  It will not modify
// either the source or "o" vector.
func (v Vector3) Add(o Vector3) Vector3 {
	return Vector3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

// Subtract will subtrace "o" from this vector, and return
// a new vector.
func (v Vector3) Subtract(o Vector3) Vector3 {
	return Vector3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
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
	return Vector3{v.X / t, v.Y / t, v.Z / t}
}

// Lerp interprolates between vectors "v" and "o" based on "t", which
// should be between [0.0, 1.0] inclusive.
func (v Vector3) Lerp(o Vector3, t float64) Vector3 {
	return v.MultiplyScalar(1.0 - t).Add(o.MultiplyScalar(t))
}
