package main

import (
  "bufio"
  "fmt"
  "flag"
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
  var location = flag.String("f", "pojmy.txt", "Text file location")
  flag.Parse()

  terms := scanFile(location)
  used := make([]Term, 0)
  
  rand.Seed(time.Now().UnixMilli())
  fmt.Print("\033[H\033[2J") // clear unix terminal

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

func scanFile(location *string) []Term {
  pojmy := make([]Term, 0)
  
  path, _ := os.Getwd()

  file, err := os.Open(fmt.Sprintf("%s/%s", path, *location))
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
