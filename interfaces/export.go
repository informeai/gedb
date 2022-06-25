package interfaces

//ExportInterface is implementation the export formats
type ExportInterface interface {
	Export(out string, payload string) error
}
