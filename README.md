long
====

[![GoDoc](https://godoc.org/github.com/logrusorgru/long?status.svg)](https://godoc.org/github.com/logrusorgru/long)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/logrusorgru/long.svg)](https://travis-ci.org/logrusorgru/long)
[![Coverage Status](https://coveralls.io/repos/logrusorgru/long/badge.svg?branch=master)](https://coveralls.io/r/logrusorgru/long?branch=master)
[![GoReportCard](http://goreportcard.com/badge/logrusorgru/long)](http://goreportcard.com/report/logrusorgru/long)
[![paypal donate](https://img.shields.io/badge/paypal%20%24-donate-orange.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=TSRPRBWXLDWRA)

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
as published by Sam Hocevar. See the LICENSE.md file for more details.
