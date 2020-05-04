package main

import (
	"fmt"

	"github.com/EnggarSe/task1-interface/motor"
	"github.com/EnggarSe/task1-interface/sepeda"
)

func main() {
	var jarak float32
	fmt.Print("Masukan Jarak Tempuh : ")
	fmt.Scanln(&jarak)
	fmt.Println("-----------Sepeda-----------------------")
	sepeda1 := sepeda.CreateSepeda(2, 5, "Merah")
	jarakSepeda := sepeda1.WaktuTempuh(jarak)
	fmt.Println(sepeda1)
	fmt.Println("Waktu Tempuh Dengan Jarak", jarak, "Km adalah", jarakSepeda, "menit")
	fmt.Println("Waktu Tempuh Untuk Method Cepat :", sepeda1.Cepat(jarakSepeda), "menit")
	fmt.Println("Waktu Tempuh Untuk Method Lambat :", sepeda1.Lambat(jarakSepeda), "menit")

	fmt.Println("-----------Motor-----------------------")
	motor1 := motor.CreateMotor(2, 3, "Hitam")
	fmt.Println(motor1)
	fmt.Println("Waktu Tempuh Dengan Jarak", jarak, "Km adalah", motor1.WaktuTempuh(jarakSepeda), "menit")
	fmt.Println("Waktu Tempuh Untuk Method Cepat :", motor1.Cepat(motor1.WaktuTempuh(jarakSepeda)), "menit")
	fmt.Println("Waktu Tempuh Untuk Method Lambat :", motor1.Lambat(motor1.WaktuTempuh(jarakSepeda)), "menit")
	fmt.Println("----------------------------------------")

}
