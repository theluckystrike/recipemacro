package recipemacro

import (
	"math"
	"testing"
)

func near(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

// Per-100 g values below are the USDA-derived rows published in the
// "macros per 100 g of common ingredients" table on RecipeTool's recipe
// macro calculator page; the expected whole-recipe grams are the rows of
// that page's chicken-and-rice meal-prep worked example.
var (
	chickenCooked = Ingredient{Name: "chicken breast, cooked", CalPer100g: 165, ProteinPer100g: 31, CarbsPer100g: 0, FatPer100g: 3.6}
	riceCooked    = Ingredient{Name: "white rice, cooked", CalPer100g: 130, ProteinPer100g: 2.7, CarbsPer100g: 28.2, FatPer100g: 0.3}
	oliveOil      = Ingredient{Name: "olive oil", CalPer100g: 900, ProteinPer100g: 0, CarbsPer100g: 0, FatPer100g: 100}
)

func with(in Ingredient, grams float64) Ingredient { in.Grams = grams; return in }

func TestSumMatchesPublishedRows(t *testing.T) {
	cases := []struct {
		in      Ingredient
		p, c, f float64
	}{
		{with(chickenCooked, 600), 186.0, 0.0, 21.6},
		{with(riceCooked, 3*158), 12.8, 133.7, 1.4},
		{with(oliveOil, 2*13.5), 0.0, 0.0, 27.0},
	}
	for _, tc := range cases {
		got := Sum([]Ingredient{tc.in})
		if !near(got.Protein, tc.p, 0.05) || !near(got.Carbs, tc.c, 0.05) || !near(got.Fat, tc.f, 0.05) {
			t.Errorf("%s: got P%.2f C%.2f F%.2f, want P%.1f C%.1f F%.1f",
				tc.in.Name, got.Protein, got.Carbs, got.Fat, tc.p, tc.c, tc.f)
		}
	}
}

func TestPerServing(t *testing.T) {
	tot := Sum([]Ingredient{with(chickenCooked, 500)})
	per := tot.PerServing(2)
	if !near(per.Calories, 412.5, 1e-9) || !near(per.Protein, 77.5, 1e-9) {
		t.Fatalf("got %+v", per)
	}
}

func TestPerServingGuardsZeroAndNegative(t *testing.T) {
	tot := Totals{Calories: 100, Protein: 10, Carbs: 5, Fat: 2}
	for _, s := range []int{0, -3} {
		if got := tot.PerServing(s); got != tot {
			t.Errorf("PerServing(%d) = %+v, want unchanged totals", s, got)
		}
	}
}

func TestEmptyRecipeIsZero(t *testing.T) {
	if got := Sum(nil); got != (Totals{}) {
		t.Fatalf("Sum(nil) = %+v", got)
	}
}
