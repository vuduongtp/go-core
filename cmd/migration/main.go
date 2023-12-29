package main

import "github.com/vuduongtp/go-core/internal/functions/migration"

func main() {
	checkErr(migration.Run())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
