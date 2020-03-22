package main


import (
    "bufio"
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)


type QA struct {
    Question string
    Answare string
}


func main() {
    
    fmt.Println("Question and Answare game! ")

    problems := flag.String("problems", "problems.cvs", "cvs files with problems and answers")
    //timer := flag.Int("timer", 30, "time limit in second")
    
    flag.Parse()

    csvFile, _ := os.Open(*problems)
    fileReader := csv.NewReader(bufio.NewReader(csvFile))
    var QAs []QA 
    for {
        line, err :=  fileReader.Read()
        if err == io.EOF {
            break;
        }else if err != nil {
            log.Fatal(err)
        }
        QAs = append(QAs, QA{
            Question : line[0],
            Answare : line[1],
            
        })
    
    }
    inputReader := bufio.NewReader(os.Stdin)
    for  _ ,  qa := range QAs {
        fmt.Printf(" %s ? ", qa.Question)
        text, _ := inputReader.ReadString('\n')
        text = strings.Replace(text, "\n", "",-1)
        fmt.Println("TEXT : ", text , " answare : ", qa.Answare)
        if strings.Compare(qa.Answare, text) == 0 {
            fmt.Println("Right answare!")
        }
    }

}
