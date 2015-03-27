// Package middleware provides some types and functions common among middleware.
package middleware

import "net/http"

type (
	// Generator represents the outer layer of a middleware that
	// parses tokens to configure the middleware instance.
	Generator func(Controller) (Middleware, error)

	// Middleware is the middle layer which represents the traditional
	// idea of middleware: it is passed the next HandlerFunc in the chain
	// and returns the inner layer, which is the actual HandlerFunc.
	Middleware func(http.HandlerFunc) http.HandlerFunc

	// A Control provides structured access to tokens from a configuration file
	// and also to properties of the server being configured. Middleware generators
	// use a Controller to construct their middleware instance.
	Controller interface {
		// Next loads the next token. Returns true if a token
		// was loaded; false otherwise. If false, all tokens
		// have already been consumed.
		Next() bool

		// NextArg loads the next token if it is on the same
		// line. Returns true if a token was loaded; false
		// otherwise. If false, all tokens on the line have
		// been consumed.
		NextArg() bool

		// NextLine loads the next token only if it is NOT on the same
		// line as the current token, and returns true if a token was
		// loaded; false otherwise. If false, there is not another token
		// or it is on the same line.
		NextLine() bool

		// NextBlock can be used as the condition of a for loop
		// to load the next token as long as it opens a block or
		// is already in a block. It returns true if a token was
		// loaded, or false when the block's closing curly brace
		// was loaded and thus the block ended. Nested blocks are
		// not (currently) supported.
		NextBlock() bool

		// Val gets the text of the current token.
		Val() string

		// Args is a convenience function that loads the next arguments
		// (tokens on the same line) into an arbitrary number of strings
		// pointed to in arguments. If there are fewer tokens available
		// than string pointers, the remaining strings will not be changed
		// and false will be returned. If there were enough tokens available
		// to fill the arguments, then true will be returned.
		Args(...*string) bool

		// RemainingArgs loads any more arguments (tokens on the same line)
		// into a slice and returns them. Open curly brace tokens also indicate
		// the end of arguments, and the curly brace is not included in
		// the return value nor is it loaded.
		RemainingArgs() []string

		// ArgErr returns an argument error, meaning that another
		// argument was expected but not found. In other words,
		// a line break, EOF, or open curly brace was encountered instead of
		// an argument.
		ArgErr() error

		// Err generates a custom parse error with a message of msg.
		Err(string) error

		// Startup registers a function to execute when the server starts.
		Startup(func() error)

		// Shutdown registers a function to execute when the server exits.
		Shutdown(func() error)

		// Root returns the file path from which the server is serving.
		Root() string

		// Host returns the hostname the server is bound to.
		Host() string

		// Port returns the port that the server is listening on.
		Port() string

		// Context returns the path scope that the Controller is in.
		// Note: This is not currently used, but may be in the future.
		Context() Path
	}
)
