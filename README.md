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
 - [2^01]: Data=0000000002,  ALLOC=0000114576,  SYS=0003639544
 - [2^02]: Data=0000000004,  ALLOC=0000114640,  SYS=0003639544
 - [2^03]: Data=0000000008,  ALLOC=0000114640,  SYS=0003639544
 - [2^04]: Data=0000000016,  ALLOC=0000114640,  SYS=0003639544
 - [2^05]: Data=0000000032,  ALLOC=0000114656,  SYS=0003639544
 - [2^06]: Data=0000000064,  ALLOC=0000114688,  SYS=0003639544
 - [2^07]: Data=0000000128,  ALLOC=0000114752,  SYS=0003639544
 - [2^08]: Data=0000000256,  ALLOC=0000114880,  SYS=0003639544
 - [2^09]: Data=0000000512,  ALLOC=0000115136,  SYS=0003639544
 - [2^10]: Data=0000001024,  ALLOC=0000115648,  SYS=0003639544
 - [2^11]: Data=0000002048,  ALLOC=0000116672,  SYS=0003639544
 - [2^12]: Data=0000004096,  ALLOC=0000118720,  SYS=0003639544
 - [2^13]: Data=0000008192,  ALLOC=0000122816,  SYS=0003639544
 - [2^14]: Data=0000016384,  ALLOC=0000131008,  SYS=0003639544
 - [2^15]: Data=0000032768,  ALLOC=0000147392,  SYS=0003639544
 - [2^16]: Data=0000065536,  ALLOC=0000180160,  SYS=0003639544
 - [2^17]: Data=0000131072,  ALLOC=0000245696,  SYS=0003639544
 - [2^18]: Data=0000262144,  ALLOC=0000376768,  SYS=0003639544
 - [2^19]: Data=0000524288,  ALLOC=0000638912,  SYS=0004720888
 - [2^20]: Data=0001048576,  ALLOC=0001163200,  SYS=0004720888
 - [2^21]: Data=0002097152,  ALLOC=0002211776,  SYS=0006883576
 - [2^22]: Data=0004194304,  ALLOC=0004311040,  SYS=0011213048
 - [2^23]: Data=0008388608,  ALLOC=0008505344,  SYS=0019871992
 - [2^24]: Data=0016777216,  ALLOC=0016894016,  SYS=0037189880
 - [2^25]: Data=0033554432,  ALLOC=0033671296,  SYS=0071825656
 - [2^26]: Data=0067108864,  ALLOC=0067225792,  SYS=0141097208
 - [2^27]: Data=0134217728,  ALLOC=0134334784,  SYS=0279902456
 - [2^28]: Data=0268435456,  ALLOC=0268552576,  SYS=0556988664
 - [2^29]: Data=0536870912,  ALLOC=0536988096,  SYS=1111161080
 - [2^30]: Data=1073741824,  ALLOC=1073859008,  SYS=2219505912
--- PASS: TestSuite (2.50s)
PASS
BenchmarkFillData-4      	    5000	    318195 ns/op
BenchmarkCryptoFillData-4	    5000	    707957 ns/op
ok  	github.com/dearing/havoc	8.627s

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
