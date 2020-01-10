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

import (
	"fmt"
	"strconv"
)

// XWing removes candidates. If in 2 columns, say 0 and 7, all instances of a particular digit, say 4, appear in the same two rows, say 4 and 6, then 1 of the 4's must be in (0, 4) or (0, 6) and the other in (7, 4) or (7, 6). Therefore all of the other 4's in those two rows can be removed. The same logic applies if rows and columns are swapped. It returns true if it changes any cells.
func (gr *Grid) xWing() bool {
	return gr.xWingGroup(col, row) || gr.xWingGroup(row, col)
}

func (gr *Grid) xWingGroup(majorGroup, minorGroup group) (res bool) {
	var digits [9][10]cell
	for ci, c := range majorGroup.unit {
		for pi, p := range c {
			val := *gr.pt(p)
			for d := 1; d <= 9; d++ {
				if val&(1<<d) != 0 {
					digits[ci][d] |= 1 << uint(pi)
				}
			}
		}
	}

	for d := 1; d <= 9; d++ {
		for c1i := 0; c1i < 8; c1i++ {
			for c2i := c1i + 1; c2i < 9; c2i++ {
				proto := digits[c1i][d]
				if count[proto] == 2 && proto == digits[c2i][d] {
					for minor := 1; minor <= 9; minor++ {
						if proto&(1<<minor) != 0 {
							for mi, m := range minorGroup.unit[minor] {
								if mi == c1i || mi == c2i {
									continue
								}
								// fmt.Printf("%s %s\n", showBits(*gr.pt(m)), showBits(cell(1<<d)))
								if gr.pt(m).andNot(cell(1 << d)) {
									res = true
									if verbose >= 1 {
										fmt.Printf("xwing: in %ss %d and %d, %d appears only in %s %d and 1 other; "+
											"removing from %s\n", majorGroup.name, c1i, c2i, d, minorGroup.name, minor, m)
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
		}
	}
	if res && verbose == 2 {
		gr.Display()
	}

	return
}

func showBits(c cell) string {
	return strconv.FormatUint(uint64(c), 2)
}
