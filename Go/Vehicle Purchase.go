package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car" || kind == "truck" {
        return true
    }
    return false
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    if option1 < option2 {
        return option1 + " is clearly the better choice."
    }
    return option2 + " is clearly the better choice."
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var percent float64 = 80.0
    if age >= 3 && age < 10 {
        percent = 70.0;
    }
    if age >= 10 {
        percent = 50.0
    }
    return (originalPrice / 100) * percent
	panic("CalculateResellPrice not implemented")
}
