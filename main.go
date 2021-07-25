package main

import (
	lediscfg "github.com/ledisdb/ledisdb/config"
	"github.com/ledisdb/ledisdb/ledis"
)

func main() {
	key := []byte("greet")
	value := []byte("hello-world")
	// Use Ledis's default config
	cfg := lediscfg.NewConfigDefault()
	l, _ := ledis.Open(cfg)
	db, _ := l.Select(0)

	db.Set(key, value)

	db.Get(key)
}
