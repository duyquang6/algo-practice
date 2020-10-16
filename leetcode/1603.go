// https://leetcode.com/problems/design-parking-system
package leetcode

type ParkingSystem struct {
	slots [3]uint16
}

func constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		slots: [3]uint16{uint16(big), uint16(medium), uint16(small)},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType-1] > 0 {
		this.slots[carType-1]--
		return true
	}
	return false
}
