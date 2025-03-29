package main

import (
	"fmt"
	"iaterminal/groq-terminal-ai/cmd"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	_ = godotenv.Load()
	
	// Executar o comando raiz
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}