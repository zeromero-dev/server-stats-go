package routes

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func ViewSystem(client *ssh.Client) {
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Error creating session: %v\n", err)
		return
	}
	defer session.Close()

	output, err := session.CombinedOutput("neofetch")
	if err != nil {
		fmt.Printf("Command execution error: %v\n", err)
		return
	}
	fmt.Println("System Information:")
	fmt.Println(string(output))
}

func ViewLoad(client *ssh.Client) {
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Error creating session: %v\n", err)
		return
	}
	defer session.Close()

	output, err := session.CombinedOutput("uptime")
	if err != nil {
		fmt.Printf("Command execution error: %v\n", err)
		return
	}
	fmt.Println("Load Resources:")
	fmt.Println(string(output))
}

func ShutOffSystem(client *ssh.Client) {
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Error creating session: %v\n", err)
		return
	}
	defer session.Close()

	err = session.Run("sudo shutdown -h now")
	if err != nil {
		fmt.Printf("Shutdown command error: %v\n", err)
		return
	}
	fmt.Println("Shutdown command executed.")
}
