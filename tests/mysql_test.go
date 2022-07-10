package tests

import (
	"log"
	"testing"

	"github.com/informeai/gedb/libs"
)

//go test -run TestMysqlConnector -v
func TestNewMysqlConnector(t *testing.T) {
	mysql_connector := libs.NewMysqlConnector("gedb", "root", "", "127.0.0.1", "3306")
	if mysql_connector == nil {
		t.Errorf("TestNewMysqlConnector: expect != nil, got: nil\n")
	}
}

//go test -run TestMysqlConnect -v
func TestMysqlConnect(t *testing.T) {
	mysql_connector := libs.NewMysqlConnector("gedb", "root", "", "127.0.0.1", "3306")
	err := mysql_connector.Connect()
	if err != nil {
		t.Errorf("TestMysqlConnect: expect not error, got: %s\n", err.Error())
	}
	log.Printf("Connection of mysql: %v\n", mysql_connector)
	mysql_connector.Close()

}

//go test -run TestMysqlListTables
func TestMysqlListTables(t *testing.T) {
	mysql_connector := libs.NewMysqlConnector("gedb", "root", "", "127.0.0.1", "3306")
	err := mysql_connector.Connect()
	if err != nil {
		t.Errorf("TestMysqlListTables: expect: != nil, got: %s\n", err.Error())
	}
	err = mysql_connector.ListTables()
	if err != nil {
		t.Errorf("TestMysqlListTables: expect: != nil, got: %s\n", err.Error())
	}
	mysql_connector.Close()
}

//go test -run TestMysqlListTables
func TestMysqlExport(t *testing.T) {
	mysql_connector := libs.NewMysqlConnector("gedb", "root", "", "127.0.0.1", "3306")
	err := mysql_connector.Connect()
	if err != nil {
		t.Errorf("TestMysqlExport: expect: != nil, got: %s\n", err.Error())
	}
	for _, value := range []string{"json", "csv"} {
		err = mysql_connector.Export("", value)
		if err != nil {
			t.Errorf("TestMysqlExport: expect: != nil, got: %s\n", err.Error())
		}
	}
	mysql_connector.Close()
}
