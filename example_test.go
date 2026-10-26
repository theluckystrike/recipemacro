package recipemacro_test

import (
	"fmt"

	"github.com/theluckystrike/recipemacro"
)

// A four-box meal prep: 600 g cooked chicken breast, 3 cups (474 g) cooked
// white rice and 2 tbsp (27 g) olive oil, split four ways.
func ExampleTotals_PerServing() {
	recipe := []recipemacro.Ingredient{
		{Name: "chicken breast, cooked", Grams: 600, CalPer100g: 165, ProteinPer100g: 31, FatPer100g: 3.6},
		{Name: "white rice, cooked", Grams: 474, CalPer100g: 130, ProteinPer100g: 2.7, CarbsPer100g: 28.2, FatPer100g: 0.3},
		{Name: "olive oil", Grams: 27, CalPer100g: 900, FatPer100g: 100},
	}
	whole := recipemacro.Sum(recipe)
	per := whole.PerServing(4)
	fmt.Printf("whole:   %.1f kcal  P %.1f g  C %.1f g  F %.1f g\n", whole.Calories, whole.Protein, whole.Carbs, whole.Fat)
	fmt.Printf("per box: %.1f kcal  P %.1f g  C %.1f g  F %.1f g\n", per.Calories, per.Protein, per.Carbs, per.Fat)
	// Output:
	// whole:   1849.2 kcal  P 198.8 g  C 133.7 g  F 50.0 g
	// per box: 462.3 kcal  P 49.7 g  C 33.4 g  F 12.5 g
}
