# maqe-bot
[![Build Status](https://travis-ci.org/Phaicom/maqe-bot.svg?branch=master)](https://travis-ci.org/Phaicom/maqe-bot)

## Getting Started

Command set for walking in 2-dimensional plane

* R: Turn 90 degree to the right
* L: Turn 90 degree to the left
* WN: Walk straight for N point(s) 

For example
```
W5RW5
```
Which mean walk straight 10 points, turn right and walk straight 10 points

## Prerequisites

* [Go](https://golang.org)
* [Dep](https://github.com/golang/dep) - Go dependency management tool 

## Installing

Run program with following args such as W5RW5RW2RW1R

```
$ go run main.go W5RW5RW2RW1R
```

Testing

```
$ go test ./models/
```

## Author

By [Reawpai Chunsoi](https://github.com/phaicom/)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details