package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func configurationQuestion(reader *bufio.Reader, description string, defaultValue string) (value string) {
	colorBlue := "\033[34m"
	colorWhite := "\033[37m"
	fmt.Print(string(colorBlue), description + " (" + defaultValue +") ")
	value, _ = reader.ReadString('\n')
	if value != "\n" {
		fmt.Println(string(colorWhite), description, value)
		return strings.TrimSuffix(value, "\n")
	}
	fmt.Println(string(colorWhite), description, defaultValue)
	return defaultValue;
}

func commandExecutor(out *bytes.Buffer, cmdError *bytes.Buffer, application string, args ...string){
	cmd := exec.Command(application, args...)	
	cmd.Stdout = out
	cmd.Stderr = cmdError
	exec_error := cmd.Run()
	cmd.Wait()
	if exec_error != nil {
		fmt.Println("Error: " + ": " + cmdError.String())
		return
	}
	cmd.Process.Kill()
}

func main() {
	colorRed := "\033[31m"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(string(colorRed), "   Welcome to App Creator Installation   ")
	fmt.Println(string(colorRed), "-----------------------------------------")
	// Configs to be collected from user
	var APP_NAME = configurationQuestion(reader, "Application Name","App-Creator")
	var WEB_APP_PORT = "4200"
	var API_APP_PORT = "8000"
	var WEB_APP_DOMAIN = configurationQuestion(reader, "Web Application Domain name","https://appcreator.com/")
	var DB_PORT = "4454"
	var DB_DATABASE = "app_creator_mysql"
	var DB_PASSWORD = "yNaJuleX41aKNiBRy54VLXMxos30"
	var DB_USERNAME = "admin"
	// Create Environment File
	filename := ".env"
    fileStat, err := os.Stat(filename)
	if fileStat != nil {
        os.Remove(filename)
		os.Remove("database")
    }
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	fmt.Fprintf(file, "APP_NAME=%s\n", APP_NAME)
	fmt.Fprintf(file, "WEB_APP_PORT=%s\n", WEB_APP_PORT)
	fmt.Fprintf(file, "API_APP_PORT=%s\n", API_APP_PORT)
	fmt.Fprintf(file, "WEB_APP_DOMAIN=%s\n", WEB_APP_DOMAIN)
	fmt.Fprintf(file, "DB_PORT=%s\n", DB_PORT)
	fmt.Fprintf(file, "DB_DATABASE=%s\n", DB_DATABASE)
	fmt.Fprintf(file, "DB_PASSWORD=%s\n", DB_PASSWORD)
	fmt.Fprintf(file, "DB_USERNAME=%s\n", DB_USERNAME)
	fmt.Println(string(colorRed), "-----------------------------------------")
	var out bytes.Buffer
	var stderr bytes.Buffer
	fmt.Println(string(colorRed), "Getting Started . . . ")
	commandExecutor(&out, &stderr, "docker-compose", "down")
	fmt.Println(string(colorRed), "Starting Application . . .")
	commandExecutor(&out, &stderr, "docker-compose", "up", "-d")
	fmt.Println(string(colorRed), "Configuring . . .")
	commandExecutor(&out, &stderr, "docker-compose exec app_creator_api php artisan config:cache")
}