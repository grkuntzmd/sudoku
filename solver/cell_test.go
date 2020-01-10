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

import "testing"

func TestXor(t *testing.T) {
	cl := cell(0b0110)
	val := cell(0b0010)
	assertTrue(t, cl.xor(val))
	assertEqual(t, cell(0b0100), cl)

	cl = cell(0b1000)
	val = cell(0b0110)
	assertFalse(t, cl.xor(val))
	assertEqual(t, cell(0b1000), cl)
}

func TestAnd(t *testing.T) {
	cl := cell(0b1010)
	val := cell(0b0110)
	assertTrue(t, cl.and(val))
	assertEqual(t, cell(0b0010), cl)

	cl = cell(0b0010)
	val = cell(0b0110)
	assertFalse(t, cl.and(val))
	assertEqual(t, cell(0b0010), cl)
}

func TestAndNot(t *testing.T) {
	cl := cell(0b1010)
	val := cell(0b0110)
	assertTrue(t, cl.andNot(val))
	assertEqual(t, cell(0b1000), cl)

	cl = cell(0b0001)
	val = cell(0b0110)
	assertFalse(t, cl.andNot(val))
	assertEqual(t, cell(0b0001), cl)
}

func TestReplace(t *testing.T) {
	cl := cell(0b1010)
	val := cell(0b0110)
	assertTrue(t, cl.replace(val))
	assertEqual(t, cell(0b0110), cl)

	cl = cell(0b0010)
	val = cell(0b0010)
	assertFalse(t, cl.replace(val))
	assertEqual(t, cell(0b0010), cl)
}

func TestOr(t *testing.T) {
	cl := cell(0b1010)
	val := cell(0b0110)
	assertTrue(t, cl.or(val))
	assertEqual(t, cell(0b1110), cl)

	cl = cell(0b1110)
	val = cell(0b0010)
	assertFalse(t, cl.or(val))
	assertEqual(t, cell(0b1110), cl)
}
