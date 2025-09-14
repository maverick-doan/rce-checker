package main

import (
	"fmt"
	"os"
	fileOps "win-rce-checker/pkg/file-ops"
	platform "win-rce-checker/pkg/platform"
	procOps "win-rce-checker/pkg/proc-ops"
)

func main() {
	// Ignore any commandline arguments
	_ = os.Args

	platformConfig, err := platform.GetPlatformConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Evidence file
	file, err := fileOps.WriteIndicatorFile(platformConfig.FilePath, platformConfig.Content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("Evidence file written.")

	// Evidence command execution
	err = procOps.ExecuteIndicatorCmd(platformConfig.Cmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("Command execution for \"%s\" failed:", platformConfig.Cmd), err)
	} else {
		fmt.Printf("Command \"%s\" executed.", platformConfig.Cmd)
	}
}
