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

// BoxLine removes candidates. When a candidate within a column or row appears only in a single box that candidate can be removed from all cells in the box, other than those in the column or row. It returns true if it removes any digits.
func (gr *Grid) boxLine() bool {
	colSel := func(p point) int {
		return p.c
	}

	rowSel := func(p point) int {
		return p.r
	}

	return gr.boxLineGroup(col, colSel, rowSel, func(i, c int) int {
		return i*3 + c/3
	}) || gr.boxLineGroup(row, rowSel, colSel, func(i, c int) int {
		return c/3*3 + i
	})
}

func (gr *Grid) boxLineGroup(g group, major, minor func(point) int, boxSel func(index, ci int) int) (res bool) {
	for ci, c := range g.unit {
		var boxes [10][3]bool // True if box that contains a specific digit.
		for _, p := range c {
			val := gr[p.r][p.c]
			for d := uint16(1); d <= 9; d++ {
				if val&(1<<d) != 0 {
					boxes[d][minor(p)/3] = true
				}
			}
		}

		for d := uint16(1); d <= 9; d++ {
			var index int
			if boxes[d][0] && !boxes[d][1] && !boxes[d][2] {
				index = 0
			} else if !boxes[d][0] && boxes[d][1] && !boxes[d][2] {
				index = 1
			} else if !boxes[d][0] && !boxes[d][1] && boxes[d][2] {
				index = 2
			} else {
				continue
			}

			for pi := 0; pi < 9; pi++ {
				p := box.unit[boxSel(index, ci)][pi]
				if major(p) == major(c[index]) {
					continue
				}

				if gr[p.r][p.c].xor(1 << d) {
					res = true
					if verbose >= 1 {
						fmt.Printf("boxline: all %d's in %s %d appear in box %d removing from %s\n", d, g.name, ci, boxSel(index, ci), p)
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
