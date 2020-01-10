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
	"math/bits"
	"strconv"
	"strings"
)

type Grid [9][9]cell

// Display prints a grid.
func (gr *Grid) Display() {
	width := gr.maxWidth() + 2

	part := strings.Repeat("\u2500", int(width*3))
	line := "\u251c" + strings.Join([]string{part, part, part}, "\u253c") + "\u2524"

	bars := strings.Repeat("\u2500", int(width)*3)
	fmt.Printf("\t%s%s%s%s%s%s%s\n", "\u250c", bars, "\u252c", bars, "\u252c", bars, "\u2510")
	for r := 0; r < 9; r++ {
		fmt.Print("\t\u2502")
		for c := 0; c < 9; c++ {
			var b strings.Builder
			for i := one; i <= 9; i++ {
				if gr[r][c]&(1<<i) != 0 {
					b.WriteString(strconv.Itoa(int(i)))
				}
			}
			if b.String() == "123456789" {
				fmt.Print(center(".", width))
			} else {
				fmt.Print(center(b.String(), width))
			}
			if c == 2 || c == 5 {
				fmt.Print("\u2502")
			}
		}
		fmt.Println("\u2502")
		if r == 2 || r == 5 {
			fmt.Println("\t" + line)
		}
	}
	fmt.Printf("\t%s%s%s%s%s%s%s\n", "\u2514", bars, "\u2534", bars, "\u2534", bars, "\u2518")
}

// digitPlaces returns an array of digits containing values where the bits (1 - 9) are set if the corresponding digit appears in that cell.
func (gr *Grid) digitPlaces(u [9]point) (places [10]cell) {
		for pi, p := range u {
			val := *gr.pt(p)
			for i := one; i <= 9; i++ {
				if val&(1<<i) != 0 {
					places[i] |= 1 << uint(pi)
				}
			}
		}

		return
}

// digitPoints builds a table of points that contain each digit.
func (gr *Grid) digitPoints(u [9]point) (points [10][]point) {
	for _, p := range u {
		val := *gr.pt(p)
		for d := one; d <= 9; d++ {
			if val&(1<<d) != 0 {
				points[d] = append(points[d], p)
			}
		}
	}

	return
}

func (gr *Grid) maxWidth() int {
	width := 0
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			count := bits.OnesCount16(uint16(gr[r][c]))
			if count == 9 {
				count = 1
			}
			if width < count {
				width = count
			}
		}
	}

	return width
}

func (gr *Grid) pt(p point) *cell {
	return &gr[p.r][p.c]
}

func (gr *Grid) solveLevel(maxLevel, level int, fns []func() bool) (int, bool) {
	for _, f := range fns {
		if f() {
			if maxLevel < level {
				maxLevel = level
			}
			return maxLevel, true
		}
	}

	return maxLevel, false
}
