package main

import (
	"embed"
	"encoding/json"
	"fmt"
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

func main() {
	dbJson, err := DB.ReadFile("cpus.json")
	if err != nil {
		panic(err)
	}
	var cpus CPUs
	json.Unmarshal([]byte(dbJson), &cpus)
	for i := range cpus.CPUs {
		fmt.Println(cpus.CPUs[i].ID)
		fmt.Println(cpus.CPUs[i].Name)
		fmt.Println(cpus.CPUs[i].Brand)
		fmt.Println(cpus.CPUs[i].Generation)
		fmt.Println(cpus.CPUs[i].GenerationCodename)
		fmt.Println(cpus.CPUs[i].Series)
		fmt.Println(cpus.CPUs[i].Specs.Cores)

	}

}
