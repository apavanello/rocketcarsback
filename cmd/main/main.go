package main

import (
	"fmt"

	s "github.com/apavanello/rocketcarsback/internal/server"
	"github.com/apavanello/rocketcarsback/versioninfo"
)

var (
	port = 50051
)

func main() {
	fmt.Println("Project: " + versioninfo.ProjectName)
	fmt.Println("Description: " + versioninfo.ProjectDescription)
	fmt.Println("Copyright: " + versioninfo.ProjectCopyright)
	fmt.Println("Version: " + versioninfo.Version)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)

	s.StartServer(port)
}
