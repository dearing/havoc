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
var size = 1024

func TestDataSet(t *testing.T) {

	DataSet(size)

	if len(Data) != size {
		t.Error("Set memory failed to set Data.")
	}
}

func TestDataFill(t *testing.T) {

	DataSet(size)

	DataFill()
	for _, v := range Data {
		if v == 0 {
			t.Error("DATA was not filled (non-zero).")
			return
		}
	}

}

func TestDataFillCrypto(t *testing.T) {

	DataSet(size)

	DataFillCrypto()

}

func TestDataReset(t *testing.T) {

	DataSet(size)
	DataFillCrypto()

	runtime.ReadMemStats(&m)
	old := m.HeapAlloc
	DataReset()

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
		DataReset()
		DataSet(1 << i)
		DataFillCrypto()
		runtime.ReadMemStats(&m)
		fmt.Printf("\t[2^%02d]: Data=%010d,  ALLOC=%010d,  SYS=%010d\n", i, len(Data), m.Alloc, m.Sys)
	}
}

func BenchmarkFillData(b *testing.B) {

	for i := 0; i < b.N; i++ {
		DataReset()
		DataSet(b.N)
		DataFill()
	}
}

func BenchmarkZeroFillData(b *testing.B) {

	for i := 0; i < b.N; i++ {
		DataReset()
		DataSet(b.N)
		DataFillZero()
	}
}

func BenchmarkCryptoFillData(b *testing.B) {

	for i := 0; i < b.N; i++ {
		DataReset()
		DataSet(b.N)
		DataFillCrypto()
	}
}
