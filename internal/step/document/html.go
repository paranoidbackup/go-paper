package document

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"time"
)

//go:embed tpl/page.gohtml
var htmlPageTemplate string

type HtmlExporter struct {
	workDir string
}

func NewHtmlExporter(workDir string) (*HtmlExporter, error) {
	return &HtmlExporter{
		workDir: workDir,
	}, nil
}

func (e *HtmlExporter) Export(input ExporterInput) error {
	docDir := fmt.Sprintf(
		"%v/docs/%v-%v",
		e.workDir,
		time.Now().Format("2006-01-02"),
		input.DocID(),
	)
	err := os.MkdirAll(docDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%v/images", docDir), os.ModePerm)
	if err != nil {
		return err
	}

	privateKeysQrCodesPaths := make([][]string, 0)
	dataQrCodesPaths := make([]string, 0)

	for privKeyIndex, pngs := range input.PrivateKeysQrCodesPNG() {
		currentKeyPaths := make([]string, 0)

		for pngIndex, pngData := range pngs {
			pngPath := fmt.Sprintf(
				"images/private_key_%v_%v.png",
				privKeyIndex,
				pngIndex,
			)
			err = os.WriteFile(fmt.Sprintf("%v/%v", docDir, pngPath), pngData, os.ModePerm)
			if err != nil {
				return err
			}
			currentKeyPaths = append(currentKeyPaths, pngPath)
		}

		privateKeysQrCodesPaths = append(privateKeysQrCodesPaths, currentKeyPaths)
	}

	for pngIndex, pngData := range input.DataQrCodesPNG() {
		pngPath := fmt.Sprintf("images/data_%v.png", pngIndex)
		err = os.WriteFile(fmt.Sprintf("%v/%v", docDir, pngPath), pngData, os.ModePerm)
		if err != nil {
			return err
		}
		dataQrCodesPaths = append(dataQrCodesPaths, pngPath)
	}

	pageData := PageData{
		ProjectID:               input.ProjectID(),
		DocID:                   input.DocID(),
		DocName:                 input.DocName(),
		Passphrases:             input.Passphrases(),
		PrivateKeysQrCodesPaths: privateKeysQrCodesPaths,
		DataQrCodesPaths:        dataQrCodesPaths,
	}

	tpl, err := template.New("page").Parse(htmlPageTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, pageData)
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("%v/index.html", docDir), buf.Bytes(), os.ModePerm)
}
