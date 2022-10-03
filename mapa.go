package main

import (
  "bufio"
  "fmt"
  "strings"
  "log"
  "math/rand"
  "time"
  "os"
)

type Term struct {
  topic string
  value string
}

func main() {
  terms := scanFile()
  rand.Seed(time.Now().UnixMilli())
  used := make([]Term, 0)
  fmt.Print("\033[H\033[2J") // clear terminal

  for len(terms) != len(used) {
    if term := terms[rand.Intn(len(terms))]; !contains(used, term) {
      used = append(used, term)
      fmt.Printf("[%d%%] %s: %s", 100*len(used)/len(terms), term.topic, term.value)
    } else {
      continue
    }

    fmt.Scanln()
  }

}

func scanFile() []Term {
  pojmy := make([]Term, 0)
  file, err := os.Open("/home/matt/Code/slepa-mapa/pojmy.txt")
  if err != nil { log.Fatal(err) }
  defer file.Close() 
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    arr := strings.Split(scanner.Text(), ": ")
    names := strings.Split(arr[1], ", ")

    for _, name := range names {
      pojmy = append(pojmy, Term{
        topic: arr[0],
        value: name,
      })
    }
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }
  return pojmy
}

func contains(list []Term, term Term) bool {
  for _, v := range list {
    if v == term {
      return true
    }
  }
  return false
}
