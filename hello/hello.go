package main

const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const spanish = "Spanish"

func Hello(name string , language string) string {
     if name == "" {
        name = "World"
      }
      
      if language == spanish { 
          return spanishPrefix + name
      }
      
     return englishPrefix + name
}