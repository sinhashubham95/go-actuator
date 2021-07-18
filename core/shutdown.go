package core

import "os"

// Shutdown is used to shutdown the application
func Shutdown() {
	// passing code 0 here to gracefully shutdown the application
	os.Exit(0)
}
