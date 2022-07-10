package libs

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/sqltocsv"
)

//MysqlConnector is struct for connect in db
type MysqlConnector struct {
	db       *sql.DB
	tables   []string
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

//getTables return tables name for connector
func (m *MysqlConnector) getTables() (err error) {
	rows, err := m.db.QueryContext(context.Background(), "SHOW TABLES")
	if err != nil {
		return
	}
	for rows.Next() {
		var table_name string
		if err := rows.Scan(&table_name); err != nil {
			log.Fatal(err)
		}
		m.tables = append(m.tables, table_name)
	}
	err = nil
	return
}

//getJSON return data of table
func (m *MysqlConnector) getJSON(sql_string string) (string, error) {
	stmt, err := m.db.Prepare(sql_string)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	tableData := make([]map[string]interface{}, 0)

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return "", err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			v := values[i]

			b, ok := v.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = v
			}
		}

		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

//getCSV return data of table in csv format
func (m *MysqlConnector) getCSV(sql_string string) (string, error) {
	stmt, err := m.db.Prepare(sql_string)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return "", err
	}
	defer rows.Close()
	converter := sqltocsv.New(rows)
	return converter.String(), nil
}

//ListTables return all tables in database
func (m *MysqlConnector) ListTables() (err error) {
	if err = m.getTables(); err != nil {
		return err
	}
	for _, table := range m.tables {
		fmt.Printf("%s\n", table)
	}
	return nil
}

//Export save data of tables in files
func (m *MysqlConnector) Export(format string, query string) (err error) {
	if err = m.getTables(); err != nil {
		return err
	}
	switch format {
	case "json":
		if query != "" {

			f, err := os.Create(fmt.Sprintf("./%s.%s", "result", format))
			if err != nil {
				return err
			}

			json_data, err := m.getJSON(query)
			if err != nil {
				return err
			}

			bytes_content := []byte(json_data)
			if _, err = f.Write(bytes_content); err != nil {
				return err
			}
			if err = f.Sync(); err != nil {
				return err
			}
			if err = f.Close(); err != nil {
				return err
			}
			break
		}
		for _, table := range m.tables {
			f, err := os.Create(fmt.Sprintf("./%s.%s", table, format))
			if err != nil {
				return err
			}
			json_data, err := m.getJSON(fmt.Sprintf("SELECT * FROM %s", table))
			if err != nil {
				return err
			}
			bytes_content := []byte(json_data)
			if _, err = f.Write(bytes_content); err != nil {
				return err
			}
			if err = f.Sync(); err != nil {
				return err
			}
			if err = f.Close(); err != nil {
				return err
			}
		}
		break
	case "csv":
		if query != "" {

			f, err := os.Create(fmt.Sprintf("./%s.%s", "result", format))
			if err != nil {
				return err
			}

			csv_data, err := m.getCSV(query)
			if err != nil {
				return err
			}

			bytes_content := []byte(csv_data)
			if _, err = f.Write(bytes_content); err != nil {
				return err
			}
			if err = f.Sync(); err != nil {
				return err
			}
			if err = f.Close(); err != nil {
				return err
			}
			break
		}
		for _, table := range m.tables {
			f, err := os.Create(fmt.Sprintf("./%s.%s", table, format))
			if err != nil {
				return err
			}
			csv_data, err := m.getCSV(fmt.Sprintf("SELECT * FROM %s", table))
			if err != nil {
				return err
			}

			bytes_content := []byte(csv_data)
			if _, err = f.Write(bytes_content); err != nil {
				return err
			}
			if err = f.Sync(); err != nil {
				return err
			}
			if err = f.Close(); err != nil {
				return err
			}
		}
	default:
		fmt.Printf("Error: Format %s not accepted\n", format)
		os.Exit(1)
	}

	return nil
}

//Connect is connect in db
func (m *MysqlConnector) Close() (err error) {
	if err = m.db.Close(); err != nil {
		return
	}
	err = nil
	return
}
