/*
* This file contains all the cryptographi operations
* to work with ENCRL, for example the encrypt function.
* */
package main

import (
  "fmt"
)

func encrypt(cipher map[string]string, file []byte) []byte{
  /*
  * This function is used to codificate a file using a given cipher
  *
  * Parameters:
  *   cipher -> The cipher to use when codificating (see more at ~/codiffications/ )
  *   file -> An array of bytes with the content of the fyle to encrypt
  *
  * Returns:
  *   An array of bytes that represent the new content of the file
  * 
  * */
  var charsCount uint
  var modifiedFile []byte
  for _, letter := range file {
    charsCount++
    value, exists := cipher[string(letter)]
    if exists {
      modifiedFile = append(modifiedFile, value[0])
    } else {
      // If the character is not located in the codification, set the new value 
      // to the old letter
      modifiedFile = append(modifiedFile, letter)
    }
  }
  fmt.Printf("[ENCRYPTION]:      Total modified characters %d.\n", charsCount)
  fmt.Printf("[ENCRYPTION]:      Byte array with the encryption loaded correctly.")
  return modifiedFile
}