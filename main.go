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
	CacheMB                  float64 `json:"cache_mb"`
	BaseFrequencyGHz         float64 `json:"base_freq_ghz"`
	BoostFrequencyGHz        float64 `json:"boost_freq_ghz"`
	TDPWatts                 int     `json:"tdp_watts"`
	Socket                   string  `json:"socket"`
	Architecture             string  `json:"architecture"`
	IntegratedGPU            string  `json:"integrated_gpu"`
	MaximumSupportedMemoryGB int     `json:"max_mem_supported_gb"`
}

func dumpID(id int, cpus CPUs) {
	fmt.Printf("ID: %s\n", cpus.CPUs[id].ID)
	fmt.Printf("├─Name: %s\n", cpus.CPUs[id].Name)
	fmt.Printf("├─Brand: %s\n", cpus.CPUs[id].Brand)
	fmt.Printf("├─Generation: %d\n", cpus.CPUs[id].Generation)
	fmt.Printf("├─Generation's codename: %s\n", cpus.CPUs[id].GenerationCodename)
	fmt.Printf("├─Series: %s\n", cpus.CPUs[id].Series)
	fmt.Println("│")
	fmt.Println("├──Specs:")
	fmt.Printf("│  ├─Total cores #: %d\n", cpus.CPUs[id].Specs.Cores)
	fmt.Printf("│  ├─Threads: %d\n", cpus.CPUs[id].Specs.Threads)
	fmt.Printf("│  ├─Cache: %.2fMB\n", cpus.CPUs[id].Specs.CacheMB)
	fmt.Printf("│  ├─Base frequency: %.2fGHz\n", cpus.CPUs[id].Specs.BaseFrequencyGHz)
	fmt.Printf("│  ├─Boost frequency: %.2fGHz\n", cpus.CPUs[id].Specs.BoostFrequencyGHz)
	fmt.Printf("│  ├─TDP: %d Watts\n", cpus.CPUs[id].Specs.TDPWatts)
	fmt.Printf("│  ├─Socket: %s\n", cpus.CPUs[id].Specs.Socket)
	fmt.Printf("│  ├─Architecture: %s\n", cpus.CPUs[id].Specs.Architecture)
	fmt.Printf("│  ├─Integrated GPU: %s\n", cpus.CPUs[id].Specs.IntegratedGPU)
	fmt.Printf("│  └─Maximum supported memory: %dGB\n", cpus.CPUs[id].Specs.MaximumSupportedMemoryGB)
	fmt.Println("├──Features:")
	for f := 0; f < len(cpus.CPUs[id].Features)-1; f++ {
		fmt.Printf("│  ├─%s\n", cpus.CPUs[id].Features[f])
	}
	fmt.Printf("│  └─%s\n", cpus.CPUs[id].Features[len(cpus.CPUs[id].Features)-1])
	fmt.Println("│")
	fmt.Printf("├─Overclockable?: %t\n", cpus.CPUs[id].Overclockable)
	fmt.Printf("└─Release date: %s\n", cpus.CPUs[id].ReleaseDate)
	fmt.Println("")

}
func dumpAllCpus(cpus CPUs) {
	for i := range cpus.CPUs {
		fmt.Println("═══════════════════════════════════════════════════════")
		fmt.Printf("[#%d]\n", i)
		fmt.Printf("ID: %s\n", cpus.CPUs[i].ID)
		fmt.Printf("├─Name: %s\n", cpus.CPUs[i].Name)
		fmt.Printf("├─Brand: %s\n", cpus.CPUs[i].Brand)
		fmt.Printf("├─Generation: %d\n", cpus.CPUs[i].Generation)
		fmt.Printf("├─Generation's codename: %s\n", cpus.CPUs[i].GenerationCodename)
		fmt.Printf("├─Series: %s\n", cpus.CPUs[i].Series)
		fmt.Println("│")
		fmt.Println("├──Specs:")
		fmt.Printf("│  ├─Total cores #: %d\n", cpus.CPUs[i].Specs.Cores)
		fmt.Printf("│  ├─Threads: %d\n", cpus.CPUs[i].Specs.Threads)
		fmt.Printf("│  ├─Cache: %.2fMB\n", cpus.CPUs[i].Specs.CacheMB)
		fmt.Printf("│  ├─Base frequency: %.2fGHz\n", cpus.CPUs[i].Specs.BaseFrequencyGHz)
		fmt.Printf("│  ├─Boost frequency: %.2fGHz\n", cpus.CPUs[i].Specs.BoostFrequencyGHz)
		fmt.Printf("│  ├─TDP: %d Watts\n", cpus.CPUs[i].Specs.TDPWatts)
		fmt.Printf("│  ├─Socket: %s\n", cpus.CPUs[i].Specs.Socket)
		fmt.Printf("│  ├─Architecture: %s\n", cpus.CPUs[i].Specs.Architecture)
		fmt.Printf("│  ├─Integrated GPU: %s\n", cpus.CPUs[i].Specs.IntegratedGPU)
		fmt.Printf("│  └─Maximum supported memory: %dGB\n", cpus.CPUs[i].Specs.MaximumSupportedMemoryGB)
		fmt.Println("│")
		fmt.Println("├──Features:")
		for f := 0; f < len(cpus.CPUs[i].Features)-1; f++ {
			fmt.Printf("│  ├─%s\n", cpus.CPUs[i].Features[f])
		}
		fmt.Printf("│  └─%s\n", cpus.CPUs[i].Features[len(cpus.CPUs[i].Features)-1])
		fmt.Println("│")
		fmt.Printf("├─Overclockable?: %t\n", cpus.CPUs[i].Overclockable)
		fmt.Printf("└─Release date: %s\n", cpus.CPUs[i].ReleaseDate)
		fmt.Println()
	}
}
func compareCpus(id1, id2 string, cpus CPUs) {}
func CaseInsensitiveContains(s, substring string) bool {
	// Code taken from https://stackoverflow.com/questions/24836044/case-insensitive-string-search-in-golang
	s, substring = strings.ToLower(s), strings.ToLower(substring)
	return strings.Contains(s, substring)
}
func main() {
	idToSearch := pflag.String("id", "none", "CPU ID to search in database")
	searchTerm := pflag.String("search", "none", "Term to search in the CPU names.")
	compareIDs := pflag.StringSlice("compare", nil, "CPUs IDs to compare (need two comma-sepparated values)")
	dumpAllFlag := pflag.Bool("dump-all", false, "Display all the CPUs' stats in the database")
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
			dumpID(idFound, cpus)

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
			fmt.Printf("No matches found for term: %s\n", *searchTerm)
		}
	}
	if len(*compareIDs) > 0 {
		if len(*compareIDs) == 2 {
			id1, id2 := (*compareIDs)[0], (*compareIDs)[1]
			compareCpus(id1, id2, cpus)

		} else {
			fmt.Println("Please provide two CPUs ids to compare.")
		}
	}
	if *dumpAllFlag {
		dumpAllCpus(cpus)
	}
	//dumpAllCpus(cpus)
}
