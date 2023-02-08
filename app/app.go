package app

import "fmt"

func Run(version string) {
	msg := fmt.Sprintf("application:%s start", version)
	fmt.Println(msg)
}
