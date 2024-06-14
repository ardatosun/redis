package main

import (
	"bufio"
	"io"
)

/*
RESP cheatsheet for reference

RESP data type 	Protocol 	Category 	First byte
Simple strings 	RESP2 		Simple 		+
Simple Errors 	RESP2 		Simple 		-
Integers 		RESP2 		Simple 		:
Bulk strings 	RESP2 		Aggregate 	$
Arrays 			RESP2 		Aggregate 	*

Nulls 			RESP3 		Simple 		_
Booleans 		RESP3 		Simple 		#
Doubles 		RESP3 		Simple 		,
Big numbers 	RESP3 		Simple 		(
Bulk errors 	RESP3 		Aggregate 	!
Verbatim stringsRESP3 		Aggregate 	=
Maps 			RESP3 		Aggregate 	%
Sets 			RESP3 		Aggregate 	~
Pushes 			RESP3 		Aggregate 	>
*/
const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

type Writer struct {
	writer io.Writer
}
