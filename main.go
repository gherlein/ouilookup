package main

import (
	"fmt"
	"github.com/klauspost/oui"
	"os"
)

var (
	mac    string = ""
	oui_db string = ""
)

func init() {

	if len(os.Args) > 1 {
		mac = os.Args[1]
	}
	if mac == "" {
		fmt.Printf("usage:  %s <mac address>\n", os.Args[0])
		os.Exit(3)
	}
}

func main() {
	db, err := oui.OpenStaticFile(oui_db)
	if err != nil {
		fmt.Println("Error reading the OUI database - should be at %s", oui_db)
		fmt.Println("You can download it from http://standards-oui.ieee.org/oui.txt")
		os.Exit(3)
	}
	entry, err := db.Query(mac)
	fmt.Printf("%-16s   %-50s\n", mac, entry.Manufacturer)
}
