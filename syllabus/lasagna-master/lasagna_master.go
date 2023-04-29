package lasagna

func PreparationTime(layers []string, avPrepTimePerLayer int) int {
	if avPrepTimePerLayer == 0 {
		avPrepTimePerLayer = 2
	}

	return len(layers) * avPrepTimePerLayer
}

func Quantities(layers []string) (noodles int, sauce float64) {

	for index := range layers {
		if layers[index] == "noodles" {

			noodles += 50
		}
		if layers[index] == "sauce" {

			sauce += 0.2
		}
	}

	return

}

func AddSecretIngredient(friendRecipe, ownRecipe []string) {

	(ownRecipe)[len(ownRecipe)-1] = friendRecipe[len(friendRecipe)-1]

}

func ScaleRecipe(amount []float64, quantities int) []float64 {

	var scaledAmount []float64
	for index := range amount {

		scaledAmount = append(scaledAmount, (amount[index]*float64(quantities))/2)
	}

	return scaledAmount

}
