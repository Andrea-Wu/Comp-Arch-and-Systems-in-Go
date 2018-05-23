package main

import (
    "fmt"
    "io/ioutil"
    //"unicode"
    "bufio"
    "strconv"
    "strings"
)

type node struct{
    next *node
    data int
}

func main(){
    //static
    /*
    var int_arr = [5]int{7, 10, 5, 10, 10}
    var char_arr = [5]byte{'i', 'i', 'i', 'i', 'i'}

    count := 5
    */
    str, err := ioutil.ReadFile("hehe.txt")

    if err != nil{
        fmt.Printf("cannot read from file\n")
        return
    }

    scanner := bufio.NewScanner(strings.NewReader(string(str)))

    int_arr := make([]int, 0)
    char_arr:= make([]string, 0)

    //assume there's the same number of integers as characters
    //I did not account for errorneous input
    var count int

    for scanner.Scan(){
        token := scanner.Text()
        if Int ,err := strconv.Atoi(token); err == nil{
            int_arr = append(int_arr, Int)
            count++
        }else{
            fmt.Printf("%s is not a number\n", token)
            char_arr = append(char_arr, token)
        }
    }

    for i := 0; i < count; i++{
        fmt.Printf("int at %d is %d\n", i, int_arr[i])
        fmt.Printf("char at %d is %s\n", i, char_arr[i])
    }

    var head *node = nil

    fmt.Printf("about to insert. count is %d\n", count)
  //  fmt.Printf("count is %d\n", count)
    for i := 0; i < count; i++{
        Char := char_arr[i]
        Int := int_arr[i]

        fmt.Printf("currently on %d iteration\n", i)

        if Char == "d"{
            //delete only if it exists, bro
            if head == nil{
                fmt.Printf("cannot del from empty list\n")
                continue
            }


        }else if Char == "i"{

            //this breaks it
            //var newNode *node

            newNode := &node{nil, Int}
           // newNode.data = Int

            if head == nil{
                head = newNode
                continue;
            }

            //insert into right place
                //case insert into front
                //case insert into middle
                //case insert into end

            //dynamic type declaration
            temp := head
            var prev *node

            isDuplicate := false
            for temp != nil && temp.data <= newNode.data {
                if temp.data == newNode.data{
                    fmt.Printf("can't insert duplicates\n")
                    prev = temp
                    temp = temp.next
                    isDuplicate = true
                    break
                }

                prev = temp
                temp = temp.next
            }

            //3 conditions that cause break

            if isDuplicate{
                continue
            }

            //temp is null (insert at end)
            if temp == nil{
                prev.next = newNode
                continue
            }

            //t.data > n.data
            if temp.data > newNode.data{

                //reset head
                if prev == nil{
                    head = newNode
                    head.next = temp
                }else{
                    prev.next = newNode
                    newNode.next = temp
                }
                continue
            }
        }else{
            fmt.Printf("can't do strings like that\n")
        }
    }
    fmt.Printf("about to print LL\n")
     printLL(head)
     
 }

func printLL(head *node) int{

        for temp := head; temp != nil; temp = temp.next{
            fmt.Printf("sorted: %d\n", temp.data)
        }
        return 0;

}    
     
