package main

import (
	"fmt"
	"os"
	"os/exec"

	archiver "github.com/mholt/archiver/v3"
)

func main() {
	// Check if the correct number of arguments are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <tar_file> <file to execute>")
		os.Exit(1)
	}

	tarFile := os.Args[1]
	//password := os.Args[2]
	destination := "."

	z := archiver.Tar{
		MkdirAll:               true,
		ContinueOnError:        true,
		OverwriteExisting:      true,
		ImplicitTopLevelFolder: false,
	}

	// Use archiver library to open and extract the ZIP file
	//err := archiver.Zip.Open(zipFile, password)
	err := z.Unarchive(tarFile, destination)
	if err != nil {
		fmt.Println("Failed to unzip:", err)
		os.Exit(1)
	}

	fmt.Println("Unzip completed successfully!")

	out, err := exec.Command(os.Args[2]).Output()
	//err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute:", err)
		os.Exit(1)
	}

	fmt.Printf("output %s", out)

}
