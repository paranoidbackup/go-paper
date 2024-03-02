package document

type Exporter interface {
	Export(input ExporterInput, destinationDir string) error
}
