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
	"crypto/rand"
	"runtime/debug"
)

var Data = make([]byte, 0)

// SetMemory sets the exported Data byte array to a given size
func SetMemory(size int) {
	Data = make([]byte, size)
}

// FreeMemory forces VM to release unused memory back to the system
func FreeMemory() {
	debug.FreeOSMemory()
}

// ResetMemory calls Setmemory(0) followed by FreeMemory
func ResetMemory() {
	SetMemory(0)
	FreeMemory()
}

// FillData will fill the current Data array with random data
func FillData() {
	rand.Read(Data)
}

// Forever runs forever.
func Forever() {
	for {
		// evah
	}
}
