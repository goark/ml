package errs

import (
	"testing"
)

func TestCause(t *testing.T) {
	testCases := []struct {
		err   error
		cause error
	}{
		{err: nil, cause: nil},
		{err: ErrNoImplement, cause: ErrNoImplement},
		{err: Wrap(ErrNoImplement, "wrapping error"), cause: ErrNoImplement},
	}

	for _, tc := range testCases {
		res := Cause(tc.err)
		if res != tc.cause {
			t.Errorf("Cause in \"%v\" == \"%v\", want \"%v\"", tc.err, res, tc.cause)
		}
	}
}

/* Copyright 2019 Spiegel
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* 	http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
