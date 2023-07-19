package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// HandleGracefulShutdown handles graceful shutdown by capturing OS signals and terminating the program.
func HandleGracefulShutdown() {
	// Create a buffered channel to receive OS signals (SIGINT and SIGTERM).
	quit := make(chan os.Signal, 1)

	// Ensure the channel is closed when the function returns.
	defer close(quit)

	// Notify the channel when SIGINT or SIGTERM signals are received.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait until a signal is received on the 'quit' channel.
	// The program will block at this line until one of the specified signals is caught.
	<-quit
}
