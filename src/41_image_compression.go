package main

import (
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func images() {
	if err := compressPNG("./images/1.png", -3); err != nil { // 0, -1, -2, -3
		fmt.Println("ERROR:", err)
	}
	if err := compressJPEG("./images/1.jpg", 75); err != nil { // 0 ~ 100
		fmt.Println("ERROR:", err)
	}
}

func compressJPEG(filepath string, quality int) error {
	fmt.Println("Compressing JPEG")
	if quality > 100 {
		quality = 100
	} else if quality < 0 {
		quality = 0
	}

	img, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer img.Close()

	imgData, err := jpeg.Decode(img)
	if err != nil {
		return err
	}

	outputImg, err := os.Create("./images/output.jpg")
	if err != nil {
		return err
	}
	defer outputImg.Close()

	options := &jpeg.Options{Quality: quality}
	jpeg.Encode(outputImg, imgData, options)

	return nil
}

func compressPNG(filepath string, compressLevel int) error {
	if compressLevel > 0 || compressLevel < -3 {
		return errors.New("invalid compression level range")
	}
	fmt.Println("Compressing PNG")
	img, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer img.Close()

	imgData, err := png.Decode(img)
	if err != nil {
		return err
	}

	outputImg, err := os.Create("./images/output.png")
	if err != nil {
		return err
	}
	defer outputImg.Close()

	enc := &png.Encoder{CompressionLevel: png.CompressionLevel(compressLevel)}
	enc.Encode(outputImg, imgData)

	return nil
}
