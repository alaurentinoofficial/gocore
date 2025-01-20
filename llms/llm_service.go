package llms

import "context"

type LlmService interface {
	SinglePrompt(ctx context.Context, prompt string) (string, error)
}
