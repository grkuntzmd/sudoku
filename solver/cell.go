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
	"strconv"
	"strings"
)

type (
	// Cell represents a single cell in the grid. Bits 1 through 9 are represented as a value where a bit
	// is set if that digit is a candidate for the cell.
	cell uint16
)

// String prints the string representation of a Cell.
func (c cell) String() string {
	var s strings.Builder
	for i := 1; i <= 9; i++ {
		if c&(1<<i) != 0 {
			s.WriteString(strconv.Itoa(int(i))) // nolint
		}
	}
	return s.String()
}

// and ANDs the current cell with the other cell and returns true if the current cell changes.
func (c *cell) and(o cell) bool {
	prev := *c
	*c &= o
	return *c != prev
}

// andNot ANDs the current cell with the complement of the other cell and returns true if the current cell changes.
func (c *cell) andNot(o cell) bool {
	prev := *c
	*c &= ^o
	return *c != prev
}

func (c cell) digits() string {
	var d []string
	for i := 1; i <= 9; i++ {
		if c&(1<<i) != 0 {
			d = append(d, strconv.Itoa(int(i)))
		}
	}
	return strings.Join(d, ", ")
}

// or ORs the current cell with the other cell and returns true if the current cell changes.
func (c *cell) or(o cell) bool {
	prev := *c
	*c |= o
	return *c != prev
}

// replace replaces the current cell with the other cell and returns true if the current cell changes.
func (c *cell) replace(o cell) bool {
	prev := *c
	*c = o
	return *c != prev
}

// xor XORs the current cell with the other cell and returns true if the current cell changes.
func (c *cell) xor(o cell) bool {
	prev := *c
	*c &^= o
	return *c != prev
}
