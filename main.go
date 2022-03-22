package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	win     bool
	rootCmd = &cobra.Command{
		Use:   "reflex [--win|-w] port",
		Short: "reflex is a better netcat for tunneling shell commands",
		Long:  "if you plan to run this on windows use --win or -w to bind the remote shell against cmd.exe",
		Run:   mainCommand,
	}
)

func mainCommand(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("cannot start reflex without know the port to bind")
		os.Exit(1)
	}

	port := args[0]

	if win {
		fmt.Println("to be implemented")
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to bind to port %s", port))
		os.Exit(1)
	}

	log.Println("Listening on 0.0.0.0:20080")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	var cmd *exec.Cmd

	// be aware that running this with --win cause a panic at line 32
	if !win {
		cmd = exec.Command("/bin/sh", "-i")
	} 

	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)
	cmd.Run()
}

func main() {
	rootCmd.PersistentFlags().BoolVarP(&win, "win", "w", false, "run this on windows")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
