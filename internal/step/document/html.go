package document

type HtmlExporter struct {
	workDir string
}

func NewHtmlExporter(workDir string) (*HtmlExporter, error) {
	return &HtmlExporter{
		workDir: workDir,
	}, nil
}

func (e *HtmlExporter) Export(input ExporterInput) error {
	return nil
}
