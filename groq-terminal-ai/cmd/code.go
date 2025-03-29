package cmd

import (
	"fmt"
	"strings"
	"time"

	"iaterminal/groq-terminal-ai/internal/api"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var codeCmd = &cobra.Command{
	Use:   "code [descrição]",
	Short: "Gera código baseado na descrição fornecida",
	Long:  `Utiliza a API da Groq para gerar código com base na descrição fornecida pelo usuário.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generateCode(args)
	},
}

func generateCode(args []string) {
	// Combinar todos os argumentos em uma única string
	description := "Gere código para: " + strings.Join(args, " ")
	
	// Iniciar spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Gerando código..."
	s.Color("yellow")
	s.Start()
	
	// Preparar contexto específico para geração de código
	codeContext := []api.Message{
		{Role: "system", Content: "Você é um assistente especializado em programação. Forneça apenas o código solicitado, sem explicações adicionais, a menos que sejam explicitamente solicitadas."},
	}
	
	// Consultar a API
	response, err := api.QueryGroq(description, codeContext)
	
	s.Stop()
	
	if err != nil {
		color.Red("Erro: %v", err)
		return
	}
	
	// Imprimir o código gerado
	fmt.Println("\nCódigo gerado:")
	color.Cyan("%s", response)
}

func init() {
	RootCmd.AddCommand(codeCmd)
}