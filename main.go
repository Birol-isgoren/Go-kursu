package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	//slices.Demo2()
	//functions.SelamVer()
	//var sayi = functions.Topla(2, 6)
	//fmt.Println("Sonuc", sayi)

	//sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)

	//fmt.Println(sonuc1, sonuc2, sonuc3)
	//var sonuc = functions.ToplaVariadic(3, 4, 5, 6)

	//sayi := 20
	//pointers.Demo1(&sayi)
	//fmt.Println(sayi)

	//sayilar := []int{1, 2, 3}
	//pointers.Demo2(sayilar)
	//fmt.Println("Maindeki sayi : ", sayilar[0])
	//structs.Demo2()

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım : ", carpim)

	//defer_statement.Demo2()
	//error_handling.Demo1()

	//error_handling.Demo2()

	//fmt.Println(error_handling.TahmitEt2(50))
	//string_functions.Demo2()
	//restful.Demo2()
	//project.GetAllProducts()
	project.AddProduct()
	products, err := project.GetAllProducts()
	if err == nil {
		for i := 0; i < len(products); i++ {
			fmt.Println(products[i].ProductName)
		}
	}
}
