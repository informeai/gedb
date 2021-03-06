package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/informeai/gedb/services"
)

//go test -v -run TestNewExportService
func TestNewExportService(t *testing.T) {
	exportService := services.NewExportService()
	if exportService == nil {
		t.Errorf("TestNewExportService: expect: != nil, got: nil\n")
	}

}

//go test -v -run TestExport
func TestExport(t *testing.T) {
	args := os.Args
	driver := "mysql"
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
	exportService := services.ExportService{
		Args:    args,
		Driver:  driver,
		Flags:   flags,
		IsQuery: false,
	}
	if err := exportService.Export(); err != nil {
		t.Errorf("TestExport: expect: nil, got: %s\n", err.Error())
	}

}
