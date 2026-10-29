package facade

import (
	"os"

	"github.com/goark/gocli/config"
)

const logFile = "history"

func historyPath() string {
	return config.Path(appNameShort, logFile)
}

func mkdirHistory() error {
	return os.MkdirAll(config.Dir(appNameShort), 0700)
}

/* Copyright 2017-2026 Spiegel
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
