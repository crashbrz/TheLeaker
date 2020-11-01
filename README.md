![License](https://img.shields.io/badge/license-sushiware-red)
![Issues open](https://img.shields.io/github/issues/crashbrz/TheLeaker)
![GitHub pull requests](https://img.shields.io/github/issues-pr-raw/crashbrz/TheLeaker)
![GitHub closed issues](https://img.shields.io/github/issues-closed-raw/crashbrz/TheLeaker)
![GitHub last commit](https://img.shields.io/github/last-commit/crashbrz/TheLeaker)

# TheLeaker #
The Leaker is a portable multi-format command line file crafter. TheLeaker allows us to create different files with embed payloads. The main goal of this tool is to test if the DLP solution in place is configured correctly.  

After setting up it properly, TheLeaker supports XSLX, PDF, Raw(txt), ZIP, TAR, and GZ file formats.
Also, TheLeaker can encode the payloads. Right now, hex, bin, and base64 encodings are supported. 

### Installation ###
Download the latest release and unpack it in the desired location. Remember to install GoLang in case you want to run from the source.
TheLeaker uses the Excelize and GofPdf libraries. 

Check the following links for more information:
[https://github.com/360EntSecGroup-Skylar/excelize/](https://github.com/360EntSecGroup-Skylar/excelize/)
[https://github.com/jung-kurt/gofpdf/](https://github.com/jung-kurt/gofpdf/) 

### License ###
TheLeaker is licensed under the SushiWare license. Check [docs/license.txt](docs/license.txt) for more information.

### Usage/Help ###
Please refer to the output of -h for usage information and general help. Also, you can contact me on `##spoonfed@freenode.org` (two #)

```
Usage of theleaker:
  -all
        Generates all files at once.
  -c string
        Single cells where the payload will be placed. Use a comma as a separator. Ex: A1,H4,D20
  -encode string
        Available encodes: base64 hex bin.
        To run more than one simultaneously, use comma as a separator. Ex. url,hex
        The final payload encoding follows the provided order.
  -gz
        Generates a GZ file containing a random name zipped file and the provided test string
  -o string
        File name output. The extension will be automatically added according to the selected file type. (default "tlt")
  -p string
        Any payload to be written in the file.
  -pdf
        Generates a PDF file containing the test string.
  -raw
        Generates an RAW file containing the test string.
  -tar
        Generates a TAR file containing a random name zipped file and the provided test string
  -v    Prints the current version and exit.
  -xlsx
        Generates an Excel file containing the test string. The flag -c is required.
  -zip
        Generates a ZIP file containing a random name zipped file and the provided test string.

  ```
