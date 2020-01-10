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

// yWing removes candidates. If a cell has two candidates (AB) and in the same column as AB is a cell containing AC and in the same row as AB is a cell containing BC, then if AB evaluates to A, then AC must be C, and similarly, if AB evaluates to B, then BC must be B. Therefore, the cell at the other intersection of AC and BC cannot contain a C, since C must appear in either AC or BC, therefore C can be removed from that last cell. It returns true if it changes any cells.
func (gr *Grid) yWing() bool {
	res := false
	for _, b := range box.unit {
		for _, p := range b {
			cl := *gr.pt(p)
			if count[cl] != 2 {
				continue
			}

			candidates := gr.findCandidates(p, p)

			for c1i := 0; c1i < len(candidates)-2; c1i++ {
				c1 := candidates[c1i]
				v1 := *gr.pt(c1)
				n1 := neighbors(c1)
				for c2i := c1i + 1; c2i < len(candidates)-1; c2i++ {
					c2 := candidates[c2i]
					v2 := *gr.pt(c2)
					if count[v1|v2] != 3 || cl&v1|cl&v2 != cl {
						continue
					}

					n2 := neighbors(c2)
					var overlap [9][9]bool
					for r := 0; r < 9; r++ {
						for c := 0; c < 9; c++ {
							if n1[r][c] && n2[r][c] {
								overlap[r][c] = true
							}
						}
					}
					overlap[p.r][p.c] = false
					for r := 0; r < 9; r++ {
						for c := 0; c < 9; c++ {
							if overlap[r][c] {
								bits := (*gr.pt(c1) | *gr.pt(c2)) &^ cl
								if (&gr[r][c]).andNot(bits) {
									res = true
									if verbose >= 1 {
										fmt.Printf("ywing: %s, %s, %s causes clearing %s from (%d, %d)\n",
											p, c1, c2, bits.digits(), r, c)
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

	return res
}

func (gr *Grid) findCandidates(p, skip point) []point {
	var ps []point
	ps = append(ps, gr.findCandidatesUnit(&box.unit[boxOf(p.r, p.c)], p, skip)...)
	ps = append(ps, gr.findCandidatesUnit(&col.unit[p.c], p, skip)...)
	ps = append(ps, gr.findCandidatesUnit(&row.unit[p.r], p, skip)...)
	return ps
}

func (gr *Grid) findCandidatesUnit(u *[9]point, cp, skip point) []point {
	var ps []point
	cl := *gr.pt(cp)
	for _, p := range u {
		if p == skip {
			continue
		}

		cand := *gr.pt(p)
		if count[cand] != 2 || count[cl&cand] != 1 {
			continue
		}
		ps = append(ps, p)
	}

	return ps
}

func neighbors(pt point) *[9][9]bool {
	var n [9][9]bool
	for _, u := range []*[9]point{&box.unit[boxOf(pt.r, pt.c)], &col.unit[pt.c], &row.unit[pt.r]} {
		for _, p := range u {
			if p == pt {
				continue
			}
			n[p.r][p.c] = true
		}
	}

	return &n
}
