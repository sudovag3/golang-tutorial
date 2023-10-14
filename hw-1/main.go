package main

import (
	"fmt"
	"io"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	//Получаем корректное время с помощью ntp клиента
	// github.com/beevik/ntp
	time, err := ntp.Time("0.beevik-ntp.poo.ntp.org")
	if err == nil {
		fmt.Println("Hello !")
		fmt.Println(time)
	} else {
		//Проверить в Windows exit code -
		// echo $LastExitCode
		io.WriteString(os.Stderr,
			err.Error())
		os.Exit(1)
	}

}
