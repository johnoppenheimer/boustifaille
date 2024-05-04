package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world from WEB")

	component := home("Maxime")
	component.Render(context.Background(), os.Stdout)
}
