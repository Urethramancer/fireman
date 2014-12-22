package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"

	"code.google.com/p/getopt"
)

// For command-line tools
func p(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

// For errors during setup and other critical conditions
func f(f string, v ...interface{}) {
	log.Fatalf(f, v...)
}

func initConfig() {
	getopt.Parse()
	loadConfig(cfgfile)
}

func genPassword(size int) string {
	valid := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$()-_.:;[]{}")
	pw := make([]byte, size)
	for i := 0; i < size; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(valid))))
		c := valid[n.Int64()]
		pw[i] = c
	}
	return string(pw)
}

func exists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil
}
