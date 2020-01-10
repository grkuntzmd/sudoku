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

// Nakedtriple checks a group for 3 cells with the same triple of values. If present, those values can be removed from all other cells in the group. It returns true if it changes any cells.
func (gr *Grid) nakedTriple() bool {
	return gr.nakedTripleGroup(box) || gr.nakedTripleGroup(col) || gr.nakedTripleGroup(row)
}

func (gr *Grid) nakedTripleGroup(g group) (res bool) {
	for ci, c := range g.unit {
		for p1i := 0; p1i < 7; p1i++ {
			p1 := c[p1i]
			cell1 := *gr.pt(p1)
			if count[cell1] == 1 || count[cell1] > 3 {
				continue
			}

			for p2i := p1i + 1; p2i < 8; p2i++ {
				p2 := c[p2i]
				cell2 := *gr.pt(p2)
				if count[cell1|cell2] == 1 || count[cell1|cell2] > 3 {
					continue
				}

				for p3i := p2i + 1; p3i < 9; p3i++ {
					p3 := c[p3i]
					cell3 := *gr.pt(p3)
					if count[cell3] == 1 {
						continue
					}

					comb := cell1 | cell2 | cell3
					if count[comb] > 3 {
						continue
					}

					for pi, p := range c {
						if pi == p1i || pi == p2i || pi == p3i {
							continue
						}

						if gr.pt(p).xor(comb) {
							res = true
							if verbose >= 1 {
								fmt.Printf("nakedtriple: in %s %d (%s, %s, %s) removing %s from %s\n", g.name, ci, p1, p2, p3, comb.digits(), p)
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
	if res && verbose == 2 {
		gr.Display()
	}

	return
}
