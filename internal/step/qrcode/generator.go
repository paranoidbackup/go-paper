package qrcode

import qrcode "github.com/skip2/go-qrcode"

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
	
	privateKeysQrCodesPng := make([][][]byte, 0)
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
	png, err := qrcode.Encode(data, qrcode.Highest, 512)
	if err != nil {
		return nil, err
	}
	return png, nil
}

func (g *GeneratorImpl) splitMulti(data []string) [][]string {
	return nil
}

func (g *GeneratorImpl) split(data string) []string {
	return nil
}
