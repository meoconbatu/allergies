package allergies

const testVersion = 1

var items = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(input uint) []string {
	output := make([]string, 0)
	t := input / 128
	if t > 1 {
		if t%2 != 0 {
			t = t - 1
		}
		input = input - 128*t
	}
	for i := uint(128); i >= 1; i = i / 2 {
		t = input / i
		if t >= 1 {
			output = append(output, items[i])
			input = input - t*i
		}
	}
	return output
}

func AllergicTo(input uint, allergen string) bool {
	allergens := Allergies(input)
	for _, a := range allergens {
		if allergen == a {
			return true
		}
	}
	return false
}
