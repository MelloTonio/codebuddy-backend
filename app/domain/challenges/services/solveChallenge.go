package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

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
	}

	return string(output), nil
}
