package services

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
	return nil
}
