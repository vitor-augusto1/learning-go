package main

import (
  "fmt"
  "log"
  "log/syslog"
  "os"
  "path/filepath"
)

func main() {
  programName := filepath.Base(os.Args[0])
  sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

  if err != nil {
    log.Fatal(err)
  } else {
    log.SetOutput(syslog)
  }
  log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
}
