package main

func forcontinue () int {
     sum := 1
     for ; sum < 100 ; {
          sum += sum
     }
     return sum
}