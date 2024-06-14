package main

import (
	"log"
	"sync"
)

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET":  set,
	"GET":  get,
	"HSET": hset,
	"HGET": hget,
}
var SETs = map[string]string{}
var SETsMux = sync.RWMutex{}
var HSETs = map[string]map[string]string{}
var HSETsMux = sync.RWMutex{}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].bulk}
}

func set(args []Value) Value {
	log.Println("Incoming set command")
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].bulk
	value := args[1].bulk

	SETsMux.Lock()
	SETs[key] = value
	SETsMux.Unlock()

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	log.Println("Incoming get command")
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	SETsMux.RLock()
	value, ok := SETs[key]
	SETsMux.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}

func hset(args []Value) Value {
	log.Println("Incoming hset command")
	if len(args) != 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	HSETsMux.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMux.Unlock()

	return Value{typ: "string", str: "OK"}
}

func hget(args []Value) Value {
	log.Println("Incoming hget command")
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSETsMux.RLock()
	value, ok := HSETs[hash][key]
	HSETsMux.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}
