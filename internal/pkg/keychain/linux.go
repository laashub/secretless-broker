// +build linux

package keychain

import (
  "fmt"
)

func GetGenericPassword(service, account string) (string, error) {
  return "", fmt.Errorf("No keychain provider for Linux (yet)")
}