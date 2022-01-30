package main

import "fmt"

func main() {
	fmt.Println("Hello Compiler")

	var owner string = "Mr Boogiewoo"
	fmt.Println(owner + ", owner of a lonely heart")
	fmt.Println(owner)

	what, what_stance, for_how_long := "heart", "lonely", 24
	fmt.Println(owner + ", owner of a " + what_stance + " " + what)
	fmt.Println(for_how_long)
	
	result := fmt.Sprintf("%s, owner of a %s %s for %d years.", owner, what_stance, what, for_how_long)
	fmt.Println(result)

	fmt.Printf("%s, owner of a %s %s for %d years.", owner, what_stance, what, for_how_long)

}

// format string types
// https://zetcode.com/golang/string-format/
