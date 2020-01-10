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

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Fatalf("%+v != %+v", a, b)
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
