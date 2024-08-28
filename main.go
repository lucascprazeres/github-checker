package main

import "github-checker/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
