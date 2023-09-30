package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Define Member struct
type Member struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Job     string `json:"job"`
	Reason  string `json:"reason"`
}

func main() {
	// Simple argument check
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <ID or Name>")
		return
	}

	// Get argument
	arg := os.Args[1]

	// Read the JSON file.
	jsonData, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading data.json:", err)
		return
	}

	// Unmarshal the JSON data into a slice of Member structs.
	var anggota []Member
	if err := json.Unmarshal(jsonData, &anggota); err != nil {
		fmt.Println("Error Unmarshal JSON:", err)
		return
	}

	// Mapping data by ID for ID lookUp
	idMap := make(map[int]Member)
	for _, member := range anggota {
		idMap[member.ID] = member
	}

	// Mapping data by Name for Name lookUp
	nameMap := make(map[string]Member)
	for _, member := range anggota {
		nameMap[strings.ToLower(member.Name)] = member
	}

	// LookUp process based on argument type
	if id, err := strconv.Atoi(arg); err == nil {
		// If argument is integer then lookUp by ID
		member, found := idMap[id]
		if found {
			printMember(member)
		} else {
			fmt.Println("No member found with ID:", id)
		}
	} else {
		// If argument is string then lookUp by Name
		member, found := nameMap[strings.ToLower(arg)]
		if found {
			printMember(member)
		} else {
			fmt.Println("No member found with Name:", arg)
		}
	}
}

// Function to print Member Detail
func printMember(member Member) {
	rep := len(member.Reason) + 8
	fmt.Println(strings.Repeat("=", rep))
	fmt.Printf("Id: %d\n", member.ID)
	fmt.Printf("Nama: %s\n", member.Name)
	fmt.Printf("Alamat: %s\n", member.Address)
	fmt.Printf("Pekerjaan: %s\n", member.Job)
	fmt.Printf("Alasan: %s\n", member.Reason)
	fmt.Println(strings.Repeat("=", rep))
}
