package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/mellotonio/desafiogo/app/domain/challenges"
)

func (sgs *ChallengeService) SolveChallenge(ctx context.Context, challenge challenges.Challenge) (string, error) {
	f, err := os.OpenFile("script.py", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(challenge.Answer[0].Text))
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	cmd1 := exec.Command("autopep8", "-i", "script.py")
	_, err = cmd1.Output()
	if err != nil {
		fmt.Println("Error executing command1:", err)
		return "", nil
	}

	cmd := exec.Command("python3", "script.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		cmd := exec.Command("flake8", "script.py")
		errorOutput, _ := cmd.Output()
		fmt.Println("Error executing command:", string(errorOutput))
		output = errorOutput

		chatGPTResponse, chatGPTErr := callChatGPTAPI(challenge.Answer[0].Text, string(output))
		if chatGPTErr != nil {
			return "", chatGPTErr
		}

		output = []byte(fmt.Sprintf("Jarvis (Chatgpt Helper) disse: erro encontrado %s, sugestao: %s", errorOutput, chatGPTResponse))
	}

	return string(output), nil
}

func callChatGPTAPI(answerText, errorOutput string) (string, error) {
	apiURL := "https://api.openai.com/v1/chat/completions"
	apiKey := "sk-proj-AnZRrGIjxPr2iddzcMEpT3BlbkFJhrEOr7cdiA0CRV5sKlhv" // Replace with your actual API key

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": "Analyze the following Python script error and explain what might be wrong:\n" + errorOutput + "the original python code is:" + answerText},
		},
		"max_tokens": 150,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(string(requestBody)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return "", err
	}

	choices, ok := responseMap["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("unexpected response format from ChatGPT API")
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response format from ChatGPT API")
	}

	chatGPTResponse, ok := choice["message"].(map[string]interface{})["content"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format from ChatGPT API")
	}

	return chatGPTResponse, nil
}
