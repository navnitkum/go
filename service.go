package main
import (
    "bufio"
  
    "strings"
    
	"os/exec"
)

  type Services struct {
	   Namespace string
	   ServiceName string
  }



type NamespaceDetails struct {
   
 Namespace string
 ServiceName string

   PodList map[string][]string
  
  }
  
  func (n *NamespaceDetails) GetNamespaceDetails() []NamespaceDetails {
	var npvarArray []string
	var services []Services
   var podlist []NamespaceDetails
	


	out1, _ := exec.Command("sh", "-c", "kubectl get namespaces").Output()
	inputstr1 := string(out1)
	scanner := bufio.NewScanner(strings.NewReader(inputstr1))
    for scanner.Scan() {
			tempStr := scanner.Text()
			if !strings.Contains(tempStr, "NAME") {
				
				 tempArr := strings.Fields(tempStr)
				 npvarArray = append(npvarArray, strings.TrimSpace(tempArr[0]))
				 
			 }
			
		}
		
	for _,v3 := range npvarArray{
          serviceObj := Services{}
			
		out2, _ := exec.Command("sh", "-c", "kubectl get svc -n " + v3 +"").Output()
		inputstr2 := string(out2)
		scanner = bufio.NewScanner(strings.NewReader(inputstr2))
    	for scanner.Scan() {
			tempStr := scanner.Text()
			if !strings.Contains(tempStr, "NAME") {
				
				 tempArr := strings.Fields(tempStr)
				 serviceObj.Namespace = v3
				 serviceObj.ServiceName = strings.TrimSpace(tempArr[0])
				
				 services = append(services,serviceObj)
				
			}
		}
	
	

	
		
	}
	for _,i := range services{
		
		namespace := NamespaceDetails{}

		namespace.Namespace = i.Namespace
		namespace.ServiceName = i.ServiceName

		namespace.PodList = make(map[string][]string)
	 
		 kubectlCommand := "kubectl get pods -n  " + i.Namespace + " -l $(kubectl describe svc " + i.ServiceName + " -n " + i.Namespace + " | grep -e Selector |awk '{print $2}')"
		 out3, _ := exec.Command("sh", "-c", kubectlCommand).Output()
       
		 inputstr3 := string(out3)
		 

		 scanner = bufio.NewScanner(strings.NewReader(inputstr3))
		 var PodArray []string
    	for scanner.Scan() {
			tempStr := scanner.Text()
			if !strings.Contains(tempStr, "NAME") {
				
				 tempArr3 := strings.Fields(tempStr)

				 PodArray = append(PodArray,(strings.TrimSpace(tempArr3[0])))
			}
		}
		 namespace.PodList["Pods"]= PodArray
		 podlist = append(podlist,namespace)
	}
	return	podlist	
			
 }