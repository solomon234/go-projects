package main

import "myapp/packageone"

var myVar = "my var in main"

func main() {
	// variables
	blockVar := "a var"

	packageone.PrintMe(blockVar, myVar, packageone.PackageVar)
}
