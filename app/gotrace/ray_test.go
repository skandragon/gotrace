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

func TestRay_Point(t *testing.T) {
	type fields struct {
		Origin    Vector3
		Direction Vector3
	}
	type args struct {
		t float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector3
	}{
		{
			"multiplier 1",
			fields{
				Origin:    Vector3{1, 2, 3},
				Direction: Vector3{1, 1, 1},
			},
			args{
				t: 1.0,
			},
			Vector3{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Ray{
				Origin:    tt.fields.Origin,
				Direction: tt.fields.Direction,
			}
			if got := r.Point(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ray.Point() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testVec Vector3
)

func BenchmarkRay_Point(b *testing.B) {
	ray := Ray{Origin: Vector3{1, 2, 3}, Direction: Vector3{1, 2, 3}}
	for n := 0; n < b.N; n++ {
		ray.Origin = ray.Point(1.0)
	}
	testVec = ray.Origin
}
