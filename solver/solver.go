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
	"flag"
	"fmt"
	"math/bits"
)

// Maximum difficult level found on solving.
const (
	Level0 = iota
	Level1
	Level2
	Level3
	Level4
)

type (
	group struct {
		name string
		unit
	}
	unit [9][9]point
)

var (
	verbose int

	box group = group{name: "box"} // These are all of the coordinates in a box (first dimension).
	col group = group{name: "col"} // These are all of the coordinates in a column (first dimension).
	row group = group{name: "row"} // These are all of the coordinates in a row (first dimension).

	count [1024]uint8
)

func init() {
	flag.IntVar(&verbose, "v", 0, "Verbose level (higher generates more output)")

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			p := point{r, c}
			box.unit[boxOf(r, c)][r%3*3+c%3] = p
			col.unit[c][r] = p
			row.unit[r][c] = p
		}
	}

	for i := uint16(0); i < 1024; i++ {
		count[i] = uint8(bits.OnesCount16(i))
	}
}

func boxOf(r, c int) int {
	return r/3*3 + c/3
}

func center(s string, w int) string {
	lead := (w - len(s)) / 2
	follow := w - len(s) - lead
	return fmt.Sprintf("%*s%*s", lead+len(s), s, follow, " ")
}

func comparePointSlices(s1, s2 []point) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
