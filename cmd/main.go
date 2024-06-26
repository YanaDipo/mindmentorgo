package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

func startService(name string, cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to start %s service: %v", name, err)
	}
}

func main() {
	var wg sync.WaitGroup

	databaseServiceCmd := exec.Command("go", "run", "mindmentor/services/database_service")
	wg.Add(1)
	go func() {
		startService("Database", databaseServiceCmd)
		wg.Done()
	}()

	// Запуск сервиса emotions
	emotionServiceCmd := exec.Command("go", "run", "mindmentor/services/emotions_service")
	wg.Add(1)
	go func() {
		startService("Emotion", emotionServiceCmd)
		wg.Done()
	}()

	profileServiceCmd := exec.Command("go", "run", "mindmentor/services/profile_service")
	wg.Add(1)
	go func() {
		startService("Profile", profileServiceCmd)
		wg.Done()
	}()

	// Запуск сервиса meditation
	meditationServiceCmd := exec.Command("go", "run", "mindmentor/services/meditation_service")
	wg.Add(1)
	go func() {
		startService("Meditation", meditationServiceCmd)
		wg.Done()
	}()

	// Запуск сервиса trainings
	trainingsServiceCmd := exec.Command("go", "run", "mindmentor/services/social_service")
	wg.Add(1)
	go func() {
		startService("Training", trainingsServiceCmd)
		wg.Done()
	}()

	// Запуск сервиса social
	socialServiceCmd := exec.Command("go", "run", "mindmentor/services/social_service")
	wg.Add(1)
	go func() {
		startService("Social", socialServiceCmd)
		wg.Done()
	}()

	// Запуск сервиса авторизации
	authServiceCmd := exec.Command("go", "run", "mindmentor/services/auth_service")
	wg.Add(1)
	go func() {
		startService("Auth", authServiceCmd)
		wg.Done()
	}()

	// Запуск API Gateway
	gatewayCmd := exec.Command("go", "run", "mindmentor/gateway")
	wg.Add(1)
	go func() {
		startService("Gateway", gatewayCmd)
		wg.Done()
	}()

	wg.Wait()

}
