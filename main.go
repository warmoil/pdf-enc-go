package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	//outPath := "./out"
	//inPath := "./in"
	passPath := "./pw"

	pwModelList := readPWModelsFromFile(passPath + "/pass.txt")
	for _, pdfModel := range pwModelList {
		encryptPDFBy(pdfModel.Name, pdfModel.PW)
	}
}

func readPWModelsFromFile(filePath string) []PDFModel {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return nil
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	pwModelList := make([]PDFModel, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			if len(strings.Split(line, ",")) >= 2 {
				pwModel := stringToPDFModel(line)
				pwModelList = append(pwModelList, pwModel)
			}
		}
	}

	return pwModelList
}

func stringToPDFModel(str string) PDFModel {
	slice := strings.Split(str, ",")
	return PDFModel{Name: slice[0], PW: slice[1]}
}

func encryptPDFBy(name, pw string) {
	inpath := "./in/" + name + ".pdf"
	outPath := "./out/" + name + ".pdf"
	conf := model.NewAESConfiguration(pw, pw, 256)
	api.EncryptFile(inpath, outPath, conf)
}

type PDFModel struct {
	Name string
	PW   string
}
