/*
* In this file we store all the configuration for the codifications, as well
* as the functions that load the codifications table from a json file
*
* All the codifications tables are inside the codifications/ file. For example
* codifications/caesar.json
*/
package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "flag"
  "path/filepath"
  "strconv"
)

func loadArguments() (string, string, string, string) {
  /*
  * This function is used to parse all the arguments from the "os" module. It
  * also validates the arguments before using them.
  * Then he returns a list with the arguments
  *
  * Returns:
  *   - The reading file route
  *   - The writing file route 
  *   - The codification (default is caesar)
  */
  readingFile := flag.String("r", "", "The file to read from")
  writingFile := flag.String("w", "", "The file to write at")
  codificationTable := flag.String("c", "caesar", "The codification the program should use.")
  decrypt := flag.Bool("d", false, "Do you want to decrypt instead of encrypt?")
  flag.Parse()
  fmt.Printf("\n[PARSING]:         r %s | w  %s\n", *readingFile, *writingFile)
  fmt.Printf("[CODIFICATION]:    %s\n", *codificationTable)
  fmt.Printf("[DECRYPT]:         %v", *decrypt)
  // Create basic variables to then return them and check if there•
  // is a reading file and a writing file
  var readingFileRoute, writingFileRoute string
  if *readingFile == "" || *writingFile == "" {
    // One of the flags is missing so we can't parse them
    fmt.Println("[FATAL ERROR]:     Parse error, one of the flags is missing")
    os.Exit(1)
  } 
  // Convert from the relative path to the absolute path using the filepath module
  readingFileRoute, readingErr := filepath.Abs(*readingFile)
  writingFileRoute, writingErr := filepath.Abs(*writingFile)
  if readingErr == nil && writingErr == nil {
    // There actually were two paths, so we should inform the user
    fmt.Printf("\n[READING FROM]:    %s\n", readingFileRoute)
    fmt.Printf("[WRITTING TO]:     %s\n", writingFileRoute)
    // Return the complete routes
    return readingFileRoute, writingFileRoute, *codificationTable, strconv.FormatBool(*decrypt)
  } else {
    fmt.Println("[FATAL ERROR]:     One of the routes is not valid")
    os.Exit(1)
  } 
  return "", "", "", ""
}
    
 
func loadCipher (cipher string) map[string]string {
  /*
  * This function loads a cipher from the codifications/ folder to then return it.
  * The default cipher that is loaded is the caesar cipher.
  *
  * Parameters:
  *   cipher -> The cipher to be loaded [caesar by defdault]
  *
  * Returns:
  *   The loaded cipher (map[string]string)
  */
  cipherToLoad, pathErr := filepath.Abs(fmt.Sprintf("/codifications/%s.json", cipher))
  if pathErr != nil {
    fmt.Println("[FATAL ERROR]:      Can not get the absolute path of codification")
    os.Exit(1)
  }
  jsonCipher, err := ioutil.ReadFile(cipherToLoad)
  if err != nil {
    fmt.Println("[FATAL ERROR]:     Can not load the json cipher file")
    os.Exit(1)
  }
  // Inform the user about the loading cipher
  fmt.Printf("[CIPHER]:          %s\n", cipherToLoad)
  // Parse the json file to a cipher
  var jsonCodificationData map[string]string
  err = json.Unmarshal(jsonCipher, &jsonCodificationData)
  if err != nil {
    fmt.Println("[FATAL ERROR]:      Can not Unmarshal the json cipher file\n")
    os.Exit(1)
  }
  fmt.Println("[CIPHER]:          Loaded correctly")
  return jsonCodificationData
}

func loadFile(fileRoute string) []byte {
  fmt.Printf("[LOADING]:         File located at - %s\n", fileRoute)
  file, err := os.Open(fileRoute)
  if err != nil {
    fmt.Printf("[FATAL ERROR]:     Encrl was not able to open the file located at %s.\n", fileRoute)
    os.Exit(1)
  }
  fileContent, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Printf("[FATAL ERROR]:     Can not convert the file located at %s to a byte(s) array.\n", fileRoute)
    os.Exit(1)
  }
  return fileContent
}
