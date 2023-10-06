package main



import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)


func main(){

	type dola struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		Uf          string `json:"uf"`
		Ibge        string `json:"ibge"`
		Gia         string `json:"gia"`
		Ddd         string `json:"ddd"`
		Siafi       string `json:"siafi"`
	}

	var Digite_Cep string 
	
	fmt.Printf("Digite o cep EX: 01001000: ")
	fmt.Scan(&Digite_Cep)
	fmt.Println()


	req, err := http.Get("https://viacep.com.br/ws/" + Digite_Cep + "/json/")
	if err != nil{
		fmt.Println("ERRO", err)
	}

	

	res,_ := io.ReadAll(req.Body)

	var data dola

	json.Unmarshal(res, &data)
	

	fmt.Println("Cep:", data.Cep)
	fmt.Println("Rua:", data.Logradouro)
	fmt.Println("Complemento:", data.Complemento)
	fmt.Println("Bairro:", data.Bairro)
	fmt.Println("Localidade:", data.Localidade)
	fmt.Println("Estado:", data.Uf)
	fmt.Println("Ibge:", data.Ibge)
	fmt.Println("Gia:", data.Gia)
	fmt.Println("DDD:", data.Ddd)
	fmt.Println("Siafi:", data.Siafi)


	fmt.Println()
	fmt.Printf("Google maps: ")
	fmt.Println("https://www.google.com.br/maps/place/?q=" + Digite_Cep)
	fmt.Println()

}
