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

// NakedSingle removes digits from other items in a group (box, column, row) when a cell contains a solved value and returns true if it changes any cells.
func (gr *Grid) nakedSingle() bool {
	return gr.nakedSingleGroup(box) || gr.nakedSingleGroup(col) || gr.nakedSingleGroup(row)
}

func (gr *Grid) nakedSingleGroup(g group) (res bool) {
	for ci, c := range g.unit {
		for _, p1 := range c {
			val := *gr.pt(p1)
			if count[val] != 1 {
				continue
			}

			for _, p2 := range c {
				if p1 == p2 {
					continue
				}

				if (gr.pt(p2)).xor(val) {
					res = true
					if verbose >= 1 {
						fmt.Printf("nakedsingle: in %s %d cell %s allows only %s, removed from %s\n", g.name, ci, p1, val.digits(), p2)
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
