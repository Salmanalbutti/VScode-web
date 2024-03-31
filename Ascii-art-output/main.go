package Modify

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AsciiArtOutputF(str, banner, output string) {
	// Read the contents of the file
	content, err := os.Open("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	defer content.Close()
	// Make sure argument is a printable character (i.e., between 32 and 126)
	for _, j := range str {
		if j < 32 || j > 126 {
			fmt.Println("Please enter valid characters")
			return
		}
	}
	userInput := strings.Split(str, "\\n")
	for userInput[0] == "" && len(userInput) > 1 {
		fmt.Println()
		userInput = userInput[1:]
	}
	var result string
	if !strings.Contains(output, ".txt"){
		output = output + ".txt"
	}
	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outputFile.Close()
	for _, c := range userInput {
		if c != "" {
			for l := 0; l < 8; l++ {
				for _, k := range c {
					startLine := (2 + int(k-32)*9 + l)
					content.Seek(0, 0) // Reset file position to the beginning
					content1 := bufio.NewReader(content)
					for i := 0; i < startLine; i++ {
						str, _ = content1.ReadString('\n')
						str = strings.TrimRight(str, "\r\n") // Trim Windows-style line ending
					}
					result += str
				}
				_, err := outputFile.WriteString(result + "\n")
				if err != nil {
					fmt.Println("Failed to write to output file:", err)
					return
				}
				result = ""
			}
		} else if len(userInput) != 1 {
			fmt.Println()
		}
	}
}
func AsciiArtOutput(str, output string) {
	// Read the contents of the file
	content, err := os.Open("banners/" + "standard.txt")
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	defer content.Close()
	// Make sure argument is a printable character (i.e., between 32 and 126)
	for _, j := range str {
		if j < 32 || j > 126 {
			fmt.Println("Please enter valid characters")
			return
		}
	}
	userInput := strings.Split(str, "\\n")
	for userInput[0] == "" && len(userInput) > 1 {
		fmt.Println()
		userInput = userInput[1:]
	}
	var result string
	if !strings.Contains(output, ".txt"){
		output = output + ".txt"
	}
	
	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outputFile.Close()
	for _, c := range userInput {
		if c != "" {
			for l := 0; l < 8; l++ {
				for _, k := range c {
					startLine := (2 + int(k-32)*9 + l)
					content.Seek(0, 0) // Reset file position to the beginning
					content1 := bufio.NewReader(content)
					for i := 0; i < startLine; i++ {
						str, _ = content1.ReadString('\n')
						str = strings.TrimRight(str, "\r\n") // Trim Windows-style line ending
					}
					result += str
				}
				_, err := outputFile.WriteString(result + "\n")
				if err != nil {
					fmt.Println("Failed to write to output file:", err)
					return
				}
				result = ""
			}
		} else if len(userInput) != 1 {
			fmt.Println()
		}
	}
}
func AsciiArt(fullstr, banner string) {
	var filePath string
	if banner == "standard" || banner == "" {
		filePath = "banners/standard.txt"
	} else if banner == "shadow" {
		filePath = "banners/shadow.txt"
	} else if banner == "thinkertoy" {
		filePath = "banners/thinkertoy.txt"
	} else {
		fmt.Println("Please enter correct banner name")
		return
	}
	
	// Read the contents of the file
	content, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	defer content.Close()
	// Make sure argument is a printable character (i.e., between 32 and 126)
	checkInput := os.Args[1]
	for _, j := range checkInput {
		if j < 32 || j > 126 {
			fmt.Println("Please enter valid characters")
			return
		}
	}
	userInput := strings.Split(checkInput, "\\n")
	for userInput[0] == "" && len(userInput) > 1 {
		fmt.Println()
		userInput = userInput[1:]
	}
	var str string
	var result string
	for _, c := range userInput {
		if c != "" {
			for l := 0; l < 8; l++ {
				for _, k := range c {
					startLine := (2 + int(k-32)*9 + l)
					content.Seek(0, 0) // Reset file position to the beginning
					content1 := bufio.NewReader(content)
					for i := 0; i < startLine; i++ {
						str, _ = content1.ReadString('\n')
						str = strings.TrimRight(str, "\r\n") // Trim Windows-style line ending
					}
					result += str
				}
				fmt.Println(result)
				result = ""
			}
		} else if len(userInput) != 1 {
			fmt.Println()
		}
	}
}