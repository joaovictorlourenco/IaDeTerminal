package utils

import (
	"os"
	"strings"
	"time"
)

// FormatTimestamp retorna a hora atual formatada
func FormatTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// SaveToFile salva o conteúdo em um arquivo
func SaveToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// ExtractCodeBlocks extrai blocos de código da resposta
func ExtractCodeBlocks(response string) []string {
	var codeBlocks []string
	lines := strings.Split(response, "\n")
	
	inCodeBlock := false
	currentBlock := ""
	
	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				// Fim do bloco de código
				codeBlocks = append(codeBlocks, currentBlock)
				currentBlock = ""
			}
			inCodeBlock = !inCodeBlock
			continue
		}
		
		if inCodeBlock {
			currentBlock += line + "\n"
		}
	}
	
	return codeBlocks
}

// FormatPrompt formata um prompt com contexto adicional
func FormatPrompt(userInput string, additionalContext string) string {
	if additionalContext == "" {
		return userInput
	}
	
	return additionalContext + "\n\n" + userInput
}