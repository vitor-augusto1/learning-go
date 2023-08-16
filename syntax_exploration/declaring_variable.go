package main

import "fmt"

func main() {
  var smsSendindLimit int = 100
  var costPerSMS float64 = 0.5
  var hasPermission bool = true
  var username string = "Vitor"

  fmt.Printf(
    "%v %f %v %q\n",
    smsSendindLimit,
    costPerSMS,
    hasPermission,
    username,
  )
}
