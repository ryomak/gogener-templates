package main

import (
	"[[.ModName]]/infrastructure/env"
	"[[.ModName]]/presentation/router"
)

func main() {
	r := router.Init()
	if !env.IsProduction() {
		r.NotProd()
	}
	r.Routes()
	r.Run()
}
