package utils

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
)

func PtrToSlice[T comparable](s *[]*T) []T {
	size := len(*s)
	actualSlice := make([]T, size)
	for i := 0; i < size; i++ {
		actualSlice[i] = *(*s)[i]
	}

	return actualSlice
}

func Error(message string) {
	red := color.New(color.FgRed).SprintfFunc()

	fmt.Println(red("[ОШИБКА] " + message))
}

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
		break
	case "linux":
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		break
	}
}
