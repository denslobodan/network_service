package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Network service with quotes Golang
const addr = "0.0.0.0:12345"
const proto = "tcp4"

func main() {
	// Подключение к сетевой службе.
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Не забываем закрыть ресурс.
	defer conn.Close()

	// Буфер для чтения данных из соединения.
	reader := bufio.NewReader(conn)
	// Считывание массива байт до перевода строки.
	for {
		b, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Обработка ответа.
		fmt.Print(string(b))
	}

}
