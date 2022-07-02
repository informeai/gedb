package services

import (
	"errors"

	"github.com/informeai/gedb/helpers"
	"github.com/informeai/gedb/libs"
)

//ExportService is struct for export data from driver.
type ExportService struct {
	Args    []string
	Driver  string
	Flags   map[string]string
	IsQuery bool
}

//NewExportService return instance of ExportService
func NewExportService() *ExportService {
	return &ExportService{}
}

//Export execute creation of files of data the database.
func (e *ExportService) Export() error {
	if err := helpers.VerifyFlags(e.Flags); err != nil {
		return err
	}
	switch e.Driver {
	case "mysql":
		connector := libs.NewMysqlConnector(e.Flags["-db"], e.Flags["-user"], e.Flags["-pass"], e.Flags["-host"], "3306")
		if err := connector.Connect(); err != nil {
			return err
		}
		if err := connector.Export(e.Flags["-format"]); err != nil {
			return err
		}
		if err := connector.Close(); err != nil {
			return err
		}
		break
	default:
		return errors.New("driver not permited")
	}
	return nil
}
