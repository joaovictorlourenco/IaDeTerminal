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

var explainCmd = &cobra.Command{
	Use:   "explain [termo ou conceito]",
	Short: "Explica um termo ou conceito técnico",
	Long:  `Utiliza a API da Groq para explicar termos ou conceitos técnicos de forma clara e concisa.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		explainConcept(args)
	},
}

func explainConcept(args []string) {
	// Combinar todos os argumentos em uma única string
	term := strings.Join(args, " ")
	prompt := fmt.Sprintf("Explique o seguinte conceito de forma clara e concisa: %s", term)
	
	// Iniciar spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Obtendo explicação..."
	s.Color("yellow")
	s.Start()
	
	// Preparar contexto específico para explicações
	explainContext := []api.Message{
		{Role: "system", Content: "Você é um assistente especializado em explicar conceitos técnicos de forma clara, concisa e de fácil compreensão."},
	}
	
	// Consultar a API
	response, err := api.QueryGroq(prompt, explainContext)
	s.Stop()
	
	if err != nil {
		color.Red("Erro: %v", err)
		return
	}
	fmt.Println("Chamou o explain")
	// Imprimir a explicação
	fmt.Println("\nExplicação:")
	color.Cyan("%s", response)
}

func init() {
	RootCmd.AddCommand(explainCmd)
}