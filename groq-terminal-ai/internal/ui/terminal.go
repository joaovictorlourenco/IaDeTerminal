package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

// PrintLogo imprime o logo da aplicação
func PrintLogo() {
	color.Output = colorable.NewColorableStdout()
	logo := `
  ____               _______              _           _ 
 / ___|_ __ ___   __|_   _|__ _ __ _ __ (_)_ __   __ _| |
| |  _| '__/ _ \ / _ \| |/ _ \ '__| '_ \| | '_ \ / _' | |
| |_| | | | (_) | (_) | |  __/ |  | | | | | | | | (_| | |
 \____|_|  \___/ \___/|_|\___|_|  |_| |_|_|_| |_|\__,_|_|
                                                          
`
	color.Cyan(logo)
	color.Yellow("Terminal IA - Digite 'ajuda' para ver os comandos disponíveis\n")
}

// PrintResponse imprime a resposta formatada com efeito de digitação
func PrintResponse(response string) {
	fmt.Println()
	color.New(color.FgCyan).Println("Resposta:")
	
	// Verificar se há blocos de código e formatá-los
	lines := strings.Split(response, "\n")
	inCodeBlock := false
	
	for _, line := range lines {
		// Verifica se a linha inicia um bloco de código
		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock
			
			// Imprimir o tipo de linguagem (se especificado)
			langSpec := strings.TrimPrefix(line, "```")
			if inCodeBlock && langSpec != "" {
				color.Yellow("Código %s:", langSpec)
			}
			continue
		}

		// Mudar formatação para amarelo
		if strings.HasPrefix(line, "**") && strings.HasSuffix(line, "**") {
			boldText := strings.TrimPrefix(strings.TrimSuffix(line, "**"), "**")
			// Usar fonte maior e negrito
			fmt.Println()
			line = fmt.Sprintf("\033[1m%s\033[0m", boldText)
		}

		
		// Animação de digitação caractere por caractere
		if inCodeBlock {
			// Código com cor verde
			for _, char := range line {
				color.New(color.FgGreen).Print(string(char))
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println()
		} else {
			// Texto normal
			for _, char := range line {
				fmt.Print(string(char))
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println()
		}
	}
	fmt.Println()
}

func ClearScreen(){
	fmt.Print("\033[H\033[2J")
}

// PrintHelp imprime a ajuda da aplicação
func PrintHelp() {
	color.Yellow("\nComandos disponíveis:")
	fmt.Println("- help/ajuda:      Exibe esta mensagem de ajuda")
	fmt.Println("- sair/exit/quit:   Encerra a aplicação")
    fmt.Println("- clear/cls     	 Limpa a tela")
	fmt.Println()
}

// PrintError imprime um erro formatado
func PrintError(err error) {
	color.Red("Erro: %v", err)
}
