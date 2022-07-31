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
	"reflect"
	"testing"
)

var (
	one      = Vector3{1, 1, 1}
	retVec   Vector3
	retFloat float64
	retBool  bool
)

func BenchmarkNeg(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.Neg()
	}
	retVec = a
}

func BenchmarkAdd(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.Add(one)
	}
	retVec = a
}

func BenchmarkSubtract(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.Subtract(one)
	}
	retVec = a
}

func BenchmarkMultiply(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.Multiply(one)
	}
	retVec = a
}

func BenchmarkAddScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.AddScalar(1)
	}
	retVec = a
}

func BenchmarkSubtractScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.SubtractScalar(1)
	}
	retVec = a
}

func BenchmarkMultiplyScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.MultiplyScalar(1)
	}
	retVec = a
}

func BenchmarkDivideScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.DivideScalar(1)
	}
	retVec = a
}

func BenchmarkClamp(b *testing.B) {
	a := Vector3{3, 3, 3}
	for n := 0; n < b.N; n++ {
		a = a.Clamp(0, 9)
	}
	retVec = a
}

func BenchmarkLength(b *testing.B) {
	a := Vector3{3, 4, 5}
	len := 0.0
	for n := 0; n < b.N; n++ {
		len += a.Length()
	}
	retFloat = len
}

func BenchmarkLengthSquared(b *testing.B) {
	a := Vector3{3, 4, 5}
	len := 0.0
	for n := 0; n < b.N; n++ {
		len += a.LengthSquared()
	}
	retFloat = len
}

func BenchmarkDot(b *testing.B) {
	a := Vector3{3, 4, 5}
	len := 0.0
	for n := 0; n < b.N; n++ {
		len += a.Dot(a)
	}
	retFloat = len
}

func BenchmarkCross(b *testing.B) {
	a := Vector3{3, 4, 5}
	for n := 0; n < b.N; n++ {
		a = a.Cross(one)
	}
	retVec = a
}

func BenchmarkNormalize(b *testing.B) {
	a := Vector3{1, 2, 3}
	for n := 0; n < b.N; n++ {
		a = a.Normalize()
	}
	retVec = a
}

func BenchmarkLerp(b *testing.B) {
	a := Vector3{1000, 2000, 3000}
	for n := 0; n < b.N; n++ {
		a = a.Lerp(one, 0.01)
	}
	retVec = a
}

func BenchmarkRandomSphere(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = RandomUnitSphere()
	}
	retVec = a
}

func BenchmarkRandomUnitDisk(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = RandomUnitDisk()
	}
	retVec = a
}

func BenchmarkRandomVector(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = RandomVector()
	}
	retVec = a
}

func BenchmarkNearZeroVector(b *testing.B) {
	a := Vector3{1e-9, 1e-9, 4}
	ret := true
	for n := 0; n < b.N; n++ {
		ret = NearZeroVector(a)
	}
	retBool = ret
}

func TestVector3_Add(t *testing.T) {
	tests := []struct {
		name  string
		vec   Vector3
		other Vector3
		want  Vector3
	}{
		{
			"add",
			Vector3{-1, 0, 1},
			Vector3{2, 1, -2},
			Vector3{1, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vec.Add(tt.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Subtract(t *testing.T) {
	tests := []struct {
		name  string
		vec   Vector3
		other Vector3
		want  Vector3
	}{
		{
			"subtract",
			Vector3{-1, 0, 1},
			Vector3{2, 1, -2},
			Vector3{-3, -1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vec.Subtract(tt.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Multiply(t *testing.T) {
	tests := []struct {
		name  string
		vec   Vector3
		other Vector3
		want  Vector3
	}{
		{
			"multiply",
			Vector3{-1, 0, 1},
			Vector3{2, 1, -2},
			Vector3{-2, 0, -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vec.Multiply(tt.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
