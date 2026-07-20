package main

import "fmt"

/** Выбрать один вариант
 * 1) Ошибка компиляции
 * 2) Ошибка runtime
 * 3) Напечатает 777
 * 4) Напечатает 0
 */

func main() {
  c := make(chan int)
  write(c, 777)
  fmt.Println(<-c)
}

func write(c chan<- int, number int) {
  c <- number
}



// напечатает 777, в функции начинаем записывать число 777 в канал и читаем из канала в вывод