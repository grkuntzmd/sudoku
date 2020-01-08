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

type (
	Grid [9][9]cell

	Game struct {
		Orig Grid
		Curr Grid
	}
)

// ParseGrid parses a string of digits and dots into a game structure containing two grids: the original set up and current set up. It panics on any illegal input.
func ParseGrid(input string) *Game {
	var game Game
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if input[r*9+c] == '.' {
				for i := uint16(1); i <= 9; i++ {
					game.Orig[r][c] |= 1 << i
				}
			} else {
				d, err := strconv.Atoi(string(input[r*9+c]))
				if err != nil {
					panic(fmt.Sprintf("illegal character in input grid: %s (\"%c\")", input, input[r*9+c]))
				}
				game.Orig[r][c] |= 1 << uint16(d)
			}
		}
	}

	return &game
}

// Solve solves the current grid of the given game.
func (ga *Game) Solve() (int, bool) {
	ga.Orig.Display()

	// Copy the original grid to the current grid so that the original is preserved.
	ga.Curr = ga.Orig

	maxLevel := Level0
	for !ga.Curr.validate() {
		var ok bool
		maxLevel, ok = ga.Curr.solveLevel(maxLevel, Level0, []func() bool{
			ga.Curr.nakedSingle,
			ga.Curr.hiddenSingle,
		})
		if ok {
			continue
		}

		maxLevel, ok = ga.Curr.solveLevel(maxLevel, Level1, []func() bool{
			ga.Curr.nakedPair,
			ga.Curr.hiddenPair,
			ga.Curr.boxLine,
			ga.Curr.pointingLine,
		})
		if ok {
			continue
		}

		maxLevel, ok = ga.Curr.solveLevel(maxLevel, Level2, []func() bool{})
		if ok {
			continue
		}

		maxLevel, ok = ga.Curr.solveLevel(maxLevel, Level3, []func() bool{})
		if ok {
			continue
		}

		maxLevel, ok = ga.Curr.solveLevel(maxLevel, Level4, []func() bool{})
		if ok {
			continue
		}

		break
	}

	valid := ga.Curr.validate()
	if !valid {
		fmt.Println("Not solved")
	}
	ga.Curr.Display()
	fmt.Printf("maxLevel: %d\n", maxLevel)

	return maxLevel, valid
}

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
