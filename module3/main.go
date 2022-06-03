package main

import (
	"city"
	newcolor "color"
	"digit"
	"fmt"
	. "wordz" //Добавляем пакет wordz через точку

	"github.com/fatih/color"
	"github.com/huandu/xstrings"
	_ "github.com/huandu/xstrings"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world")
	color.Red("Hello world again")

	fmt.Println(Hello)                         //Вызов переменной из пакета wordz
	fmt.Println(Random())                      //Вызов функции из пакета wordz
	fmt.Println(city.City())                   //Вызов функции из пакета city
	fmt.Println(digit.Digit())                 //Вызов функции из пакета digit
	fmt.Println(xstrings.Shuffle(city.City())) //Вызов функции из пакета xstrings

}
