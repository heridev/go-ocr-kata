package bankocr

import(
  "fmt"
  //"io/ioutil"
  "strings"
)

var altMap = map[string][]byte{
  "0" : { '8' },
  "1" : { '7' },
  "3" : { '9' },
  "5" : { '6', '9' },
  "6" : { '5','8' },
  "7" : { '1' },
  "8" : { '0', '6', '9' },
  "9" : { '8', '3' },
}

var OCRMaps = map[string]string{
  " _ | ||_|": "0",
  "     |  |": "1",
  " _  _||_ ": "2",
  " _  _| _|": "3",
  "   |_|  |": "4",
  " _ |_  _|": "5",
  " _ |_ |_|": "6",
  " _   |  |": "7",
  " _ |_||_|": "8",
  " _ |_| _|": "9",
}

func ExampleStringReplace() {
  accountNumber := "490067715"

  accountByteNumber := []byte(accountNumber)
  accountByteNumber[2] = '9'
  fmt.Println(string(accountByteNumber))
  // Output:
  // 499067715
}

func ExampleGetAlternatives() {
  accountNumber := "888888888"
  //alternatives := []string{}
  alternatives, _ := getAlternativeAccountNumbers(accountNumber)

  //fmt.Println(status)
  for _, alter := range alternatives {
    fmt.Println(alter)
  }

  // Output:
  // 888886888
  // 888888988
  // 888888880
}

func ExampleChecksum(){
  invalidNumber := "111111111"
  validNumber := "123456789"
  fmt.Println(validateChecksum(invalidNumber))
  fmt.Println(validateChecksum(validNumber))

  // Output:
  // false
  // true 
}

func ExampleParseLinesIntoNumberCase1() {
  ocrnumbers := "    _  _     _  _  _  _  _ \n" +
                "  | _| _||_||_ |_   ||_||_|\n" +
                "  ||_  _|  | _||_|  ||_| _|\n"
  fmt.Println(GetAccountNumber(ocrnumbers))
  // Output:
  // 123456789
}

func ExampleReadLineOneFromFileTest() {
  lines := getLinesFromFile()

  for i:=0; i < len(lines) - 1; i += 4 {
    blockLines := lines[i:4 + i]
    ocrLines := strings.Join(blockLines, "\n")
    fmt.Println(GetAccountNumber(ocrLines))
  }

  // Output:
  // 123456789
  // 123456789
  // 123456789
  // 111111111
  // 1?1?1????
}

func ExampleConvertOCRIntoNumbers() {

  getAccountsFromFile("bankSource.txt", "output.txt")
  byteStrings := readFile("output.txt")
  fmt.Println(byteStrings)
  //lines := strings.Split(byteStrings, "\n")

  // Output:
  // 123456789
  // 123456789
  // 123456789
  // 711111111
  // 1?1?1???? ILL
  //
}

func ExampleWriteOnFile() {
  WriteOnFile("test", "output.txt")
  byteStrings := readFile("output.txt")
  lines := strings.Split(byteStrings, "\n")
  fmt.Println(lines[0])

  // Output:
  // test
}

func ExampleGetLineNumbers() {
  lines := getLinesFromFile()
  totLines := len(lines)
  fmt.Println(totLines)

  // Output:
  // 20
}

func ExampleCheckNumbersOnFile() {
  //saveNumbersIntoFile()
  // Output:
  // 
  //  ERR
  //  ILL
}

func ExampleCharactersPerLine() {
  lines := getLinesFromFile()
  fmt.Println(len(lines[0]))
  totLines := len(lines) - 1
  fmt.Println(len(lines[totLines]))

  // Output:
  // 27
  // 0
}

func ExampleGetState(){
  accountNumber := "888888888"
  state := getState(accountNumber)
  fmt.Println(state)

  // Output:
  // 888888888 AMB [ 888886888 888888988 888888880 ]

}

