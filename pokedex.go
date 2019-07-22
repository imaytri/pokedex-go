package main

import "fmt"
import "os"
import "strings"
import "net/http"
import "encoding/json"

func main(){
	var q string
	var q_url string
	var err error
	// var output string
	poke := new(pokemon)
	for i:=1; i<len(os.Args); i++ {
		q = q + os.Args[i] + " "
	}

	q_url = changeStrFormat(q, true)

	// call api
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + q_url)
	if err != nil {
        fmt.Println(err)
    }
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&poke)
	fmt.Printf("\n#%d - %s \n", poke.ID, changeStrFormat(poke.Name, false))
	fmt.Printf("Type : ")
	for i:=0; i<len(poke.Types); i++ {
		fmt.Print(changeStrFormat(poke.Types[i].Type.Name, false))
		if(i < len(poke.Types)-1){
			fmt.Print(" - ")
		} else {
			fmt.Print("\n\nAbilities:\n")
		} 
	}

	for i:=0; i<len(poke.Abilities); i++ {
		fmt.Printf("%d. %s", i+1, poke.Abilities[i].Ability.Name)
		var url string = poke.Abilities[i].Ability.URL
		abi := new(ability)
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		err = json.NewDecoder(response.Body).Decode(&abi)
		for j:=0; j<len(abi.EffectEntries); j++ {
			fmt.Printf("\n%s", abi.EffectEntries[j].Effect)
		}

		if(i<len(poke.Abilities)-1){
			fmt.Print("\n\n")
		}
	}
	
}

func changeStrFormat(str string, is_to_url bool) string {
	var res string 
	if is_to_url {
		res = strings.ReplaceAll(strings.TrimSpace(strings.ToLower(str)), " ", "-")
	} else {
		res = strings.Title(strings.ReplaceAll(str, "-", " "))
	}
	return res
}