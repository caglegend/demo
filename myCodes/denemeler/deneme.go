package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	currentTime := time.Now()

	// Şu anki tarih ve saat bilgisini yazdırma
	fmt.Println(currentTime)

	// Özel bir biçimde yazdırma
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(formattedTime)

	sformattedTime := currentTime.Format(time.RFC3339)
	fmt.Println(sformattedTime)

	fmt.Println("Architecture:", runtime.GOARCH)
}
