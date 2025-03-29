package config

import (
	"os"
	"strconv"
)

// Configuration contém as configurações da aplicação
type Configuration struct {
	Model            string
	Temperature      float64
	MaxTokens        int
	ContextSize      int
	DefaultPrompt    string
	GroqApiEndpoint  string
}

// GetConfig carrega as configurações do ambiente ou usa valores padrão
func GetConfig() Configuration {
	config := Configuration{
		Model:           getEnvOrDefault("GROQ_MODEL", "llama3-70b-8192"),
		Temperature:     getEnvAsFloatOrDefault("GROQ_TEMPERATURE", 0.7),
		MaxTokens:       getEnvAsIntOrDefault("GROQ_MAX_TOKENS", 2000),
		ContextSize:     getEnvAsIntOrDefault("GROQ_CONTEXT_SIZE", 10),
		DefaultPrompt:   getEnvOrDefault("GROQ_DEFAULT_PROMPT", "Você é um assistente de IA útil e conciso."),
		GroqApiEndpoint: getEnvOrDefault("GROQ_API_ENDPOINT", "https://api.groq.com/openai/v1/chat/completions"),
	}
	
	return config
}

// Função auxiliar para obter variável de ambiente ou valor padrão
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Função auxiliar para obter variável de ambiente como float ou valor padrão
func getEnvAsFloatOrDefault(key string, defaultValue float64) float64 {
	strValue := os.Getenv(key)
	if strValue == "" {
		return defaultValue
	}
	
	value, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return defaultValue
	}
	
	return value
}

// Função auxiliar para obter variável de ambiente como int ou valor padrão
func getEnvAsIntOrDefault(key string, defaultValue int) int {
	strValue := os.Getenv(key)
	if strValue == "" {
		return defaultValue
	}
	
	value, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue
	}
	
	return value
}