package crypt

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

//go:embed assets/eff_large_wordlist.txt
var effLargeWordlist string

type PassphraseGenerator interface {
	GeneratePassphrase(length int) (string, error)
}

type Random interface {
	Intn(n int) int
}

type randomImpl struct{}

type passphraseGeneratorImpl struct {
	words  map[string]string
	random Random
}

func newPassphraseGenerator() (*passphraseGeneratorImpl, error) {
	words, err := initializeWordlist(effLargeWordlist)
	if err != nil {
		return nil, err
	}
	return &passphraseGeneratorImpl{
		words:  words,
		random: &randomImpl{},
	}, nil
}

func (g *passphraseGeneratorImpl) GeneratePassphrase(length int) (string, error) {
	words := make([]string, length)
	for i := 0; i < len(words); i++ {
		words[i] = g.generateWord()
	}
	return strings.Join(words, " "), nil
}

func initializeWordlist(source string) (map[string]string, error) {
	words := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(source))

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\r\n")
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, "\t")
		if len(parts) != 2 {
			return nil, fmt.Errorf("wrong line format, must be 2 parts: %v", line)
		}

		id := parts[0]
		if !regexp.MustCompile("^[1-6]{5}$").MatchString(id) {
			return nil, fmt.Errorf("wrong word id format: %v", line)
		}

		words[id] = parts[1]
	}

	return words, nil
}

func (g *passphraseGeneratorImpl) generateWord() string {
	for {
		id := g.generateId()
		if word, exists := g.words[id]; exists {
			return word
		}
	}
}

func (g *passphraseGeneratorImpl) generateId() string {
	id := make([]string, 5)
	for i := 0; i < len(id); i++ {
		id[i] = fmt.Sprintf("%v", g.throwDice())
	}
	return strings.Join(id, "")
}

func (g *passphraseGeneratorImpl) throwDice() int {
	return g.random.Intn(6) + 1
}

func (r *randomImpl) Intn(n int) int {
	return rand.Intn(n)
}
