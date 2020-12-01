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
     
      prefix := englishPrefix
      switch language {
      	case spanish:
      	  prefix = spanishPrefix
        case french:
          prefix = frenchPrefix	    
      }
      
      
      if language == spanish { 
          return spanishPrefix + name
      }
      
       if language == french { 
          return frenchPrefix + name
      }
      
     return prefix + name
}