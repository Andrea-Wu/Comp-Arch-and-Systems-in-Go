package main
            //corner case: i is len(uncompressed_string) -1

import(
    "fmt"
//    "strings"
//    "bufio"
    "os"
    "bytes"
    "strconv"
)

func main(){
    //command line argument

    if len(os.Args) != 2{
        fmt.Printf("invalid number of arguments\n")
        return
    }    

    uncompressed_string := os.Args[1]
    var compressed_string bytes.Buffer
    count := 1
    j := 0
    i := 1

    length := len(uncompressed_string)
    for i < length {
        if uncompressed_string[i] == uncompressed_string[j] {
            count++
        }else{
            //concat to string yo

            compressed_string.WriteString(string(uncompressed_string[j]))
            //strconv.Itoa converts Int to String
            compressed_string.WriteString(strconv.Itoa(count))

           count = 1 
        }

        j++
        i++
    }

    compressed_string.WriteString(string(uncompressed_string[length -1]))
    compressed_string.WriteString(strconv.Itoa(count))
    count = 0
 

    fmt.Printf("%s\n", compressed_string.String())
}
