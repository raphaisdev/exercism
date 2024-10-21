package chance

import (
    "math/rand"
    "time"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
    return 1 + rand.Intn(20-1+1)
	panic("Please implement the RollADie function")
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    return float64(rand.Intn(1200)) / 100.00
	panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
    a := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 10; i++ {
        rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
    }
    return a

	panic("Please implement the ShuffleAnimals function")
}
