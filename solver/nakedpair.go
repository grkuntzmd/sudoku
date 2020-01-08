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

// NakedPair checks a group for 2 cells with the same pair of values. If present, those values can be removed from all other cells in the group. It returns true if it removes any digits.
func (gr *Grid) nakedPair() bool {
	return gr.nakedPairGroup(box) || gr.nakedPairGroup(col) || gr.nakedPairGroup(row)
}

func (gr *Grid) nakedPairGroup(g group) (res bool) {
	for ci, c := range g.unit {
	outer:
		for _, p1 := range c {
			val := gr[p1.r][p1.c]
			if count[val] != 2 {
				continue
			}

			for _, p2 := range c {
				if p1 == p2 {
					continue
				}

				cell2 := &gr[p2.r][p2.c]
				if *cell2 != val {
					continue
				}

				for _, p3 := range c {
					if p1 == p3 || p2 == p3 {
						continue
					}

					if gr[p3.r][p3.c].xor(val) {
						res = true
						if verbose >= 1 {
							fmt.Printf("nakedpair: in %s %d removed %s from %s\n", g.name, ci, val.digits(), p3)
						}
						if verbose >= 3 {
							gr.Display()
						}
					}
				}
				continue outer
			}
		}
	}
	if res && verbose >= 2 {
		gr.Display()
	}

	return
}
