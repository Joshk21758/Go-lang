package main

import (

  "fmt"
  "strconv"
  "os"
  "stdin"
  "bufio"

  )

func main(){
  scanner := bufio.NewScanner(os.stdin)
  println("Enter your weight: ")
  scanner.Scan()
  wght := strconv.ParseInt(scanner.Text(), 10, 20)
  //Decision
  if wght < 30 && wght < 45 {
    println("You are not built")
    
    }

  else if wght > 45 {
    println("You are built :) ")
    
    }

  else {
    println("What is your weight? ")
    
    }

  }