package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 2
	}
	return len(layers) * avgTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		default:
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsRecipe, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendsRecipe[len(friendsRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsNeededPerPortion []float64, portions int) (totalAmounts []float64) {
	for i := 0; i < len(amountsNeededPerPortion); i++ {
		totalAmounts = append(totalAmounts, amountsNeededPerPortion[i]*float64(portions)/2)
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
