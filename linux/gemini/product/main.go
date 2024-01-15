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
	      genai.Text("You are a product manager targeting a Gen-Z audience."),
	      genai.Text("Create exciting advertising copy for products and their simple description."),
	      genai.Text("Keep copy under a few sentences long."),
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
	var output string
	var formattedText, geminiPrompt string
	geminiPrompt = text.FgRed.Sprint("")

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				output = fmt.Sprint(part)
				formattedText = text.WrapSoft(output, 80)
				for _, line := range strings.Split(formattedText, "\n") {
					// fmt.Printf("Gemini: %s\n", line)
					fmt.Printf("%s %s\n",geminiPrompt, line)
				}
			}
		}
	}
	fmt.Println("---")
}
