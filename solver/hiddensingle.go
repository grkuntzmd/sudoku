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

// HiddenSingle solves a cell if it contains the only instance of a digit within its group (box, column, row) and returns true if it removes any digits.
func (gr *Grid) hiddenSingle() bool {
	return gr.hiddenSingleGroup(box) || gr.hiddenSingleGroup(col) || gr.hiddenSingleGroup(row)
}

func (gr *Grid) hiddenSingleGroup(g group) (res bool) {
	for ci, c := range g.unit {
		var digits [10][]point // Points that contain a specific digit.
		for _, p := range c {
			val := gr[p.r][p.c]
			for i := one; i <= 9; i++ {
				if val&(1<<i) != 0 {
					digits[i] = append(digits[i], p)
				}
			}
		}

		for i := one; i <= 9; i++ {
			if len(digits[i]) == 1 {
				p := digits[i][0]
				if gr[p.r][p.c].replace(1 << i) {
					res = true
					if verbose >= 1 {
						fmt.Printf("hiddensingle: in %s %d set %s to %d\n", g.name, ci, p, i)
					}
					if verbose >= 3 {
						gr.Display()
					}
				}
			}
		}
	}
	if res && verbose >= 2 {
		gr.Display()
	}

	return
}
