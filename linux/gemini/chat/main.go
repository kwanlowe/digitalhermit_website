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
	// "bufio"
	"sort"
	"strings"
	"golang.org/x/term"
	"github.com/fatih/color"
	markdown "github.com/MichaelMure/go-term-markdown"
	"gemini-chat/library"
	"github.com/petermattis/prompt"
)
const (
	appname string = "gemini-chat"
)

var geminiKeywords = []string{
        "QUIT",
        "EXPLAIN",
        "TRANSLATE",
}

func init() {
        sort.Strings(geminiKeywords)
}


func inputFinished(text string) bool {
        text = strings.TrimSpace(text)
        return strings.HasSuffix(text, ";")
}

func completer(text []rune, wordStart, wordEnd int) []string {
        word := strings.ToUpper(string(text[wordStart:wordEnd]))
        i := sort.Search(len(geminiKeywords), func(i int) bool {
                return geminiKeywords[i] >= word
        })
        if i >= len(geminiKeywords) {
                return nil
        }
        word += "\xff"
        j := sort.Search(len(geminiKeywords), func(i int) bool {
                return geminiKeywords[i] >= word
        })
        return geminiKeywords[i:j]
}


func main() {
	var pathToImage1, configFile string
	// reader := bufio.NewReader(os.Stdin)
	userPrompt := text.FgBlue.Sprint("Prompt: ")

	flag.StringVar(&configFile, "c", "config", "Configuration File")
	flag.StringVar(&pathToImage1, "i", "image1.jpg", "Image location")
	flag.Parse()

	config, _ := readConfig.LoadConfig(configFile, appname)

	modelType := config.GetString("gemini.model")

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
	  log.Fatal(err)
	}
	defer client.Close()
	
	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel(modelType)
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
	  &genai.Content{
	    Parts: []genai.Part{
	      genai.Text("Hello."),
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

	 	p, err := prompt.New(
                	prompt.WithCompleter(completer),
                	prompt.WithHistory(os.ExpandEnv("${HOME}/.gemini-chat_history"), -1),
                	prompt.WithInputFinished(inputFinished))

		if err != nil {
			log.Fatal(err)
		}
        
		// geminiPrompt, _ := reader.ReadString('\n')
		geminiPrompt, err := p.ReadLine("gemini> ")
		if err != nil {
			log.Fatal(err)
		}	

		geminiPrompt = strings.ToValidUTF8(geminiPrompt, " ")
		if (len(geminiPrompt)) > 1 {
			resp, err := cs.SendMessage(ctx, genai.Text(geminiPrompt))
			if err != nil {
			  	log.Fatal(err)
			}
			printResponse(resp)
		}
	}
}

func printResponse(resp *genai.GenerateContentResponse) {
	var output string
	var markdownText, geminiPrompt string
	geminiPrompt = text.FgRed.Sprint("Gemini:")

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				output = fmt.Sprint(part)
				width, _, _ := term.GetSize(0) 
				width = width - len(geminiPrompt)
				markdownText = string(markdown.Render(output, width, 6))
				for _, line := range strings.Split(markdownText, "\n") {
					color.Blue(line)
				}
			}
		}
	}
	fmt.Println("---")
}
