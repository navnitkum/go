
package main

import (
   
    "fmt"
    "log"
   
	"encoding/json"
	"bytes"
	"crypto/tls"
	"net/http"
)

type PodDetails struct{
	
	Services []NamespaceDetails
	
 }

func main() {
    Namespacedetails := NamespaceDetails{}
	NamespaceList := []NamespaceDetails{}

	for{
	NamespaceList = Namespacedetails.GetNamespaceDetails()

	podDetails := PodDetails{}
	podDetails.Services = NamespaceList
   
   jsonData, err := json.Marshal(podDetails)
	if err != nil { log.Println(err) } 
		fmt.Println(string(jsonData))
		tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	    }
		
				clientIngestion := &http.Client{Transport: tr}
		
				var byteR = bytes.NewReader(jsonData)
				req1, _ := http.NewRequest("POST", "http://localhost:8082/api/v1/ingest/analytics", byteR)
		
				req1.Header.Set("content-type", "application/json")
				req1.Header.Set("topic-name", "ServicePodMetric")
				req1.Header.Set("x-auth-header", "abc")
				res, _ := clientIngestion.Do(req1)
				fmt.Println("***************Result from Ingestion API*******\n", res)
		
			
	}
}
