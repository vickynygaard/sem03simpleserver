package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"github.com/vickynygaard/is105sem03/mycrypt"
	"github.com/vickynygaard/funtemps/conv"
	"strconv"
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
                                                _, err = conn.Write([]byte(string(kryptertMelding))
                                         case "Kjevik;SN39040;18.03.2022 01:50;6":
                                                 FarhenheitString, err := CelsiusToFarhenheitLine(string(dekryptertMelding))
                                                 if err !=  nil  {
                                                         log.Println(err)
                                                }
						kryptertMelding := mycrypt.Krypter([]rune(FarhenheitString), mycrypt.ALF_SEM03, 4)
						log.Println("Kryptert melding: ", string(kryptertMelding)))
                                                _, err = conn.Write([]byte(string("Kryptert melding"))
                                       default:
                                           _, err != c.Write(buf[:n])
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
func  CelsiusToFarhenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
       
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFarhenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}
func conv.CelsiusToFarhenheitLine(line string) (string, error) {

        dividedString := strings.Split(line, ";")
	var err error
	
	if (len(dividedString) == 4) {
		dividedString[3], err = CelsiusToFarhenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil
	
	/*	
	return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
        */
}
