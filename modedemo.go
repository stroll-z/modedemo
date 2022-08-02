package modedemo

import "fmt"

func Version() string {
	return "v1.0.0"
}

func MakeVersion(m, j int) string {
	return fmt.Sprintf("v%d.%d.0", m, j)
}
