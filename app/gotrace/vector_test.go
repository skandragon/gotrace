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

import "testing"

var (
	one = Vector3{1, 1, 1}
)

func BenchmarkAdd(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.Add(one)
	}
}

func BenchmarkSubtract(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.Subtract(one)
	}
}

func BenchmarkMultiply(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.Multiply(one)
	}
}

func BenchmarkAddScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.AddScalar(1)
	}
}

func BenchmarkSubtractScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.SubtractScalar(1)
	}
}

func BenchmarkMultiplyScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.MultiplyScalar(1)
	}
}

func BenchmarkDivideScalar(b *testing.B) {
	a := Vector3{}
	for n := 0; n < b.N; n++ {
		a = a.DivideScalar(1)
	}
}
