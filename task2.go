package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var result []BrainrotMeme

	for _, meme := range memes {
		if meme.Views > minViews {
			result = append(result, meme)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TrendLevel > result[j].TrendLevel
	})

	return result
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)

	for _, meme := range memes {
		impact[meme.Category] += meme.Views
	}

	return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var result []string

	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			result = append(result, meme.Name)
		}
	}

	return result
}

func main() {
	memes := []BrainrotMeme{
		{Name: "Sigma Grindset", TrendLevel: 9, Category: "Sigma", Views: 120.5},
		{Name: "Skibidi Toilet", TrendLevel: 10, Category: "Skibidi", Views: 250.0},
		{Name: "Mewing Challenge", TrendLevel: 7, Category: "Mewing", Views: 45.2},
		{Name: "Subo Bratik Dance", TrendLevel: 6, Category: "Subo Bratik", Views: 35.8},
		{Name: "TUNTUNTUNSAHUR Beat", TrendLevel: 8, Category: "TUNTUNTUNSAHUR", Views: 78.3},
		{Name: "Sigma Male Energy", TrendLevel: 5, Category: "Sigma", Views: 55.1},
		{Name: "Random Brainrot", TrendLevel: 4, Category: "Other", Views: 12.5},
		{Name: "Skibidi Rizz", TrendLevel: 9, Category: "Skibidi", Views: 95.7},
	}

	fmt.Println("1. Топ трендовые мемы с просмотрами > 50 миллионов:")
	fmt.Println("(отсортированы по убыванию TrendLevel)\n")
	topTrending := FindTopTrending(memes, 50.0)
	for i, meme := range topTrending {
		fmt.Printf("   %d. %s | TrendLevel: %d | Просмотры: %.1fM | Категория: %s\n",
			i+1, meme.Name, meme.TrendLevel, meme.Views, meme.Category)
	}

	fmt.Println("\n2. Суммарное влияние по категориям (в миллионах просмотров):\n")
	categoryImpact := CalculateCategoryImpact(memes)
	var categories []string
	for cat := range categoryImpact {
		categories = append(categories, cat)
	}
	sort.Strings(categories)
	for _, cat := range categories {
		fmt.Printf("   %s: %.1fM просмотров\n", cat, categoryImpact[cat])
	}

	fmt.Println("\n3. Мемы, удовлетворяющие сложному условию:")
	filtered := FilterByComplexCondition(memes)
	for i, name := range filtered {
		fmt.Printf("   %d. %s\n", i+1, name)
	}
}
