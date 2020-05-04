package main

import (
	"fmt"

	"github.com/EnggarSe/task1-interface/motor"
	"github.com/EnggarSe/task1-interface/sepeda"
	"github.com/EnggarSe/task1-interface/speed"
)

func main() {
	var jarak float32
	var k speed.Maju
	fmt.Print("Masukan Jarak Tempuh : ")
	fmt.Scanln(&jarak)
	fmt.Println("-----------Sepeda-----------------------")
	sepeda1 := sepeda.CreateSepeda(2, 5, "Merah")
	jarakSepeda := sepeda1.WaktuTempuh(jarak)
	fmt.Println(sepeda1)
	fmt.Println("Waktu Tempuh Dengan Jarak", jarak, "Km adalah", jarakSepeda, "menit")
	k = sepeda1
	fmt.Println("Waktu Tempuh Untuk Method Cepat :", k.Cepat(jarakSepeda), "menit")
	fmt.Println("Waktu Tempuh Untuk Method Lambat :", k.Lambat(jarakSepeda), "menit")

	fmt.Println("-----------Motor-----------------------")
	motor1 := motor.CreateMotor(2, 3, "Hitam")
	fmt.Println(motor1)
	fmt.Println("Waktu Tempuh Dengan Jarak", jarak, "Km adalah", motor1.WaktuTempuh(jarakSepeda), "menit")
	k = motor1
	fmt.Println("Waktu Tempuh Untuk Method Cepat :", k.Cepat(motor1.WaktuTempuh(jarakSepeda)), "menit")
	fmt.Println("Waktu Tempuh Untuk Method Lambat :", k.Lambat(motor1.WaktuTempuh(jarakSepeda)), "menit")
	fmt.Println("----------------------------------------")

}
