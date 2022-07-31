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

// Ray defines a ray.
type Ray struct {
	Origin    Vector3
	Direction Vector3
}

// Point returns the point the ray points to.
func (r Ray) Point(t float64) Vector3 {
	b := r.Direction.MultiplyScalar(t)
	return r.Origin.Add(b)
}
