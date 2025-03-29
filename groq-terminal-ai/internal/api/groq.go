package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

type GroqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// QueryGroq envia uma solicitação para a API da Groq
func QueryGroq(prompt string, context []Message) (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	
	if apiKey == "" {
		return "", fmt.Errorf("API key não encontrada. Configure a variável GROQ_API_KEY no arquivo .env")
	}

	// Adicionar a mensagem atual ao contexto
	messages := append(context, Message{Role: "user", Content: prompt})

	reqBody, err := json.Marshal(GroqRequest{
		Model:       "llama3-70b-8192", // ou outro modelo de sua escolha
		Messages:    messages,
		Temperature: 0.7,
		MaxTokens:   2000,
	})
	if err != nil {
		return "", fmt.Errorf("erro ao criar corpo da requisição: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(reqBody))
	
	if err != nil {
		return "", fmt.Errorf("erro ao criar requisição: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("erro na requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %v", err)
	}

	var groqResp GroqResponse
	if err := json.Unmarshal(body, &groqResp); err != nil {
		return "", fmt.Errorf("erro ao processar resposta: %v\nResposta bruta: %s", err, string(body))
	}

	// Verificar se há erro na resposta da API
	if groqResp.Error != nil && groqResp.Error.Message != "" {
		return "", fmt.Errorf("erro da API Groq: %s", groqResp.Error.Message)
	}

	if len(groqResp.Choices) > 0 {
		return groqResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("resposta vazia da API")
}