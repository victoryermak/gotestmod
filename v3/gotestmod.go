package gotestmod

import "fmt"

// Hi returns a friendly greeting
func Hi(name string, residence string) string {
  if residence =="UA" {
   return fmt.Sprintf("Hi, %s!!!", name)
  } else{
    return fmt.Sprintf("Fuck you, %s!!!", name)
  }
}
