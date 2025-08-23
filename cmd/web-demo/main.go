package main

// https://github.com/gin-gonic/gin			// web api
// https://github.com/uber-go/zap			// logging
// https://github.com/mattn/go-sqlite3		// Database connection
// https://github.com/golang-migrate/migrate // migrations
// https://vektra.github.io/mockery/latest/	// Database mocking
// https://go.dev/doc/tutorial/add-a-test	// testing
// https://medium.com/@rnp0728/secure-password-hashing-in-go-a-comprehensive-guide-5500e19e7c1f // pwd hash

// crud
// serve basic webpage
// unit tests with mock
// logging
// package organization

// solve 4 of each
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/

import (
	"web/example/internal/app"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	app.Launch()
}
