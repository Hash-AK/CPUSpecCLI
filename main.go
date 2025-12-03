package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

//go:embed cpus.json
var DB embed.FS

type CPUs struct {
	CPUs []CPU `json:"cpus"`
}
type CPU struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Brand              string   `json:"brand"`
	Generation         int      `json:"generation"`
	GenerationCodename string   `json:"generation_codename"`
	Series             string   `json:"series"`
	Specs              Specs    `json:"specs"`
	Features           []string `json:"features"`
	Overclockable      bool     `json:"overclockable"`
	ReleaseDate        string   `json:"release_date"`
}
type Specs struct {
	Cores                    int     `json:"cores"`
	Threads                  int     `json:"threads"`
	CacheL1KB                float64 `json:"cache_l1_kb"`
	CacheL2MB                float64 `json:"cache_l2_mb"`
	CacheL3MB                float64 `json:"cache_l3_mb"`
	BaseFrequencyGHz         float64 `json:"base_freq_ghz"`
	BoostFrequencyGHz        float64 `json:"boost_freq_ghz"`
	TDPWatts                 int     `json:"tdp_watts"`
	Socket                   string  `json:"socket"`
	Architecture             string  `json:"architecture"`
	IntegratedGPU            string  `json:"integrated_gpu"`
	MaximumSupportedMemoryGB int     `json:"max_mem_supported_gb"`
}

func dumpAllCpus(cpus CPUs) {
	for i := range cpus.CPUs {
		fmt.Printf("[#%d]\n", i)
		fmt.Printf("ID: %s\n", cpus.CPUs[i].ID)
		fmt.Printf("├─Name: %s\n", cpus.CPUs[i].Name)
		fmt.Printf("├─Brand: %s\n", cpus.CPUs[i].Brand)
		fmt.Printf("├─Generation: %d\n", cpus.CPUs[i].Generation)
		fmt.Printf("├─Generation's codename: %s\n", cpus.CPUs[i].GenerationCodename)
		fmt.Printf("├─Series: %s\n", cpus.CPUs[i].Series)
		fmt.Println("└┬─Specs:")
		fmt.Printf(" ├─Total cores #: %d\n", cpus.CPUs[i].Specs.Cores)
		fmt.Printf(" ├─Threads: %d\n", cpus.CPUs[i].Specs.Threads)
		fmt.Printf(" ├─Base frequency: %fGHz\n", cpus.CPUs[i].Specs.BaseFrequencyGHz)
		fmt.Printf(" ├─Boost frequency: %fGHz\n", cpus.CPUs[i].Specs.BoostFrequencyGHz)
		fmt.Printf(" ├─TDP: %d Watts\n", cpus.CPUs[i].Specs.TDPWatts)
		fmt.Printf(" ├─Socket: %s\n", cpus.CPUs[i].Specs.Socket)
		fmt.Printf(" ├─Architecture: %s\n", cpus.CPUs[i].Specs.Architecture)
		fmt.Printf(" ├─Integrated GPU: %s\n", cpus.CPUs[i].Specs.IntegratedGPU)
		fmt.Printf(" ├─Maximum supported memory: %dGB\n", cpus.CPUs[i].Specs.MaximumSupportedMemoryGB)
		fmt.Println("┌┘")
		fmt.Println("└┬─Features:")
		for f := range cpus.CPUs[i].Features {
			fmt.Printf(" ├─%s\n", cpus.CPUs[i].Features[f])
		}
		fmt.Println("┌┘")
		fmt.Printf("├─Overclockable?: %t\n", cpus.CPUs[i].Overclockable)
		fmt.Printf("└─Release date: %s\n", cpus.CPUs[i].ReleaseDate)
		fmt.Println()
	}
}
func CaseInsensitiveContains(s, substring string) bool {
	// Code taken from https://stackoverflow.com/questions/24836044/case-insensitive-string-search-in-golang
	s, substring = strings.ToLower(s), strings.ToLower(substring)
	return strings.Contains(s, substring)
}
func main() {
	idToSearch := pflag.String("id", "none", "CPU ID to search in database")
	searchTerm := pflag.String("search", "none", "Term to search in the CPU names.")
	dbJson, err := DB.ReadFile("cpus.json")
	if err != nil {
		panic(err)
	}
	var cpus CPUs
	json.Unmarshal([]byte(dbJson), &cpus)

	pflag.Parse()
	if *idToSearch != "none" {
		var idFound int
		var foundMatch bool
		for i := range cpus.CPUs {
			if cpus.CPUs[i].ID == *idToSearch {
				idFound = i
				foundMatch = true
			}
		}
		if foundMatch {
			fmt.Printf("Found a match for id %s : %s\n", *idToSearch, cpus.CPUs[idFound].Name)

		} else {
			fmt.Printf("Could not find any CPU in database with id %s\n", *idToSearch)
		}

	}
	if *searchTerm != "none" {
		var foundIDs []int

		for i := range cpus.CPUs {
			if CaseInsensitiveContains(cpus.CPUs[i].Name, *searchTerm) {
				foundIDs = append(foundIDs, i)
			}
		}
		if len(foundIDs) > 0 {
			fmt.Printf("Found %d matches:\n", len(foundIDs))
			for i := range foundIDs {
				fmt.Printf("Name: %s, ID: %s\n", cpus.CPUs[i].Name, cpus.CPUs[i].ID)
			}
		} else {
			fmt.Printf("Not matches found for term: %s\n", *searchTerm)
		}
	}
	//dumpAllCpus(cpus)
}
