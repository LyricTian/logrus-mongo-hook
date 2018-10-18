# An asynchronous MongoDB Hook for [Logrus](https://github.com/sirupsen/logrus)

[![Build][Build-Status-Image]][Build-Status-Url] [![Codecov][codecov-image]][codecov-url] [![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

## Quick Start

### Download and install

```bash
$ go get -u -v github.com/LyricTian/logrus-mongo-hook
```

### Usage

```go
package main

import (
    "github.com/LyricTian/logrus-mongo-hook"
    "github.com/sirupsen/logrus"
)

func main() {
    hook := mongohook.DefaultWithURL("127.0.0.1:27017","test","t_log")
    defer hook.Flush()

    log := logrus.New()
    log.AddHook(hook)
}
```

## MIT License

    Copyright (c) 2018 Lyric

[Build-Status-Url]: https://travis-ci.org/LyricTian/logrus-mongo-hook
[Build-Status-Image]: https://travis-ci.org/LyricTian/logrus-mongo-hook.svg?branch=master
[codecov-url]: https://codecov.io/gh/LyricTian/logrus-mongo-hook
[codecov-image]: https://codecov.io/gh/LyricTian/logrus-mongo-hook/branch/master/graph/badge.svg
[reportcard-url]: https://goreportcard.com/report/github.com/LyricTian/logrus-mongo-hook
[reportcard-image]: https://goreportcard.com/badge/github.com/LyricTian/logrus-mongo-hook
[godoc-url]: https://godoc.org/github.com/LyricTian/logrus-mongo-hook
[godoc-image]: https://godoc.org/github.com/LyricTian/logrus-mongo-hook?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg
