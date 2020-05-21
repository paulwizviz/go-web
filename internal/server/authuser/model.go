// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authuser

import (
	"fmt"
	"strings"
)

const bearerPrefix = "Bearer"

type jwtToken string

func (t jwtToken) addBearerPrefix() jwtToken {
	prefix := fmt.Sprintf("%s %s", bearerPrefix, string(t))
	return jwtToken(prefix)
}

// RemoveBearerPrefix remove bearer proefix
func (t jwtToken) removeBearerPrefix() jwtToken {
	if strings.Contains(string(t), bearerPrefix) {
		result := strings.TrimLeft(string(t), bearerPrefix)
		result = strings.TrimSpace(result)
		return jwtToken(result)
	}
	return t
}
