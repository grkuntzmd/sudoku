/*
 * Copyright © 2020, G.Ralph Kuntz, MD.
 *
 * Licensed under the Apache License, Version 2.0(the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIC
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package solver

import "fmt"

// PointingLine removes candidates. When a candidate within a box appears only in a single column or row, that candidate can be removed from all cells in the column or row outside of the box. It returns true if it changes any cells.
func (gr *Grid) pointingLine() bool {
	return gr.pointingLineGroup(func(p point) *[9]point {
		return &col.unit[p.c]
	}, func(p point) int {
		return p.c
	}) || gr.pointingLineGroup(func(p point) *[9]point {
		return &row.unit[p.r]
	}, func(p point) int {
		return p.r
	})
}

func (gr *Grid) pointingLineGroup(sel func(point) *[9]point, axis func(point) int) bool {
	res := false
	for bi, b := range box.unit {
		points := gr.digitPoints(b)

		// Loop through the digits and determine if all of them are on the same line (col or row). If so, then all other cells in that line that are not in the current box can have those digits removed.
	outer:
		for d := 1; d <= 9; d++ {
			a := axis(points[d][0])
			for _, p := range points[d][1:] {
				if axis(p) != a {
					continue outer
				}
			}

			for _, p := range sel(points[d][0]) {
				if p.r/3*3+p.c/3 == bi {
					continue
				}

				if gr.pt(p).xor(1 << d) {
					res = true
					if verbose >= 1 {
						fmt.Printf("pointingline: in box %d removing %d from %s\n", bi, d, p)
					}
					if verbose >= 3 {
						gr.Display()
					}
				}
			}
		}
	}
	if res && verbose == 2 {
		gr.Display()
	}

	return res
}
