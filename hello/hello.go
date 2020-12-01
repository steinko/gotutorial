package main

const englishPrefix = "Hello "

func Hello(name string , language string) string {
     if name == "" {
        name = "World"
      }
      
      if language == "Spanish" { 
          return "Hola " + name
      }
      
     return englishPrefix + name
}