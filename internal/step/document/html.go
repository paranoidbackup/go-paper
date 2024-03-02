package document

type HtmlExporter struct{}

func NewHtmlExporter() *HtmlExporter {
	return &HtmlExporter{}
}

func (e *HtmlExporter) Export(input ExporterInput, destinationDir string) error {
	return nil
}
