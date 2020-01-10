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
)

// HiddenQuad removes other digits from a quad of cells in a group (box, column, row) when that quad contains the only occurrances of the digits in the group. It returns true if it changes any cells.
func (gr *Grid) hiddenQuad() bool {
	return gr.hiddenQuadGroup(box) || gr.hiddenQuadGroup(col) || gr.hiddenQuadGroup(row)
}

func (gr *Grid) hiddenQuadGroup(g group) (res bool) {
	for ci, c := range g.unit {
		places := gr.digitPlaces(c)

		for i1 := 1; i1 <= 6; i1++ {
			p1 := places[i1]
			c1 := count[p1]
			if c1 == 1 || c1 > 4 {
				continue
			}

			for i2 := i1 + 1; i2 <= 7; i2++ {
				p2 := places[i2]
				c2 := count[p2]
				if c2 == 1 || c2 > 4 || count[p1|p2] > 4 {
					continue
				}

				for i3 := i2 + 1; i3 <= 8; i3++ {
					p3 := places[i3]
					c3 := count[p3]
					if c3 == 1 || c3 > 4 || count[p1|p2|p3] > 4 {
						continue
					}

					for i4 := i3 + 1; i4 <= 9; i4++ {
						p4 := places[i4]
						c4 := count[p4]
						comb := p1 | p2 | p3 | p4
						if c4 == 1 || c4 > 4 || count[comb] != 4 {
							continue
						}

						bits := cell(1<<i1 | 1<<i2 | 1<<i3 | 1<<i4)
						for pi, p := range c {
							if comb&(1<<pi) != 0 {
								if gr.pt(p).and(bits) {
									res = true
									if verbose >= 1 {
										fmt.Printf("hiddenquad: in %s %d limits %s to %s\n", g.name, ci, p, bits.digits())
									}
									if verbose >= 3 {
										gr.Display()
									}
								}
							}
						}
						if res {
							gr.Display()
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
