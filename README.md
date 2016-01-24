HAVOC
============

[![forthebadge](http://forthebadge.com/images/badges/just-plain-nasty.svg)](http://forthebadge.com)

Bad library for making a bad application.

- see [havoc_server] for the webserver!

```
go test -v -bench=.
=== RUN   TestDataSet
--- PASS: TestDataSet (0.00s)
=== RUN   TestDataFill
--- PASS: TestDataFill (0.00s)
=== RUN   TestDataFillCrypto
--- PASS: TestDataFillCrypto (0.00s)
=== RUN   TestDataReset
--- PASS: TestDataReset (0.00s)
=== RUN   TestSuite
	[2^01]: Data=0000000002,  ALLOC=0000114992,  SYS=0003901688
	[2^02]: Data=0000000004,  ALLOC=0000115056,  SYS=0003901688
	[2^03]: Data=0000000008,  ALLOC=0000115056,  SYS=0003901688
	[2^04]: Data=0000000016,  ALLOC=0000115056,  SYS=0003901688
	[2^05]: Data=0000000032,  ALLOC=0000115072,  SYS=0003901688
	[2^06]: Data=0000000064,  ALLOC=0000115104,  SYS=0003901688
	[2^07]: Data=0000000128,  ALLOC=0000115168,  SYS=0003901688
	[2^08]: Data=0000000256,  ALLOC=0000115296,  SYS=0003901688
	[2^09]: Data=0000000512,  ALLOC=0000115552,  SYS=0003901688
	[2^10]: Data=0000001024,  ALLOC=0000116064,  SYS=0003901688
	[2^11]: Data=0000002048,  ALLOC=0000117088,  SYS=0003901688
	[2^12]: Data=0000004096,  ALLOC=0000119136,  SYS=0003901688
	[2^13]: Data=0000008192,  ALLOC=0000123232,  SYS=0003901688
	[2^14]: Data=0000016384,  ALLOC=0000131424,  SYS=0003901688
	[2^15]: Data=0000032768,  ALLOC=0000147808,  SYS=0003901688
	[2^16]: Data=0000065536,  ALLOC=0000180576,  SYS=0003901688
	[2^17]: Data=0000131072,  ALLOC=0000246112,  SYS=0003901688
	[2^18]: Data=0000262144,  ALLOC=0000377184,  SYS=0003901688
	[2^19]: Data=0000524288,  ALLOC=0000639328,  SYS=0004983032
	[2^20]: Data=0001048576,  ALLOC=0001163616,  SYS=0004983032
--- PASS: TestSuite (0.13s)
PASS
BenchmarkFillData-4      	    5000	    310299 ns/op
BenchmarkCryptoFillData-4	    5000	    704925 ns/op
ok  	github.com/dearing/havoc	5.272s

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

[havoc_server]: https://github.com/dearing/havoc_server
[issues]: https://github.com/dearing/havoc/issues
