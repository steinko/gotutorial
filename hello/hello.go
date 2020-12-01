package main

const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjore "

const spanish = "Spanish"
const french = "French"

func Hello(name string , language string) string {
     if name == "" {
        name = "World"
      }
      
      if language == spanish { 
          return spanishPrefix + name
      }
      
       if language == french { 
          return frenchPrefix + name
      }
      
     return englishPrefix + name
}