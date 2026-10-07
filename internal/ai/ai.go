package ai

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/openai/openai-go/v3"
	"gopkg.in/yaml.v3"
)

type PromptConfig struct {
	Model           string `yaml:"model"`
	DeveloperPrompt string `yaml:"developer_prompt"`
	UserPrompt      string `yaml:"user_prompt"`
}

type AI struct {
	client  openai.Client
	prompts map[string]PromptConfig
}

func New(client openai.Client) (*AI, error) {
	prompts, err := loadPrompts("./prompts")
	if err != nil {
		return nil, fmt.Errorf("failed to load prompts: %w", err)
	}

	return &AI{
		client:  client,
		prompts: prompts,
	}, nil
}

func loadPrompts(dirPath string) (map[string]PromptConfig, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	prompts := make(map[string]PromptConfig)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(dirPath, entry.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		var cfg PromptConfig
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", filePath, err)
		}

		promptKey := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))

		prompts[promptKey] = cfg
	}

	return prompts, nil
}
