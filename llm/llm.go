package llm

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/cfindlayisme/whatcanimake/model"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
)

func GetRecipes(ingrediants string, count int) (model.Recipes, error) {

	llm, err := openai.NewChat(openai.WithModel("gpt-3.5-turbo"))
	if err != nil {
		log.Fatal(err)
	}
	systemPrompt := "You are a backend REST API that responds to requests with a list of ingrediants with " + strconv.Itoa(count) + " recipes you can make with those ingredients. You respond with a JSON format provided in the examples below based off the ingrediants provided by the user, but do not provide just the examples back to the user. Do not use ingredients that are not provided by the user, and do not use ingredients from the examples unless they are also provided by the user as they only have the list of ingredients they have provided."

	bytes, err := os.ReadFile("model/mocks/mocks.json")
	if err != nil {
		log.Fatal(err)
	}

	mockData := string(bytes)

	completion, err := llm.Call(context.Background(), []schema.ChatMessage{
		schema.SystemChatMessage{Content: systemPrompt},
		schema.SystemChatMessage{Content: mockData},
		schema.HumanChatMessage{Content: ingrediants},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(completion.Content)

	recipesRecieved := model.Recipes{}
	err = json.Unmarshal([]byte(completion.Content), &recipesRecieved)
	if err != nil {
		log.Fatal(err)
	}

	return recipesRecieved, nil
}
