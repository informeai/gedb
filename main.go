package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/informeai/gedb/services"
	"github.com/informeai/gedb/templates"
)

var (
	host   string
	db     string
	user   string
	pass   string
	format string
	port   string
	query  string
)

func main() {
	args := os.Args
	if len(os.Args) < 4 {
		fmt.Print(templates.Header)
		fmt.Print(templates.Usage)
		os.Exit(0)
	}
	setFlags := flag.NewFlagSet("", flag.ExitOnError)
	setFlags.StringVar(&host, "host", "127.0.0.1", "host of connection")
	setFlags.StringVar(&db, "db", "", "database name")
	setFlags.StringVar(&user, "u", "root", "username of connection")
	setFlags.StringVar(&pass, "p", "", "password of connection")
	setFlags.StringVar(&format, "f", "json", "format to export")
	setFlags.StringVar(&port, "P", "", "port to connection")
	setFlags.StringVar(&query, "q", "", "query for search")
	setFlags.Parse(os.Args[3:])

	flags := map[string]string{
		"-host": host,
		"-db":   db,
		"-u":    user,
		"-p":    pass,
		"-f":    format,
		"-P":    port,
		"-q":    query,
	}
	if os.Args[1] != "export" {
		log.Println("command not permited")
		fmt.Print(templates.Header)
		fmt.Print(templates.Usage)
		os.Exit(0)
	}
	fmt.Println("Initiate Export...")
	exportService := services.ExportService{
		Args:    args,
		Flags:   flags,
		Driver:  os.Args[2],
		IsQuery: false,
	}
	if err := exportService.Export(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Export Success")
}
