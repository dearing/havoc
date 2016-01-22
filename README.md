HAVOC
============

[![forthebadge](http://forthebadge.com/images/badges/just-plain-nasty.svg)](http://forthebadge.com)

Bad library for making a bad application.

```
go test -v -bench=.
=== RUN   TestSetMemory
--- PASS: TestSetMemory (0.00s)
=== RUN   TestResetMemory
--- PASS: TestResetMemory (0.00s)
=== RUN   TestSuite
	[2^01]: Data=0000000002,  ALLOC=0000112128,  SYS=0004983032
	[2^02]: Data=0000000004,  ALLOC=0000112192,  SYS=0004983032
	[2^03]: Data=0000000008,  ALLOC=0000112192,  SYS=0004983032
	[2^04]: Data=0000000016,  ALLOC=0000112192,  SYS=0004983032
	[2^05]: Data=0000000032,  ALLOC=0000112208,  SYS=0004983032
	[2^06]: Data=0000000064,  ALLOC=0000112240,  SYS=0004983032
	[2^07]: Data=0000000128,  ALLOC=0000112304,  SYS=0004983032
	[2^08]: Data=0000000256,  ALLOC=0000112432,  SYS=0004983032
	[2^09]: Data=0000000512,  ALLOC=0000112688,  SYS=0004983032
	[2^10]: Data=0000001024,  ALLOC=0000113200,  SYS=0004983032
	[2^11]: Data=0000002048,  ALLOC=0000114224,  SYS=0004983032
	[2^12]: Data=0000004096,  ALLOC=0000116272,  SYS=0004983032
	[2^13]: Data=0000008192,  ALLOC=0000120368,  SYS=0004983032
	[2^14]: Data=0000016384,  ALLOC=0000128560,  SYS=0004983032
	[2^15]: Data=0000032768,  ALLOC=0000144944,  SYS=0004983032
	[2^16]: Data=0000065536,  ALLOC=0000177712,  SYS=0004983032
	[2^17]: Data=0000131072,  ALLOC=0000243248,  SYS=0004983032
	[2^18]: Data=0000262144,  ALLOC=0000374320,  SYS=0004983032
	[2^19]: Data=0000524288,  ALLOC=0000636464,  SYS=0004983032
	[2^20]: Data=0001048576,  ALLOC=0001160752,  SYS=0004983032
--- PASS: TestSuite (0.16s)
PASS
BenchmarkMem-4	


    5000	    718237 ns/op
ok  	github.com/dearing/havoc	3.793s

```

TODO & Help Wanted
------------
 - see [issues]

Contributing
------------
1. Fork the repository on Github
2. Create a named feature branch (like `add_component_x`)
3. Write your change
4. Write tests for your change (if applicable)
5. Run the tests, ensuring they all pass
6. Submit a Pull Request using Github

License and Authors
-------------------
Author: Jacob Dearing // jacob.dearing@gmail.com

```
The MIT License (MIT)

Copyright (c) 2016 Jacob Dearing

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
