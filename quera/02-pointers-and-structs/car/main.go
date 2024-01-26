package main

import "fmt"

type Car struct {
	speed, battery int
}

func NewCar(speed, battery int) *Car {
	return &Car{speed, battery}
}
func GetSpeed(car *Car) int {
	return car.speed
}
func GetBattery(car *Car) int {
	return car.battery
}
func ChargeCar(car *Car, minutes int) {
	battery := car.battery + (minutes / 2)
	if battery > 100 {
		battery = 100
	}
	car.battery = battery
}
func TryFinish(car *Car, distance int) string {
	neededBattery := distance / 2
	if neededBattery > car.battery {
		car.battery = 0
		return ""
	}
	car.battery -= neededBattery
	return fmt.Sprintf("%.2f", float64(distance)/float64(car.speed))
}
