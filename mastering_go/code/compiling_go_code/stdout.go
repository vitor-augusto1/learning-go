package main

import (
  "io"
  "os"
)

// This code will use the io package instead of the fmt package.
// The os package is used for reading the command-line arguments
// of the program and for accessing os.Stdout.

func main() {
  myString := ""
  arguments := os.Args
  if len(arguments) == 1 {
    myString = "Please give me one argument!"
  } else {
    myString = arguments[1]
  }

  // Using the os.Stdout to print in the UNIX standard output
  // We are using the os.Stdout because it implements the io.Writer interface
  io.WriteString(os.Stdout, myString)
  io.WriteString(os.Stdout, "/n")
}
