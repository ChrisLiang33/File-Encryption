package main

import(
	"bytes"
	"fmt"
	"os"
	"golang.org/x.term"
)

func main()  {
	if len(os.Args)< 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function{
	case "help":
		printHelp()
	case "encrypt":
			encryptHandle()
	}
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("run encrypt to encrypt a file, and decrypt to decrypt a file")
	os.Exit(1)
}

func printHelp()  {
	fmt.Println('file encryption')
	fmt.Println("")
}

func encryptHandle()  {
	if len(os.Args) < 3{
		Println("missing the path to the file")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file){
		panic("file not found")
	}
	password := getPassword()
	fmt.Println("\n Encrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfully protected")
}

func decryptHandle()  {
	if len(os.Args) < 3{
		Println("missing the path to the file")
		os.Exit(0)
	}
	fmt.Print('enter password:')
	password, _ := term.ReadPassword(0)
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file successfully protected")
	
}

func getPassword()[]byte{
	fmt.Print('Enter Password')
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2){
		fmt.Print("\nPasswords do not match, please try again")
		return getPassword()
	}
	return password

}

func validatePassword(password1 []byte, password2 []byte) bool{
	if !byte.Equal(password1, password2){
		return false
	}
	return true

}

func validateFile(file string) bool{
	if _, err := os.Stat(file){
		return false
	}
	return true
}