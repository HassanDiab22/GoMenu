package utils

import (
	"fmt"
	"gomenu/models"
	"log"
	"os"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func GenerateQRCode(menu models.Menu) {
	fmt.Println("Generating QR code...")

	refUrl := fmt.Sprintf("https://localhost:3000/menus/%d", menu.ID)
	qrc, err := qrcode.New(refUrl)
	if err != nil {
		log.Fatalf("could not generate QRCode: %v", err)
	}

	qrFilePath := fmt.Sprintf("./assets/QR/%d.jpeg", menu.ID)
	if err := os.MkdirAll("./assets/QR", os.ModePerm); err != nil {
		log.Fatalf("could not create assets directory: %v", err)
	}

	w, err := standard.New(qrFilePath)
	if err != nil {
		log.Fatalf("standard.New failed: %v", err)
	}

	if err = qrc.Save(w); err != nil {
		log.Fatalf("could not save image: %v", err)
	}

	fmt.Println("QR code generated and saved successfully at", qrFilePath)
}
