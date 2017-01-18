// Copyright (c) 2016 - Max Ekman <max@looplab.se>
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

package eventhorizon

import (
	"context"
	"testing"
)

func TestContextNamespace(t *testing.T) {
	ctx := context.Background()

	if ns := Namespace(ctx); ns != DefaultNamespace {
		t.Error("the namespace should be the default:", ns)
	}

	ctx = WithNamespace(ctx, "ns")
	if ns := Namespace(ctx); ns != "ns" {
		t.Error("the namespace should be correct:", ns)
	}
}
