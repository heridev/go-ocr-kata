package bankocr

import (
  "strings"
  "io/ioutil"
  "strconv"
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
    accountNumber += getMapValue(key)
    key = ""
    charNum = charNum+3
  }

  return accountNumber
}

func getMapValue(key string) string {
  mapValue := OCRMaps[key]
  if mapValue != "" {
    return mapValue
  }
  return "?"
}

func validateChecksum(numberString string) bool {
  parsedNumber := ""
  total := 0
  length := len(numberString) - 1
  for i, j:= length, 1; i >= 0; i, j = i - 1, j + 1 {
    parsedNumber = string(numberString[i])
    number, _ := strconv.Atoi(parsedNumber)
    total += number * j
  }
  return total % 11 == 0
}

func readFile(path string) string {
  fileBytes, _ := ioutil.ReadFile(path)
  byteStrings := string(fileBytes)
  return byteStrings
}

func check(err error){
  if err != nil{
    panic(err)
  }
}

func getAccountsFromFile(fromFile string, toFile string) {
  lines := getLinesFromFile()
  result:= ""
  accountNumber:=""
  for i:=0; i < len(lines) - 1; i += 4 {
    blockLines := lines[i:4 + i]
    ocrLines := strings.Join(blockLines, "\n")
    accountNumber = GetAccountNumber(ocrLines)
    status := getState(accountNumber)
    result += status
    result += "\n"
  }
  WriteOnFile(result, toFile)
}

func getState(accountNumber string) string {
  if (strings.Index(accountNumber, "?") != -1){
    return accountNumber + " ILL"
  }
  if !(validateChecksum(accountNumber)){
    alternatives, length := getAlternativeAccountNumbers(accountNumber)
    if length > 0 {
      if length == 1 {
        return alternatives[0]
      }
      if length > 1{
        alts := accountNumber + " AMB ["
        for _, val := range alternatives{
          if val != "" {
            alts += " " 
            alts += val
          }
        }
        alts +=" ]"
        return alts
      }
    }else{
      return accountNumber + " ERR"
    }
  }
  return accountNumber
}

func WriteOnFile(input string, path string){
  byteString := []byte(input)
  err := ioutil.WriteFile(path, byteString, 777)
  check(err)
}

func getLinesFromFile() []string{
  byteStrings := readFile("bankSource.txt")
  lines := strings.Split(byteStrings, "\n")
  totLines := len(lines) - 1
  return lines[0: totLines]
}

func getAlternativeAccountNumbers(accountNumber string) ([]string, int){
  altArray := make([]string, 100)
  key := ""
  alternatives := []byte{}
  j := 0
  for i:= 0; i<len(accountNumber); i++{
    key = string(accountNumber[i])
    alternatives = altMap[key]
    tmpAccountNumber := []byte(accountNumber)
    for _, alternative := range alternatives {
      tmpAccountNumber[i] = alternative
      if validateChecksum(string(tmpAccountNumber)) {
        altArray[j] = string(tmpAccountNumber)
        j++
      }
      tmpAccountNumber = []byte(accountNumber)
    }
  }
  return altArray, j
}

