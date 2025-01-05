package main

import (
"bufio"
"fmt"
"log"
"os"
"net"
"strings"
)

func main(){
  scanner:= bufio.NewScanner(os.Stdin)
  fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

  for scanner.Scan(){
    checkDomain(scanner.Text())
  }

  if err:= scanner.Err(); err!=nil{
    log.Fatal("Error couldnt read from input: %v", err)
  }
}

func checkDomain(domain string){
  var hasMX, hasSPF, hasDMARC bool
  var spfRecord, dmarcRecord string

  mxRecords, err:= net.LookupMX(domain)
  
  if err!= nil{
    log.Printf("Error: %v",err)
  }
  if len(mxRecords) > 0 {
    hasMX = true
  }

  txtRecords, err:= net.LookupTXT(domain)
  if err != nil{
    log.Printf("Error %v", err)
  }

  for _, i := range txtRecords{
    if strings.HasPrefix(i, "v=spf1"){
      hasSPF = true
      spfRecord = i
      break
    }
  }

  dmarcRecords, err:= net.LookupTXT("_dmarc."+domain)
  if err != nil{
    log.Printf("Error %v", err)
  }

  for _, i := range dmarcRecords{
    if strings.HasPrefix(i, "v=DMARC1"){
    hasDMARC = true
    dmarcRecord = i
    break
    }
  }

  fmt.Printf("%v","%v","%v","%v","%v","%v",domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
