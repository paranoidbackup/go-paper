package document

type Exporter interface {
	Export(input ExporterInput) error
}
