package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err);
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press any key, ESC to quit")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		if key == keyboard.KeyEsc {
			break
		}

		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
	}
}