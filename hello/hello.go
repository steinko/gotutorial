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
     return greetingPrefix(language) + name
}
 
 func greetingPrefix(language string) (prefix string) {    
      switch language {
      	case spanish:
      	  prefix = spanishPrefix
        case french:
          prefix = frenchPrefix	 
          
          default:
           prefix = englishPrefix  
      }
      return
 }
   