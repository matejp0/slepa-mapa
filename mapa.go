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
  "sort"
)

type Term struct {
  topic string
  value string
}

func main() {
  var location = flag.String("f", "pojmy.md", "Text file location")
  flag.Parse()

  terms := scanFile(location)
  //used := make([]Term, 0)
  used := make(map[Term]int64)

  rand.Seed(time.Now().UnixMilli())
  fmt.Print("\033[H\033[2J") // clear unix terminal
  fmt.Println("Počet pojmů:", len(terms))

  for len(terms) != len(used) {
    i := rand.Intn(len(terms))
    term := terms[i]
    var start int64
    if used[term] == 0 {
      start = time.Now().UnixMilli()
      fmt.Printf("[%d%%] %s: %s", 100*len(used)/len(terms), term.topic, term.value)
      fmt.Scanln()
      used[term] = time.Now().UnixMilli() - start
    } else {
      continue
    }
  }
  sort.SliceStable(terms, func(i, j int) bool{
    return used[terms[i]] > used[terms[j]]
  })

  fmt.Println("=============================")

  for i := 0; i < len(terms) / 5; i++ {
    fmt.Printf("%s: %s (%dms) \n", terms[i].topic, terms[i].value, used[terms[i]])
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
    if (arr[0][0] == '#') {continue}
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
