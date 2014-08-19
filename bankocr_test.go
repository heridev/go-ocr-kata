package bankocr

import(
  "fmt"
  //"io/ioutil"
  "strings"
)

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

func ExampleParseLinesIntoNumberCase1() {
  ocrnumbers := "    _  _     _  _  _  _  _ \n" +
                "  | _| _||_||_ |_   ||_||_|\n" +
                "  ||_  _|  | _||_|  ||_| _|\n"
  fmt.Println(GetAccountNumber(ocrnumbers))
  // Output:
  // 123456789
}

//func ExampleValidate

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
}

func ExampleGetLineNumbers() {
  lines := getLinesFromFile()
  //fmt.Println(lines[11])
  totLines := len(lines)
  fmt.Println(totLines)

  // Output:
  // 12
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

func printValue(value string) string {
  showValue := value
  showValue += "============="
  return showValue
}

