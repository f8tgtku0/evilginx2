package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kgretzky/evilginx2/core"
	"github.com/kgretzky/evilginx2/log"
)

const (
	version = "2.4.0"
)

func main() {
	// Parse command-line flags
	phishlets_dir := flag.String("p", "", "Path to phishlets directory")
	redirect_dir := flag.String("t", "", "Path to redirect pages directory")
	config_dir := flag.String("c", "", "Path to configuration directory")
	debug_log := flag.Bool("debug", false, "Enable debug logging")
	developer_mode := flag.Bool("developer", false, "Enable developer mode (bypass certificate errors)")
	flag.Parse()

	exe_path, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to get executable path: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	if *debug_log {
		log.SetLevel(log.DEBUG)
	}

	log.Info("evilginx2 v%s - by @kgretzky and @an0nud4y", version)
	log.Info("https://github.com/an0nud4y/evilginx2")

	// Initialize the core application
	app, err := core.NewEvilginx(exe_path, *config_dir, *phishlets_dir, *redirect_dir, *developer_mode, *debug_log)
	if err != nil {
		log.Fatal("%v", err)
		os.Exit(1)
	}

	// Start the application
	app.Start()

	// Handle OS signals for graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sig:
		log.Info("shutting down...")
		app.Stop()
	}
}
