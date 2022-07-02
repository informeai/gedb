package tests

import (
	"flag"
	"fmt"
	"log"
	"testing"

	"github.com/informeai/gedb/helpers"
	"github.com/informeai/gedb/templates"
)

//go test -run -v TestVerifyFlags
func TestVerifyFlags(t *testing.T) {
	hostFlag := flag.String("host", "127.0.0.1", "host of connection")
	databaseFlag := flag.String("db", "gedb", "database of connection")
	usernameFlag := flag.String("user", "admin", "username of connection")
	passwordFlag := flag.String("pass", "secret", "password of connection")
	formatFlag := flag.String("format", "json", "format to export")
	portFlag := flag.String("P", "3306", "port to connection")
	flags := map[string]string{
		"-host":   *hostFlag,
		"-db":     *databaseFlag,
		"-user":   *usernameFlag,
		"-pass":   *passwordFlag,
		"-format": *formatFlag,
		"-P":      *portFlag,
	}
	fmt.Printf("flags: %+v\n", flag.Args())
	flag.Parse()
	err := helpers.VerifyFlags(flags)
	if err != nil {
		t.Errorf("TestVerifyFlags: expect: nil, got: %s\n", err.Error())
		log.Printf("%s\n", templates.Header)
		log.Printf("%s\n", templates.Usage)
	}
}
