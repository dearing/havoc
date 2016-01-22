// Copyright © 2016 Jacob Dearing
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package havoc

import (
	"fmt"
	"runtime"
	"testing"
)

var m = runtime.MemStats{}

func TestSetMemory(t *testing.T) {
	SetMemory(1024)
	if len(Data) != 1024 {
		t.Error("Set memory failed to set Data.")
	}
}

func TestResetMemory(t *testing.T) {
	SetMemory(1024 * 1024)
	runtime.ReadMemStats(&m)
	old := m.HeapAlloc
	ResetMemory()

	if len(Data) != 0 {
		t.Error("Reset memory failed to clear Data.")
	}

	runtime.ReadMemStats(&m)
	if m.HeapAlloc >= old {
		t.Errorf("Reset memory failed release memory! %d >= %d\n", m.HeapAlloc, old)
	}

}
func TestSuite(t *testing.T) {
	for i := uint(1); i < 21; i++ {

		ResetMemory()
		SetMemory(1 << i)
		FillData()
		runtime.ReadMemStats(&m)
		fmt.Printf("\t[2^%02d]: Data=%010d,  ALLOC=%010d,  SYS=%010d\n", i, len(Data), m.Alloc, m.Sys)
		//fmt.Printf("\t[2^%02d]: Data=%010d, HALLOC=%010d, HREL=%010d\n", i, len(Data), m.HeapAlloc, m.HeapReleased)
	}
}

func BenchmarkMem(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ResetMemory()
		SetMemory(b.N)
		FillData()
	}
	println()
}
