[![Coverage Status](https://coveralls.io/repos/github/JanukanS/bytefifo/badge.svg?branch=main)](https://coveralls.io/github/JanukanS/bytefifo?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/JanukanS/bytefifo.svg)](https://pkg.go.dev/github.com/JanukanS/bytefifo)
# bytefifo
A Go package implementing a data structure for storing byte data representing iterations of a single resource (i.e. a GET request on a dynamic resource, file undergoing updates etc.). The data structure resembles a FIFO queue, implemented using Go Rings.
