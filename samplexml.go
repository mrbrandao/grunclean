
import (
	"encoding/xml"
	"fmt"
	"grunclean/model"
	"io/ioutil"
	"net/http"
	"time"
)

/*XmlTest is a study case using xml instead of json.
It is here just as a sample, it prints Version and Project names using xml
*/
func XmlTest() {
	//Iniciando cliente e setando timeout
	cliente := &http.Client{
		Timeout: time.Second * 10,
	}

	//Fazendo a requisicao
	request, err := http.NewRequest("GET", url+"/api", nil)
	if err != nil {
		fmt.Printf("[main] Error on create request url", err.Error())
		return
	}

	//Efetuando o request
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Printf("[main] Error on Do request url", err.Error())
		return
	}
	defer resposta.Body.Close()

	//Contruindo o Body da Resposta
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Printf("[main] Error on build the resposta body", err.Error())
			return
		}
		//fmt.Println(string(corpo))
		post := model.ApiV{}
		err = xml.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Printf("[main] Error on Unmarshal xml", err.Error())
			return
		}
		fmt.Printf("This rundeck api Version is: %+v\r\n", post.Version)

		fmt.Println("Starting NewRequest...")
		request, _ := http.NewRequest("GET", url+"/api/"+post.Version+"/projects?authtoken="+token, nil)
		//EUREKA --> Recebendo as saidas em json
		//request.Header.Set("Accept", "application/json")
		resposta, _ = cliente.Do(request)
		defer resposta.Body.Close()
		corpo, _ = ioutil.ReadAll(resposta.Body)
		fmt.Println(string(corpo))
		posta := model.Project{}
		err = xml.Unmarshal(corpo, &posta)
		fmt.Printf("Listing the project: %+v\r\n", posta.Count)
		fmt.Printf("Listing the project: %+v\r\n", posta.Name[0])
	}
}
