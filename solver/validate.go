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

func (gr *Grid) validate() bool {
	return gr.validateGroup(box) && gr.validateGroup(col) && gr.validateGroup(row)
}

func (gr *Grid) validateGroup(g group) bool {
	for _, c := range g.unit {
		var cells [10]int
		for _, p := range c {
			val := *gr.pt(p)
			if count[val] == 0 {
				gr.Display()
				panic("empty cell")
			}
			for d := 1; d <= 9; d++ {
				if val&(1<<d) != 0 {
					cells[d]++
				}
			}
		}

		for d := 1; d <= 9; d++ {
			if cells[d] != 1 {
				return false
			}
		}
	}

	return true
}
