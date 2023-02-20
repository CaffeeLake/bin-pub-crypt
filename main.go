package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/go-compile/rome"
	"github.com/go-compile/rome/brainpool"
)

func main() {
	for i, v := range os.Args {
		if v == "keygen" || v == "-k" {
			keygen(os.Args[i+1])
		} else if v == "encrypt" || v == "-e" {
			encrypt(os.Args[i+1], os.Args[i+2], os.Args[i+3])
		} else if v == "decrypt" || v == "-d" {
			decrypt(os.Args[i+1], os.Args[i+2], os.Args[i+3])
		} else if len(os.Args) == 1 || v == "help" || v == "--help" || v == "-h" {
			help()
		}
	}
	os.Exit(0)
}

func help() {
	fmt.Println("Usage: " + os.Args[0] + " [command]")
	fmt.Println("Commands:")
	fmt.Println("	keygen [keyfilename]")
	fmt.Println("	encrypt [publickeyfile] [inputfile] [outputfile]")
	fmt.Println("	decrypt [privatekeyfile] [inputfile] [outputfile]")
}

func keygen(file string) error {
	k, err := brainpool.GenerateP512r1()
	if err != nil {
		panic(err)
	}
	public, err := k.Public().Key()
	if err != nil {
		panic(err)
	}
	private, err := k.Private()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(file+".pub", public, 0777)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(file+".key", private, 0777)
	if err != nil {
		panic(err)
	}
	return err
}

func encrypt(publickey string, decfile string, encfile string) error {
	public, err := loadpublickey(publickey)
	if err != nil {
		panic(err)
	}

	f, err := os.ReadFile(decfile)
	if err != nil {
		panic(err)
	}

	x, err := public.Encrypt(f, rome.CipherChaCha20Poly1305, sha256.New())
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(encfile, x, 0777)
	if err != nil {
		panic(err)
	}
	return err
}

func decrypt(privatekey string, encfile string, decfile string) error {
	private, err := loadprivatekey(privatekey)
	if err != nil {
		panic(err)
	}

	f, err := os.ReadFile(encfile)
	if err != nil {
		panic(err)
	}

	x, err := private.Decrypt(f, rome.CipherChaCha20Poly1305, sha256.New())
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(decfile, x, 0777)
	if err != nil {
		panic(err)
	}
	return err
}

func loadpublickey(file string) (*rome.ECPublicKey, error) {
	publickey, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	public, err := rome.ParseECPublic(publickey)
	if err != nil {
		panic(err)
	}
	return public, err
}

func loadprivatekey(file string) (*rome.ECKey, error) {
	privatekey, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	private, err := rome.ParseECPrivate(privatekey)
	if err != nil {
		panic(err)
	}
	return private, err
}
