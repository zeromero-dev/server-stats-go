package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"

	"github.com/joho/godotenv"

	routes "github.com/zeromero-dev/server-stats-go/internal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; will prompt for SSH credentials")
	}

	username := os.Getenv("SSH_USERNAME")
	ip := os.Getenv("SSH_IP")
	if username == "" || ip == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter SSH credentials (username@ip): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}
		input = strings.TrimSpace(input)
		parts := strings.Split(input, "@")
		if len(parts) != 2 {
			log.Fatal("Invalid format. Expected format: username@ip")
		}
		username = parts[0]
		ip = parts[1]
	}

	fmt.Printf("Connecting to %s@%s...\n", username, ip)
	fmt.Print("Enter password: ")
	password, err := readPassword()
	if err != nil {
		log.Fatalf("Error reading password: %v", err)
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		log.Fatalf("Failed to dial SSH: %v", err)
	}
	defer client.Close()

	for {
		fmt.Println("\nMenu Options:")
		fmt.Println("1. View system")
		fmt.Println("2. View Load resources")
		fmt.Println("3. Shut off system")
		fmt.Print("Enter option number: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			routes.ViewSystem(client)
		case 2:
			routes.ViewLoad(client)
		case 3:
			routes.ShutOffSystem(client)
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func readPassword() (string, error) {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println() //next line after password input.
	return strings.TrimSpace(string(bytePassword)), nil
}
