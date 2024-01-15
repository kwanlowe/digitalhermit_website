package main

import (
	"github.com/google/generative-ai-go/genai"
	"github.com/jedib0t/go-pretty/v6/text"
	"google.golang.org/api/option"
	"context"
	"os"
	"log"
 	"fmt"
	"flag"
	"bufio"
	"strings"
)

func main() {
	var pathToImage1 string
	reader := bufio.NewReader(os.Stdin)
	userPrompt := text.FgBlue.Sprint("Prompt: ")

	flag.StringVar(&pathToImage1, "i", "image1.jpg", "Image location")
	//pathToImage1 := "images/crowd1.png"	
	flag.Parse()


	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
	  log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
	  &genai.Content{
	    Parts: []genai.Part{
	      genai.Text("Create a Golang function based on the input."),
	      genai.Text("Add informative comments to the code."),
	      genai.Text("Return only the Golang code."),
	    },
	    Role: "user",
	  },
	  &genai.Content{
	    Parts: []genai.Part{
	      genai.Text("Great to meet you. What would you like to know?"),
	    },
	    Role: "model",
	  },
	}

	for {
		fmt.Print(userPrompt)
		prompt, _ := reader.ReadString('\n')
		prompt = strings.ToValidUTF8(prompt, " ")
		resp, err := cs.SendMessage(ctx, genai.Text(prompt))
		if err != nil {
		  log.Fatal(err)
		}
		printResponse(resp)
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
        for _, cand := range resp.Candidates {
                for _, part := range cand.Content.Parts {
                        fmt.Println(part)
                }
        }
        fmt.Println("---")
}
