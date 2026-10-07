package ai

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3/responses"
)

func (ai *AI) SummarizeTeamReport(ctx context.Context) (teamReport string, err error) {
	prompt, ok := ai.prompts["summorizing-reports"]
	if !ok {
		return "", fmt.Errorf("unknown prompt")
	}

	response, err := ai.client.Responses.New(context.Background(), responses.ResponseNewParams{
		Model: prompt.Model,
		Reasoning: responses.ReasoningParam{
			Effort: responses.ReasoningEffortLow,
		},
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: responses.ResponseInputParam{
				responses.ResponseInputItemParamOfMessage(
					prompt.DeveloperPrompt,
					responses.EasyInputMessageRoleDeveloper,
				),
				responses.ResponseInputItemParamOfMessage(
					prompt.UserPrompt,
					responses.EasyInputMessageRoleUser,
				),
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to get ai response: %v", err)
	}

	return response.OutputText(), nil
}
