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
	"testing"
)

func TestXor(t *testing.T) {
	cl := cell(0x6)
	val := cell(0x2)
	assertTrue(t, cl.xor(val))
	assertEqual(t, cell(0x4), cl)

	cl = cell(0x8)
	val = cell(0x6)
	assertFalse(t, cl.xor(val))
	assertEqual(t, cell(0x8), cl)
}

func TestAnd(t *testing.T) {
	cl := cell(0xa)
	val := cell(0x6)
	assertTrue(t, cl.and(val))
	assertEqual(t, cell(0x2), cl)

	cl = cell(0x2)
	val = cell(0x6)
	assertFalse(t, cl.and(val))
	assertEqual(t, cell(0x2), cl)
}

func TestReplace(t *testing.T) {
	cl := cell(0xa)
	val := cell(0x6)
	assertTrue(t, cl.replace(val))
	assertEqual(t, cell(0x6), cl)

	cl = cell(0x2)
	val = cell(0x2)
	assertFalse(t, cl.replace(val))
	assertEqual(t, cell(0x2), cl)
}

func TestOr(t *testing.T) {
	cl := cell(0xa)
	val := cell(0x6)
	assertTrue(t, cl.or(val))
	assertEqual(t, cell(0xe), cl)

	cl = cell(0xe)
	val = cell(0x2)
	assertFalse(t, cl.or(val))
	assertEqual(t, cell(0xe), cl)
}

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func assertFalse(t *testing.T, a bool) {
	if a {
		t.Fatal("should be false")
	}
}

func assertTrue(t *testing.T, a bool) {
	if !a {
		t.Fatal("should be true")
	}
}
