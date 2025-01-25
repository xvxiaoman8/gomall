// Copyright 2024 CloudWeGo Authors
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

// Code generated by hertz generator.

package checkout

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xvxiaoman8/gomall/app/frontend/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Auth()}
}

func _checkoutMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkout0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkoutresultMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _checkoutwaitingMw() []app.HandlerFunc {
	// your code...
	return nil
}
