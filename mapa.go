package main

import (
  "bufio"
  "fmt"
  "strings"
  "log"
  "math/rand"
  "time"
  "os"
  "golang.org/x/exp/slices"
)

func main() {
  terms := scanFile()
  rand.Seed(time.Now().UnixMilli())
  used := make([]string, 0)
  fmt.Print("\033[H\033[2J") // clear terminal
  for true {
    if(len(terms) == len(used)) { 
      fmt.Println("Done")
      break
    }
    
    if term := terms[rand.Intn(len(terms))]; !slices.Contains(used, term) {
      used = append(used, term)
      percent := 100*len(used)/len(terms)
      fmt.Printf("[%d%%] %v", percent, term)
    } else {
      continue
    }
    fmt.Scanln()
  }

}

func scanFile() []string {
  pojmy := make([]string, 0)
  file, err := os.Open("/home/matt/Code/slepa-mapa/pojmy.txt")
  if err != nil { log.Fatal(err) }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    arr := strings.Split(scanner.Text(), ": ")
    category := arr[0]
    names := strings.Split(arr[1], ", ")
    for _, name := range names {
      pojmy = append(pojmy, fmt.Sprintf("%s: %s", category, name))
    }
  }
  if err := scanner.Err(); err != nil { log.Fatal(err) }
  return pojmy
}
