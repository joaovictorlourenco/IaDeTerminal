package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"iaterminal/groq-terminal-ai/internal/api"
	"iaterminal/groq-terminal-ai/internal/ui"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var contextMessages []api.Message

var RootCmd = &cobra.Command{
	Use:   "groq-ai",
	Short: "Terminal IA usando a API da Groq",
	Long:  `Uma aplicação de terminal inteligente que utiliza a API da Groq para processamento de linguagem natural.`,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveMode()
	},
}

func interactiveMode() {
	
	ui.PrintLogo()

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("\n> ")
		scanner.Scan()
		userInput := scanner.Text()
		lowerInput := strings.ToLower(userInput)
		
		if lowerInput == "sair" || lowerInput == "exit" || lowerInput == "quit" {
			break
		}

		if lowerInput == "help" || lowerInput == "ajuda" {
			ui.PrintHelp()
			continue
		}

		if lowerInput == "clear" || lowerInput == "cls" || lowerInput == "limpar" {
			// Limpar a tela
			ui.ClearScreen()
			continue
		}


		// Iniciar spinner para feedback visual
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Processando..."
		s.Color("yellow")
		s.Start()
		
		// Consultar a API
		response, err := api.QueryGroq(userInput, contextMessages)


		s.Stop()
		
		if err != nil {
			ui.PrintError(err)
			continue
		}
		
		// Adicionar a resposta ao contexto
		contextMessages = append(contextMessages, api.Message{Role: "user", Content: userInput})
		contextMessages = append(contextMessages, api.Message{Role: "assistant", Content: response})
		
		// Limitar o histórico de contexto (opcional, para evitar tokens excessivos)
		if len(contextMessages) > 10 {
			contextMessages = contextMessages[2:] // Remove as duas mensagens mais antigas
		}
		
		// Extrair blocos de código (opcional)
		ui.PrintResponse(response)
	}
}

func init() {
	// Inicializar o contexto
	contextMessages = make([]api.Message, 0)
}