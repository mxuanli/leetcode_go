package main

type ParkingSystem struct {
	carMap map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	carMap := make(map[int]int)
	carMap[1] = big
	carMap[2] = medium
	carMap[3] = small
	p := ParkingSystem{carMap}
	return p
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.carMap[carType] > 0 {
		this.carMap[carType] -= 1
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
