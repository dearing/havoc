HAVOC
============

Bad library for making a bad application.

```
go test -v -bench=.
=== RUN   TestSetMemory
--- PASS: TestSetMemory (0.00s)
=== RUN   TestResetMemory
--- PASS: TestResetMemory (0.00s)
=== RUN   TestSuite
    [ len(Data) ], [ mem alloc ], [ mem sys ]
[1^ 1]: Data=         2, ALLOC=    113792, SYS=   3639544
[1^ 2]: Data=         4, ALLOC=    113856, SYS=   3639544
[1^ 3]: Data=         8, ALLOC=    113856, SYS=   3639544
[1^ 4]: Data=        16, ALLOC=    113856, SYS=   3639544
[1^ 5]: Data=        32, ALLOC=    113872, SYS=   3639544
[1^ 6]: Data=        64, ALLOC=    113904, SYS=   3639544
[1^ 7]: Data=       128, ALLOC=    113968, SYS=   3639544
[1^ 8]: Data=       256, ALLOC=    114096, SYS=   3639544
[1^ 9]: Data=       512, ALLOC=    114352, SYS=   3639544
[1^10]: Data=      1024, ALLOC=    114864, SYS=   3639544
[1^11]: Data=      2048, ALLOC=    115888, SYS=   3639544
[1^12]: Data=      4096, ALLOC=    117936, SYS=   3639544
[1^13]: Data=      8192, ALLOC=    122032, SYS=   3639544
[1^14]: Data=     16384, ALLOC=    130224, SYS=   3639544
[1^15]: Data=     32768, ALLOC=    146608, SYS=   3639544
[1^16]: Data=     65536, ALLOC=    179376, SYS=   3639544
[1^17]: Data=    131072, ALLOC=    244912, SYS=   3639544
[1^18]: Data=    262144, ALLOC=    375984, SYS=   3639544
[1^19]: Data=    524288, ALLOC=    638128, SYS=   4720888
[1^20]: Data=   1048576, ALLOC=   1162416, SYS=   4720888
[1^21]: Data=   2097152, ALLOC=   2210992, SYS=   6883576
[1^22]: Data=   4194304, ALLOC=   4312048, SYS=  11213048
[1^23]: Data=   8388608, ALLOC=   8506352, SYS=  19871992
[1^24]: Data=  16777216, ALLOC=  16895024, SYS=  37189880
[1^25]: Data=  33554432, ALLOC=  33676520, SYS=  71825656
--- PASS: TestSuite (5.41s)
PASS
BenchmarkMem-4	


   10000	    586819 ns/op
ok  	github.com/dearing/havoc	11.282s

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
