package main

import "fmt"

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func main() {
	nums := []int{1, 2, 3, 4}
	c1 := make(chan int, len(nums)) //пишем сюда числа
	c2 := make(chan int, len(nums)) // сюда пишем x*2
	go func() {
		defer close(c1)
		for _, num := range nums {
			writeToChannel(c1, num)
		}
	}()

	go func() {
		defer close(c2)
		calculate(c1, c2)
	}()

	for doubleNum := range c2 {
		fmt.Println(doubleNum)
	}
}

func writeToChannel(c1 chan int, num int) {
	c1 <- num
}

func calculate(c1 chan int, c2 chan int) {
	for num := range c1 {
		c2 <- num * 2
	}
}

//в го - если функция - сендер(посылает данные в каннал), то ей можно закрывать канал.
// но если функция ресивер - запрещено

// select в языке Go позволяет выбирать из нескольких каналов операции отправки или приёма данных. Он предоставляет возможность ожидать одной из нескольких операций, которая готова к выполнению, и выполнять соответствующий код в зависимости от этой операции.

// В контексте данной программы, select используется для безопасной отправки данных в канал c2 в функции calculate. В случае, если чтение из канала c2 заблокировано (например, если канал c2 заполнен), select сразу же завершится и выполнится инструкция return, что приведёт к закрытию канала c2. Это помогает избежать блокировки программы в случае, если чтение из канала c1 завершится раньше, чем данные будут полностью обработаны и отправлены в канал c2.