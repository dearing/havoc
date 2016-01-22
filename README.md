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
	[2^01]: Data=0000000002, ALLOC=0000112544, SYS=0003901688
	[2^02]: Data=0000000004, ALLOC=0000112608, SYS=0003901688
	[2^03]: Data=0000000008, ALLOC=0000112608, SYS=0003901688
	[2^04]: Data=0000000016, ALLOC=0000112608, SYS=0003901688
	[2^05]: Data=0000000032, ALLOC=0000112624, SYS=0003901688
	[2^06]: Data=0000000064, ALLOC=0000112656, SYS=0003901688
	[2^07]: Data=0000000128, ALLOC=0000112720, SYS=0003901688
	[2^08]: Data=0000000256, ALLOC=0000112848, SYS=0003901688
	[2^09]: Data=0000000512, ALLOC=0000113104, SYS=0003901688
	[2^10]: Data=0000001024, ALLOC=0000113616, SYS=0003901688
	[2^11]: Data=0000002048, ALLOC=0000114640, SYS=0003901688
	[2^12]: Data=0000004096, ALLOC=0000116688, SYS=0003901688
	[2^13]: Data=0000008192, ALLOC=0000120784, SYS=0003901688
	[2^14]: Data=0000016384, ALLOC=0000128976, SYS=0003901688
	[2^15]: Data=0000032768, ALLOC=0000145360, SYS=0003901688
	[2^16]: Data=0000065536, ALLOC=0000178128, SYS=0003901688
	[2^17]: Data=0000131072, ALLOC=0000243664, SYS=0003901688
	[2^18]: Data=0000262144, ALLOC=0000374736, SYS=0003901688
	[2^19]: Data=0000524288, ALLOC=0000636880, SYS=0004983032
	[2^20]: Data=0001048576, ALLOC=0001161168, SYS=0004983032
	[2^21]: Data=0002097152, ALLOC=0002209744, SYS=0007145720
	[2^22]: Data=0004194304, ALLOC=0004310800, SYS=0011475192
	[2^23]: Data=0008388608, ALLOC=0008505104, SYS=0020134136
	[2^24]: Data=0016777216, ALLOC=0016893776, SYS=0037452024
	[2^25]: Data=0033554432, ALLOC=0033675272, SYS=0072087800
--- PASS: TestSuite (5.38s)
PASS
BenchmarkMem-4	


    5000	    752262 ns/op
ok  	github.com/dearing/havoc	9.182s

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
