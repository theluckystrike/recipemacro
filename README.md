# recipemacro

`recipemacro` is a tiny Go library for per-serving recipe nutrition:
it totals calories and macros from per-100 g ingredient data, then
divides the result across servings.

The math is the same kind behind the free
[recipe nutrition calculator](https://recipetool.net/) at
RecipeTool, a calorie and macro counter you can paste any recipe
into; this library gives you the same arithmetic in Go so you can
compute it locally.

## Usage

```go
ings := []recipemacro.Ingredient{
    {Name: "chicken breast", Grams: 500, CalPer100g: 165, ProteinPer100g: 31},
}
t := recipemacro.Sum(ings)
per := t.PerServing(2) // 412.5 kcal, 77.5 g protein per serving
```

## License

MIT
