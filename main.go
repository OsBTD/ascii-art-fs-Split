package main

import "ascii/ascii"

func main() {
	_, _, Banner := ascii.Input()
	Map := ascii.Populate()
	ascii.Print(Map, Banner)
}
