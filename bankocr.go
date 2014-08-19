package bankocr

import (
  "strings"
  "io/ioutil"
)


func GetAccountNumber(ocr string) string {
  lines := strings.Split(ocr, "\n")[0:3]

  charNum := 0
  var key string
  accountNumber := ""
  for i:=0; i<9 ; i++{
    for _, val := range lines{
      key = key+val[charNum : charNum+3]
    }
    accountNumber += OCRMaps[key]
    key = ""
    charNum = charNum+3
  }

  return accountNumber
}

func getLinesFromFile() []string{
  fileBytes, _ := ioutil.ReadFile("bankSource.txt")
  byteStrings := string(fileBytes)
  lines := strings.Split(byteStrings, "\n")
  totLines := len(lines) - 1
  return lines[0: totLines]
}

