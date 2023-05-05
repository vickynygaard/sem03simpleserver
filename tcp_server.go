package main

import (
	"io"
	"log"
	"net"
	"sync"
	"github.com/vickynygaard/is105sem03/mycrypt"
	"strings"
	"errors"

)

func main() {

	var wg sync.WaitGroup

	server, err := net.Listen("tcp", "172.17.0.5:8080")
	if err != nil {
		log.Fatal(err)
	}
        log.Printf("bundet til %s", server.Addr().String())
        wg.Add(1)

        go func() {
                defer wg.Done()
                for {
                        log.Println("f  r server.Accept() kallet")
                        conn, err := server.Accept()
                        if err != nil {
                                return
                        }
                        go func(c net.Conn) {
                                defer c.Close()
                                for {
                                        buf := make([]byte, 1024)
                                        n, err := c.Read(buf)
                                        if err != nil {
                                                if err != io.EOF {
         log.Println(err)
                                                }
                                                return // fra for l  kke
                                        }
                                        dekryptertMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
                                         log.Println("Dekryptert melding: ", string(dekryptertMelding))
                                        switch msg := string(dekryptertMelding); msg {
                                         case "ping":
                                                kryptertMelding := mycrypt.Krypter([]rune("pong"), mycrypt.ALF_SEM03, 4)
                                                log.Println("Kryptert melding: ", string(kryptertMelding))
                                                _, err = conn.Write([]byte(string(kryptertMelding)))
                                         case "Kjevik;SN39040;18.03.2022 01:50;6":
                                                 FarhenheitString, err := CelsiusToFarhenheitLine(string(dekryptertMelding))
                                                 if err !=  nil  {
                                                         log.Println(err)
                                                }
                                                _, err = conn.Write([]byte(string(FarhenheitString)))
                                       default:
                                                 _, err = c.Write(buf[:n])
            
                                          
                 
                                        } 
                                        if err != nil {
                                                if err != io.EOF {
                                                        log.Println(err)
                                                }
                                                return // fra for l  kke
                                        }
                                }
                        }(conn)
                }
        }()
        wg.Wait()

}
