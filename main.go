package main

import (
	con "webapiingo/connection"
	rt "webapiingo/routing"
)

func main() {
	con.DataMigration()
	rt.HandlerRouting()
}
