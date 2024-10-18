package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        timePerLayer = 2
    }
    return len(layers) * timePerLayer
    panic("Please implement the function")
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    for _, v := range layers {
        if v == "noodles" {
            noodles += 50
            continue
        }
        if v == "sauce" {
            sauce += 0.2
            continue
        }
    }
    return  noodles, sauce
    panic("Please implement the function")
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
    return append(myList[:len(myList)-1], friendsList[len(friendsList)-1])
    panic("Please implement the function")
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(layers []float64, portions int) []float64 {
    var scaled []float64
    for _, v := range layers {
		scaled = append(scaled, v / float64(2) * float64(portions))
    }
    return scaled
    panic("Please implement the function")
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
