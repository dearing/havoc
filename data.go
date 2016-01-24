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
)

var Data = make([]byte, 0)

// DataSet sets the exported Data byte array to a given size
func DataSet(size int) {
	Data = make([]byte, size)
}

// ResetData calls DataSet(0) followed by FreeMemory
func ResetData() {
	DataSet(0)
	FreeMemory()
}

// DataFill will fill the current Data array with ones
func DataFill() {
	for i := 0; i < len(Data); i++ {
		Data[i] = 1
	}
}

// DataFillZero will fill the current Data array with zeros
func DataFillZero() {
	for i := 0; i < len(Data); i++ {
		Data[i] = 0
	}
}

// DataFillCrypto will fill the current Data array with random data
func DataFillCrypto() {
	rand.Read(Data)
}
