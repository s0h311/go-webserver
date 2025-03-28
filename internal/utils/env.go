package utils

import (
  "fmt"
  "os"
)

func GetEnv(name string) string {
  value, exists := os.LookupEnv(name)

  if !exists {
    PrintlnColorful(Red, fmt.Sprintf("Env %s not found", name))
    os.Exit(1)
  }

  return value
}
