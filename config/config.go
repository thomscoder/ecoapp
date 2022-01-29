package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

var (
	file *os.File
	err  error
)

// Create the certificates
func createCA() error {
	cmd := exec.Command("sh", "-c", "cd certificates && mkcert localhost")
	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

// Get the port number
func getPortNumber() string {
	var port string
	fmt.Print("Port number: ")
	fmt.Scanf("%v", &port)
	return port
}

// Write the server .env
func CreateServerEnvFile() {
	go createCA()
	pNumber := getPortNumber()
	file, err = os.Create(".env")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("HOST=localhost\n")
	file.WriteString("PORT=" + pNumber + "\n")
}

// Read the generated env file
func GetEnvVariables() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("An error occurred loading the environment variables")
	}
	host := os.Getenv("HOST")
	port := ":" + os.Getenv("PORT")
	return host, port
}
