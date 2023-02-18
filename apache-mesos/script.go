package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type myapps struct {
	Apps jsonData `json:"app"`
}

type jsonData struct {
	Id        string        `json:"id"`
	Instances int32         `json:"instance"`
	Container containerData `json:"container"`
}

type containerData struct {
	Docker myDocker `json:"docker"`
}

type myDocker struct {
	Image string `json:"image"`
}

type pod struct {
	ApiVersion string
	Kind       string
	Metadata   myMetaData
	Spec       mySpec
}

type myMetaData struct {
	Name      string
	Namespace string
}

type mySpec struct {
	Container []mycontainer
}

type mycontainer struct {
	Name  string
	Image string
}

func main() {

	res, err1 := http.Get("http://74.220.19.20:8080/v2/apps/docker-nginx")
	if err1 != nil {
		panic(err1)
	}

	defer res.Body.Close()

	databyte, _ := ioutil.ReadAll(res.Body)

	var myJsonData myapps

	err := json.Unmarshal(databyte, &myJsonData)
	if err != nil {
		panic(err)
	}

	myPodData := pod{
		ApiVersion: "v1",
		Kind:       "pod",
		Metadata: myMetaData{
			Name: myJsonData.Apps.Id,
		},
		Spec: mySpec{
			Container: []mycontainer{
				{Name: "", Image: myJsonData.Apps.Container.Docker.Image},
			},
		},
	}

	fmt.Printf(" The value of the struct is %v", myPodData)

	j, _ := json.Marshal(myPodData)

	fmt.Println(string(j))

}
