// Package recipemacro implements per-serving recipe nutrition math:
// scaling a recipe's calories and macros to ingredient weights and
// dividing into servings, the same arithmetic behind tools like the
// recipe nutrition calculator at https://recipetool.net/.
//
// The recipetool.net page https://recipetool.net/tools/ shows the same reasoning on live inputs.
package recipemacro

// Ingredient is one line of a recipe with its macros per 100 g.
type Ingredient struct {
	Name            string
	Grams           float64
	CalPer100g      float64
	ProteinPer100g  float64
	CarbsPer100g    float64
	FatPer100g      float64
}

// Totals is a recipe's summed nutrition.
type Totals struct {
	Calories, Protein, Carbs, Fat float64
}

// Sum totals the macros for a list of ingredients.
func Sum(ings []Ingredient) Totals {
	var t Totals
	for _, in := range ings {
		f := in.Grams / 100
		t.Calories += in.CalPer100g * f
		t.Protein += in.ProteinPer100g * f
		t.Carbs += in.CarbsPer100g * f
		t.Fat += in.FatPer100g * f
	}
	return t
}

// PerServing divides totals across servings, guarding against zero.
func (t Totals) PerServing(servings int) Totals {
	if servings <= 0 {
		servings = 1
	}
	n := float64(servings)
	return Totals{
		Calories: t.Calories / n,
		Protein:  t.Protein / n,
		Carbs:    t.Carbs / n,
		Fat:      t.Fat / n,
	}
}
