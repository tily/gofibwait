# gofibwait

[![Build Status](https://travis-ci.org/tily/gofibwait.svg?branch=master)](https://travis-ci.org/tily/gofibwait)
[![Code Climate](https://codeclimate.com/github/tily/gofibwait/badges/gpa.svg)](https://codeclimate.com/github/tily/gofibwait)
[![Issue Count](https://codeclimate.com/github/tily/gofibwait/badges/issue_count.svg)](https://codeclimate.com/github/tily/gofibwait)
[![Coverage Status](https://coveralls.io/repos/github/tily/gofibwait/badge.svg?branch=master)](https://coveralls.io/github/tily/gofibwait?branch=master)
[![GoDoc](https://godoc.org/github.com/tily/gofibwait?status.svg)](http://godoc.org/github.com/tily/gofibwait)

## Usage

```
w := gofibwait.NewWaiter(60) // maximum is 60 sec
w.Wait() // waits 0 sec
w.Wait() // waits 1 sec
w.Wait() // waits 1 sec
w.Wait() // waits 2 sec
w.Wait() // waits 3 sec
w.Wait() // waits 5 sec
w.Wait() // waits 8 sec
w.Wait() // waits 13 sec
w.Wait() // waits 21 sec
w.Wait() // waits 34 sec
w.Wait() // waits 55 sec
w.Wait() // waits 89 sec
w.Wait() // waits 89 sec
w.Wait() // waits 89 sec
w.Reset()
w.Wait() // waits 0 sec
```
