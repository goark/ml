package ecode

import "fmt"

//ECode is error ECodeber for "spiegel-im-spiegel/ml/ecode/..." packages
type ECode int

const (
	ErrNullPointer ECode = iota + 1
	ErrNoImplement
	ErrInvalidRequest
)

var errMessage = map[ECode]string{
	ErrNullPointer:    "Null reference instance",
	ErrNoImplement:    "This style is not implementation",
	ErrInvalidRequest: "invalid request",
}

func (n ECode) Error() string {
	if s, ok := errMessage[n]; ok {
		return s
	}
	return fmt.Sprintf("unknown error (%d)", int(n))
}

/* Copyright 2019-2021 Spiegel
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
