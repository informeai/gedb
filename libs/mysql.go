package libs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//MysqlConnector is struct for connect in db
type MysqlConnector struct {
	db       *sql.DB
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

//NewMysqlConnector return instance of MysqlConnector
func NewMysqlConnector(database, username, password, host, port string) *MysqlConnector {
	if len(port) == 0 {
		port = "3306"
	}
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	return &MysqlConnector{Database: database, Username: username, Password: password, Host: host, Port: port}
}

//Connect is connect in db
func (m *MysqlConnector) Connect() (err error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.Username, m.Password, m.Host, m.Port, m.Database))
	if err != nil {
		return
	}
	m.db = db
	err = nil
	return
}

//Connect is connect in db
func (m *MysqlConnector) Close() (err error) {
	if err = m.db.Close(); err != nil {
		return
	}
	err = nil
	return
}
