// Copyright 2020 FastWeGo
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

package template_message_test

import (
	"fmt"

	"github.com/icarephone/microapp"
	"github.com/icarephone/microapp/apis/template_message"
)

func ExampleSend() {
	var ctx *microapp.MicroApp

	payload := []byte("{}")
	resp, err := template_message.Send(ctx, payload)

	fmt.Println(resp, err)
}
