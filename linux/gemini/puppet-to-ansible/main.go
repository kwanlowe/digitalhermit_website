package main

import (
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"context"
	"os"
	"log"
	"fmt"
	"flag"
)

func main() {
	var pathToPuppetCode string
	flag.StringVar(&pathToPuppetCode, "i", "example.pp", "Puppet code location")
	flag.Parse()

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
	  log.Fatal(err)
	}
	defer client.Close()

	// For text-and-image input (multimodal), use the gemini-pro-vision model
	model := client.GenerativeModel("gemini-pro")

	puppetCode, err := os.ReadFile(pathToPuppetCode)
	if err != nil {
	  log.Fatal(err)
	}


	prompt := []genai.Part{
	  genai.Text(puppetCode),
	  genai.Text("This is a snippet of puppet automation code. Convert this to Ansible code."),
	}
	resp, err := model.GenerateContent(ctx, prompt...)

	if err != nil {
	  log.Fatal(err)
	}
	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
