package testmod

import "fmt"

func Hi(name, lang string) (string, error) {
  switch lang {
  case "en": return fmt.Sprintf("Hi %s!", name), nil
  case "ru": return fmt.Sprintf("Привет %s!", name), nil
  default  : return "", errors.New("unknown language")
  }
}
