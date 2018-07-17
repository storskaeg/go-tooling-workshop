// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sum

import "testing"

func TestRecursive(t *testing.T) {
	// Implement the body of this test, calling recursive.
	happyVals := []int{0, 1, 2, 3}
	sadVals := []int{}

	gotHappy := recursive(happyVals)
	gotSad := recursive(sadVals)

	wantHappy := 6
	wantSad := 0

	if gotHappy != wantHappy {
		t.Errorf("want: %v; got: %v", wantHappy, gotHappy)
	}

	if gotSad != wantSad {
		t.Errorf("want: %v; got: %v", wantSad, gotSad)
	}
}
