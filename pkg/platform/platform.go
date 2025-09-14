package platform

import (
	"fmt"
	"runtime"
)

type Platform string

type PlatformConfig struct {
	OS       Platform
	FilePath string
	Content  string
	Cmd      string
}

const (
	Windows Platform = "windows"
	Linux   Platform = "linux"
	Unknown Platform = "unknown"
)

func GetPlatform() Platform {
	switch runtime.GOOS {
	case "windows":
		return Windows
	case "linux":
		return Linux
	default:
		return Unknown
	}
}

func GetPlatformConfig() (PlatformConfig, error) {
	platform := GetPlatform()
	switch platform {
	case Windows:
		return PlatformConfig{
			OS:       Windows,
			FilePath: "C:\\Users\\Public\\exploit_success.txt",
			Content:  "Exploited",
			Cmd:      "calc.exe",
		}, nil
	case Linux:
		return PlatformConfig{
			OS:       Linux,
			FilePath: "/tmp/exploit_success.txt",
			Content:  "Exploited",
			Cmd:      "echo \"Exploited\" > /tmp/exploit_success.txt",
		}, nil
	default:
		return PlatformConfig{},
			fmt.Errorf("unknown platform: %s", platform)
	}
}
