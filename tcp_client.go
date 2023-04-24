package main

import (
	"github.com/Tiril14/is105sem03/mycrypt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.4:8000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("os.Args[1] = ", os.Args[1])

	// Encrypt the message
	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)

	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt the response from the server
	deCrypt := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
	log.Println("Dekrypter melding: ", string(deCrypt))

	response := string(deCrypt)
	log.Printf("reply from proxy: %s", response)

}
