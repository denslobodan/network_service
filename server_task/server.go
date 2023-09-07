package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const addr = "0.0.0.0:12345"
const proto = "tcp4"

var qoutes = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		conn.Write([]byte(getRandomQuote(qoutes) + "\n"))
	}

}

func getRandomQuote(quotes []string) string {
	if len(quotes) == 0 {
		return ""
	}
	time.Sleep(3 * time.Second)
	return quotes[rand.Intn(len(quotes))]
}

func main() {
	// Запуск сетевой службы по протоколу TCP
	// на порту 12345.
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Подключения обрабатываются в бесконечном цикле.
	// Иначе после обслуживания первого подключения сервер
	//завершит работу.
	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
		fmt.Println("Connected by", conn.RemoteAddr())
	}

}
