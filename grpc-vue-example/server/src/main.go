package main

import (
	"flag"
	"[[.ModName]]/src/interface/rpc/server"
)

var (
	addr   = flag.String("addr", ":8080", "addr to bind")
	dbconf = flag.String("dbconf", "dbconfig.yml", "database config")
	env    = flag.String("env", "development", "application envirionment")
)

func init() {
	flag.Parse()
}

func main() {
	server.GrpcRun(*dbconf, *env, *addr)
}
