package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/pkcs12"
)

func main() {
	pfxFile := flag.String("pfx", "", "PFX file to break it")
	flag.Parse()

	if *pfxFile == "" {
		log.Fatal("Can't find the pfx file. You must run the command as: break_pfx file")
	}

	data, err := ioutil.ReadFile(*pfxFile)
	if err != nil {
		log.Fatal(err)
	}

	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							password := fmt.Sprintf("%d%d%d%d%d%d", a, b, c, d, e, f)
							fmt.Printf("\rTrying to break with the password %s", password)

							_, _, err = pkcs12.Decode(data, password)
							if err == nil {
								fmt.Printf("\n\n=> The password is %s\n", password)
								os.Exit(0)
							}

						}
					}
				}
			}
		}
	}
}
