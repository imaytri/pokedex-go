package main

import "fmt"
import "os"
import "strings"
import "net/http"
import "encoding/json"
import "strconv"

var poke = new(pokemon)
var it = new(item)
var test int

func main(){
	var q string
	var q_url string
	var stat int

	for i:=1; i<len(os.Args); i++ {
		q = q + os.Args[i] + " "
	}
	q_url = changeStrFormat(q, true)
	if q_url=="" {
		fmt.Println("Please fill Pokemon Species / Item parameter")
		os.Exit(1)
	}
	if len(os.Args)==2 {
		_, err := strconv.ParseInt(os.Args[1], 10, 0)
		if err==nil {
			fmt.Println("Pokemon Species or Item not found.")
			os.Exit(1)
		}
	}

	// Process
	stat, _ = callApi("https://pokeapi.co/api/v2/pokemon/" + q_url, poke)
	if (stat != 404){
		printResult("pokemon")
	} else {
		stat, _ = callApi("https://pokeapi.co/api/v2/item/" + q_url, it)
		if (stat != 404){
			printResult("item")		
		} else {
			printResult("")
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

func callApi(url string, target interface{}) (int, error) {
	response, err := http.Get(url)
	if err != nil {
        fmt.Println(err)
    }
	defer response.Body.Close()
	return response.StatusCode, json.NewDecoder(response.Body).Decode(&target)
}

func printResult(opt string) {
	switch opt {
	case "pokemon" :
		fmt.Printf("\n#%d - %s \n", poke.ID, changeStrFormat(poke.Name, false))
		fmt.Printf("Type : ")
		for i:=0; i<len(poke.Types); i++ {
			fmt.Print(changeStrFormat(poke.Types[i].Type.Name, false))
			if (i < len(poke.Types)-1){
				fmt.Print(" - ")
			} else {
				fmt.Print("\n\nAbilities:\n")
			} 
		}

		for i:=0; i<len(poke.Abilities); i++ {
			fmt.Printf("%d. %s", i+1, changeStrFormat(poke.Abilities[i].Ability.Name, false))
			var url string = poke.Abilities[i].Ability.URL
			abi := new(ability)
			callApi(url, abi)
			for j:=0; j<len(abi.EffectEntries); j++ {
				fmt.Printf("\n%s", abi.EffectEntries[j].Effect)
			}

			if (i<len(poke.Abilities)-1){
				fmt.Print("\n\n")
			}
		}
	case "item": 
		fmt.Printf("\nItem : %s\nCost : %d\n", changeStrFormat(it.Name, false), it.Cost)
		for i:=0; i<len(it.EffectEntries); i++ {
			fmt.Printf("Entries : %s ", it.EffectEntries[i].Effect)
		}
	default:
		fmt.Println("Pokemon Species or Item not found.")
	}
}