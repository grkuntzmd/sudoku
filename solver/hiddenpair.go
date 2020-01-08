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

// HiddenPair removes other digits from a pair of cells in a group (box, column, row) when that pair contains the only occurrances of the digits in the group and returns true if it removes any digits.
func (gr *Grid) hiddenPair() bool {
	return gr.hiddenPairGroup(box) || gr.hiddenPairGroup(col) || gr.hiddenPairGroup(row)
}

func (gr *Grid) hiddenPairGroup(g group) (res bool) {
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

		for i1 := cell(1); i1 <= 8; i1++ {
			if len(digits[i1]) != 2 {
				continue
			}

			for i2 := i1 + 1; i2 <= 9; i2++ {
				if len(digits[i2]) != 2 {
					continue
				}

				if comparePointSlices(digits[i1], digits[i2]) {
					comb := cell(1<<i1 | 1<<i2)
					for k := 0; k < 2; k++ {
						p := digits[i1][k]
						if gr[p.r][p.c].replace(comb) {
							res = true
							if verbose >= 1 {
								fmt.Printf("hiddenpair: in %s %d limits %s to %s\n", g.name, ci, p, comb.digits())
							}
							if verbose >= 3 {
								gr.Display()
							}
						}
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
