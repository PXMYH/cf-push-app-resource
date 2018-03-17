package main

import {

}

func parrot(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
}

func main {
	// check input 
	if len(os.Args) < 4 {
    parrot("Missing required working dir arg!")
		os.Exit(1)
	}

	cf_uri := os.Args[0]
	cf_username := os.Args[2]

}