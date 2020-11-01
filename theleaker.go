/*
 * --------------------------------------------------------------------------------
 * "SUSHIWARE LICENSE" (Revision 01):
 * As long as you retain this notice, you can do whatever you want with this code.
 * If we meet someday around the universe, and you think this code was useful,
 * if you want, you can pay me a sushi round in return.
 * Ewerson Guimaraes a.k.a Crash
 * --------------------------------------------------------------------------------
 */

package main

import (
	"archive/tar"
	"archive/zip"
	//"bufio"
	//"compress/gzip"
	b64 "encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const TltVersion = "0.1"
//#######################################################################
func WriteExcel(argCells string,payload string, outPut string, zipfile *bool)  {
	if argCells == "" {
		fmt.Println("The flag -xlsx was provided, it requires at least one cell to be written. For example: -c A1")
		os.Exit(0)
	}else {
		outPut = outPut+".xlsx"
		xlsxWrite := excelize.NewFile()
		cells := strings.Split(argCells, ",")
		//fmt.Println(cells[0]) // debug
		//fmt.Println("Array size:",len(cells)) //debug
		for i := 0; i < len(cells); i++ {
			fmt.Println("Writing the cell:", cells[i])
			xlsxWrite.SetCellValue("Sheet1", cells[i], payload)
			err := xlsxWrite.SaveAs(outPut)
			if err != nil {
				fmt.Println(err)
			}  else {
				fmt.Println("The Excel file " + outPut + " has been created, containing the payload:" + payload)
			}
		}
	}
	//if *zipfile {
	//	Writezip(outPut)
	//}
}
//#######################################################################
func WritePdf(payload string, outPut string, zipfile *bool ) {
	outPut = outPut+".pdf"
	createPdf := gofpdf.New("P", "mm", "A4", "")
	createPdf.AddPage()
	createPdf.SetFont("Arial", "", 11)
// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	createPdf.CellFormat(190, 7, payload, "0", 0, "LM", false, 0, "")
	// err := createPdf.OutputFileAndClose(outPut)
	fileFinal := createPdf.OutputFileAndClose(outPut)
	if fileFinal != nil {
		fmt.Println(fileFinal)
	}  else {

		fmt.Println("The PDF file " + outPut + " has been created, containing the payload:" + payload)
	}

}
//#######################################################################
func WriteZip (payload string,outPut string, ) {
	outFile, err := os.Create(outPut + ".zip")
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()
	w := zip.NewWriter(outFile)

	var files = []struct {
		Name, Body string
	}{
		{ strconv.Itoa(rand.Int()),payload}, //Add a random  numeric file name to zip and the payload
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The ZIP file " + outPut + ".zip has been created, containing the payload:" + payload)
	}
}
//#######################################################################
func WriteRaw (payload string, outPut string) {
	rw := []byte(payload)
	err := ioutil.WriteFile(outPut, rw,700)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The RAW file " + outPut + " has been created, containing the payload:" + payload)
	}
}
//#######################################################################
func WriteTar (payload string, outPut string) {
	tarPath := outPut +".tar"
	files := map[string]string{
		strconv.Itoa(rand.Int()): payload,
		}
	tarWrite := func(data map[string]string) error {
		tarFile, err := os.Create(tarPath)
		if err != nil {
			return err
		}
		defer tarFile.Close()

		tw := tar.NewWriter(tarFile)
		defer tw.Close()

		for name, content := range data {
			hdr := &tar.Header{
				Name: name,
				Mode: 0600,
				Size: int64(len(content)),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			if _, err := tw.Write([]byte(content)); err != nil {
				return err
			}
		}
		return nil
	}

	if err := tarWrite(files); err != nil {
		log.Fatal(err)
	} else {
			fmt.Println("The TAR file " + outPut + ".tar has been created, containing the payload:" + payload)
		}
	}

//#######################################################################
//#######################################################################
func WriteGz (payload string, outPut string) {

}
//#######################################################################

func main() {
	var argCells string
	flag.StringVar(&argCells, "c", "", "Single cells where the payload will be placed. Use a comma as a separator. Ex: A1,H4,D20")
	var payload string
	flag.StringVar(&payload, "p", "", "Any payload to be written in the file.")
	var encodeType string
	flag.StringVar(&encodeType, "encode", "", "Available encodes: base64 hex bin.\nTo run more than one simultaneously, use comma as a separator. Ex. url,hex\nThe final payload encoding follows the provided order.")
	var outPut string
	flag.StringVar(&outPut, "o", "tlt", "File name output. The extension will be automatically added according to the selected file type.")
	version := flag.Bool("v", false, "Prints the current version and exit.")
	zipfile := flag.Bool("zip", false, "Generates a ZIP file containing a random name zipped file and the provided test string.")
	excel := flag.Bool("xlsx", false, "Generates an Excel file containing the test string. The flag -c is required.")
	pdf := flag.Bool("pdf", false, "Generates a PDF file containing the test string.")
	allFiles := flag.Bool("all", false, "Generates all files at once.")
	rawfile := flag.Bool("raw", false, "Generates an RAW file containing the test string.")
	tarfile := flag.Bool("tar", false, "Generates a TAR file containing a random name zipped file and the provided test string")
	gzfile := flag.Bool("gz", false, "Generates a GZ file containing a random name zipped file and the provided test string")
	flag.Parse()


	if *version {
		fmt.Println(TltVersion)
		fmt.Println("Under the SushiWare license.")
		os.Exit(0)
	}

	if encodeType == "" {
	} else {
		encodeSplit := strings.Split(encodeType, ",")
		for i := 0; i < len(encodeSplit); i++ {

			if encodeSplit[i] == "base64" {
				payload = b64.StdEncoding.EncodeToString([]byte(payload))
				fmt.Println("Encode:", encodeSplit[i]) // debug
				fmt.Println("Payload:", payload)       //
			} else if encodeSplit[i] == "hex" {
				payload = hex.EncodeToString([]byte(payload))
				fmt.Println("Encode:", encodeSplit[i]) // debug
				fmt.Println("Payload:", payload)       //
			} else if encodeSplit[i] == "bin" {
				payloadsize := len(payload)
				for _, bina := range payload {
					payload = fmt.Sprintf("%s%b", payload, bina)
				}
				payload = payload[payloadsize:]
				fmt.Println("Encode:", encodeSplit[i])   // debug
				fmt.Println("Payload Encoded:", payload) //
			}

		}
	}
		if *excel {
			WriteExcel(argCells, payload, outPut, zipfile)
		}
		if *pdf {
			WritePdf(payload, outPut, zipfile)
		}
		if *zipfile {
			WriteZip(payload, outPut)
		}
		if *rawfile {
			WriteRaw(payload, outPut)
		}
		if *tarfile {
			WriteTar(payload, outPut)
		}
		if *gzfile {
			WriteGz(payload, outPut)
		}
		if *allFiles {
			WriteExcel(argCells,payload,outPut,zipfile)
			WritePdf(payload,outPut, zipfile)
			WriteRaw(payload, outPut)
			WriteZip(payload, outPut)
			WriteTar(payload, outPut)
			WriteGz(payload, outPut)
		}


	}
