package config

import "fmt"

const (
	Yellow = "\033[33m" // Yellow color for normal information
	Red    = "\033[31m" // Red color for warnings and errors
	Reset  = "\033[0m"  // Reset the color
)

func Info(message string) {
	fmt.Printf("%s[INF] %s%s\n", Yellow, message, Reset)
}
func Warn(message string) {
	fmt.Printf("%s[WRN] %s%s\n", Red, message, Reset)
}

// ShowASCII displays the ASCII art banner
func ShowASCII() {
	asciiArt := `
    _____ _           _           _     _          __  __    _    
    |  ___| | __ _ ___| |__       | |   | |    __ _|  \/  |  / \   
    | |_  | |/ _' / __| '_ \ _____| |   | |   / _' | |\/| | / _ \  
    |  _| | | (_| \__ \ | | |_____| |___| |__| (_| | |  | |/ ___ \ 
    |_|   |_|\__,_|___/_| |_|     |_____|_____\__,_|_|  |_/_/   \_\  v1.0

                    By $ilent $pectre5
	`
	fmt.Println(asciiArt)
	fmt.Println("")
}
