long
====

[![GoDoc](https://godoc.org/github.com/logrusorgru/long?status.svg)](https://godoc.org/github.com/logrusorgru/long)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/long.svg)](https://travis-ci.org/logrusorgru/long)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/long/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/long?branch=master)
[![GoReportCard](https://goreportcard.com/badge/logrusorgru/long)](https://goreportcard.com/report/logrusorgru/long)
[![Gitter](https://img.shields.io/badge/chat-on_gitter-46bc99.svg?logo=data:image%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGhlaWdodD0iMTQiIHdpZHRoPSIxNCI%2BPGcgZmlsbD0iI2ZmZiI%2BPHJlY3QgeD0iMCIgeT0iMyIgd2lkdGg9IjEiIGhlaWdodD0iNSIvPjxyZWN0IHg9IjIiIHk9IjQiIHdpZHRoPSIxIiBoZWlnaHQ9IjciLz48cmVjdCB4PSI0IiB5PSI0IiB3aWR0aD0iMSIgaGVpZ2h0PSI3Ii8%2BPHJlY3QgeD0iNiIgeT0iNCIgd2lkdGg9IjEiIGhlaWdodD0iNCIvPjwvZz48L3N2Zz4%3D&logoWidth=10)](https://gitter.im/logrusorgru/long?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

The _long_ is a method of encoding integers. It's similar to
[varint](https://developers.google.com/protocol-buffers/docs/encoding#varints).
But it optimised for negative numbers too.

# Representation

On the [paly.golang.org](http://play.golang.org/p/-eNR53vrL5)

```
// i  - inverse flag (bit)
// n  - next byte flag (bit)
// d  - data (bit)
// [] - byte
// there is an example number (t):
[dddd dddd][dddd dddd][dddd dddd][dddd dddd]
// if t < ^t the t will be inverted and the inverse bit will be set
// decoded:
// .head
// [dddd ddin]
// .body
// [dddd dddn]
// [dddd dddn]
//  etc
// length of result from 1 to 10 inclusive
```

# Get

```
go get github.com/logrusorgru/long
cd $GOPATH/github.com/logrusorgru/long
go test
```

# Licensing

Copyright © 2015 Konstantin Ivanov <kostyarin.ivanov@gmail.com>  
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.
