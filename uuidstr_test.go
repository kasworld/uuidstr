// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package uuidstr

import (
	"testing"

	"github.com/kasworld/g2rand"
)

func TestNew(t *testing.T) {
	uuid := New()
	t.Logf("%v", uuid)
	t.Logf("%v", Parse(uuid))
	for i := 0; i < 10; i++ {
		uuid := New()
		t.Logf("%v", uuid)
	}
}

func TestNewByFn(t *testing.T) {
	rnd := g2rand.New()
	uuid := NewByFn(rnd.Read)
	t.Logf("%v", uuid)
	t.Logf("%v", Parse(uuid))
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New()
	}
}

func BenchmarkNewByFn(b *testing.B) {
	rnd := g2rand.New()
	for i := 0; i < b.N; i++ {
		_ = NewByFn(rnd.Read)
	}
}
