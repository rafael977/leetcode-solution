package designparkingsystem

/*
 * @lc app=leetcode id=1603 lang=golang
 *
 * [1603] Design Parking System
 */

// @lc code=start
type ParkingSystem struct {
	cap map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		cap: map[int]int{
			1: big,
			2: medium,
			3: small,
		},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.cap[carType] == 0 {
		return false
	}
	this.cap[carType]--
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end
