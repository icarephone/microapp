// Copyright 2020 Icarephone
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

package data_caching_test

import (
	"fmt"
	"net/url"

	"github.com/icarephone/microapp"
	"github.com/icarephone/microapp/apis/data_caching"
)

func ExampleSetUserStorage() {
	var ctx *microapp.MicroApp

	payload := []byte("{}")
	params := url.Values{}
	resp, err := data_caching.SetUserStorage(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleRemoveUserStorage() {
	var ctx *microapp.MicroApp

	payload := []byte("{}")
	params := url.Values{}
	resp, err := data_caching.RemoveUserStorage(ctx, payload, params)

	fmt.Println(resp, err)
}
