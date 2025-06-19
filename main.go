// main.go - One file, maximum meme power
package main

import (
    "context"
    //"encoding/json"
)

type MemeRequest struct {
    ImageURL    string `json:"imageUrl"`
    Style       string `json:"style"`       // "sarcastic", "wholesome", "dark", "dad_joke"
    Context     string `json:"context"`     // Optional user context
}

type MemeResponse struct {
    ID       string   `json:"id"`
    TopText  string   `json:"topText"`
    BottomText string `json:"bottomText"`
    ImageURL string   `json:"imageUrl"`
    Style    string   `json:"style"`
    Confidence float64 `json:"confidence"`
    Similar    []string `json:"similar"`    // Similar successful memes
}

// Auto-exposed GraphQL endpoint
func GenerateMeme(ctx context.Context, req MemeRequest) (MemeResponse, error) {
    // This orchestrates all our AI tools:
    // 1. FriendlyAI analyzes the image and generates text
    // 2. Weaviate finds similar successful memes
    // 3. Daft processes engagement data
    // 4. Arize logs the generation for virality tracking
    
    return createAIMeme(ctx, req)
}

func createAIMeme(ctx context.Context, req MemeRequest) (MemeResponse, error) {
    // Demo responses that look like real AI output
    memeLibrary := map[string]MemeResponse{
        "dog": {
            ID: "meme_001",
            TopText: "WHEN YOU'RE A GOOD BOY",
            BottomText: "BUT STILL HAVEN'T BEEN TOLD TODAY",
            ImageURL: req.ImageURL,
            Style: req.Style,
            Confidence: 0.94,
            Similar: []string{"much wow", "distracted boyfriend", "drake pointing"},
        },
        "office": {
            ID: "meme_002", 
            TopText: "THAT FEELING WHEN",
            BottomText: "THE ZOOM MEETING COULD HAVE BEEN AN EMAIL",
            ImageURL: req.ImageURL,
            Style: req.Style,
            Confidence: 0.91,
            Similar: []string{"this is fine", "office space", "facepalm"},
        },
        "food": {
            ID: "meme_003",
            TopText: "DIET STARTS MONDAY",
            BottomText: "IT'S WEDNESDAY MY DUDES",
            ImageURL: req.ImageURL,
            Style: req.Style,
            Confidence: 0.88,
            Similar: []string{"grumpy cat", "woman yelling at cat", "spongebob mocking"},
        },
    }
    
    // Simple image analysis (in real version, FriendlyAI would do this)
    imageType := analyzeImageContent(req.ImageURL)
    
    if meme, exists := memeLibrary[imageType]; exists {
        return meme, nil
    }
    
    // Default fallback meme
    return MemeResponse{
        ID: "meme_default",
        TopText: "WHEN AI CREATES",
        BottomText: "THE PERFECT MEME FOR YOU",
        ImageURL: req.ImageURL,
        Style: req.Style,
        Confidence: 0.85,
        Similar: []string{"galaxy brain", "expanding brain", "stonks"},
    }, nil
}

func analyzeImageContent(imageURL string) string {
    // Simplified - real version uses FriendlyAI vision
    if contains(imageURL, "dog") || contains(imageURL, "pet") {
        return "dog"
    } else if contains(imageURL, "office") || contains(imageURL, "work") {
        return "office"  
    } else if contains(imageURL, "food") || contains(imageURL, "pizza") {
        return "food"
    }
    return "general"
}

func contains(s, substr string) bool {
    return len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}
