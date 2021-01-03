package fetch_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/spiegel-im-spiegel/ml/ecode"
	"github.com/spiegel-im-spiegel/ml/fetch"
)

func TestFetch(t *testing.T) {
	testCases := []struct {
		u   string
		err error
	}{
		{u: "https://git.io/vFR5M", err: nil},
		{u: "http://not.exist/", err: ecode.ErrInvalidRequest},
	}

	for _, tc := range testCases {
		resp, err := fetch.New(
			fetch.WithContext(context.Background()),
			fetch.WithHttpClient(http.DefaultClient),
		).Get(tc.u)
		if !errors.Is(err, tc.err) {
			t.Errorf("Fetch.Get(\"%v\") is \"%+v\", want \"%+v\"", tc.u, err, tc.err)
		} else if err == nil {
			resp.Body.Close()
		} else {
			fmt.Printf("Info: %+v\n", err)
		}
	}
}

/* Copyright 2021 Spiegel
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
