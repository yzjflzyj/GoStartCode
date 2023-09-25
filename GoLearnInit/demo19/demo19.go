package main

import "fmt"

func main() {
	fmt.Println("Email file created successfully.")
	m := map[string]string{"脑子进": "煎鱼了", "煎鱼": "进脑子了"}
	fmt.Printf("m1: %v, len: %d\n", m, len(m))
	clear(m)
	fmt.Printf("m2: %v, len: %d\n", m, len(m))
}
