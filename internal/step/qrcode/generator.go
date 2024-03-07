package qrcode

import "github.com/skip2/go-qrcode"

const partialSize = 1000

type Generator interface {
	Generate(input Input) (*Output, error)
}

type GeneratorImpl struct {
}

func NewGenerator() (*GeneratorImpl, error) {
	return &GeneratorImpl{}, nil
}

func (g *GeneratorImpl) Generate(input Input) (*Output, error) {
	privateKeysSplit := g.splitMulti(input.PrivateKeys())
	dataSplit := g.split(input.Data())

	privateKeysQrCodesPng := make([][][]byte, len(privateKeysSplit))
	dataQrCodesPng := make([][]byte, 0)

	for i, pKey := range privateKeysSplit {
		if len(privateKeysQrCodesPng[i]) == 0 {
			privateKeysQrCodesPng[i] = make([][]byte, 0)
		}
		for _, pKeyPart := range pKey {
			png, err := g.encode(pKeyPart)
			if err != nil {
				return nil, err
			}
			privateKeysQrCodesPng[i] = append(privateKeysQrCodesPng[i], png)
		}
	}
	for _, dataPart := range dataSplit {
		png, err := g.encode(dataPart)
		if err != nil {
			return nil, err
		}
		dataQrCodesPng = append(dataQrCodesPng, png)
	}

	return &Output{
		projectId:             input.ProjectID(),
		docId:                 input.DocID(),
		docName:               input.DocName(),
		publicKeys:            input.PublicKeys(),
		passphrases:           input.Passphrases(),
		privateKeysSplit:      privateKeysSplit,
		dataSplit:             dataSplit,
		privateKeysQrCodesPng: privateKeysQrCodesPng,
		dataQrCodesPng:        dataQrCodesPng,
	}, nil
}

func (g *GeneratorImpl) encode(data string) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(data, qrcode.Highest, 1024)
	if err != nil {
		return nil, err
	}
	return png, nil
}

func (g *GeneratorImpl) splitMulti(data []string) [][]string {
	result := make([][]string, 0)
	for _, v := range data {
		result = append(result, g.split(v))
	}
	return result
}

func (g *GeneratorImpl) split(data string) []string {
	runes := []rune(data)
	result := make([]string, 0)

	i := 0
	for {
		start := i
		end := i + partialSize - 1
		if end >= len(runes) {
			end = len(runes) - 1
		}
		for end < len(runes)-1 && (string(runes[end]) == "\n" || string(runes[end]) == " ") {
			end++
		}
		result = append(result, string(runes[start:end]))
		if end == len(runes)-1 {
			break
		}
		i = end + 1
	}

	return result
}
