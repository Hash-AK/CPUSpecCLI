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

const (
	ColorGreen  = "\033[32m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

type CPUs struct {
	CPUs []CPU `json:"cpus"`
}
type CPU struct {
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	Brand              string     `json:"brand"`
	Generation         int        `json:"generation"`
	GenerationCodename string     `json:"generation_codename"`
	Series             string     `json:"series"`
	Specs              Specs      `json:"specs"`
	Benchmarks         Benchmarks `json:"benchmarks"`
	Features           []string   `json:"features"`
	Overclockable      bool       `json:"overclockable"`
	ReleaseDate        string     `json:"release_date"`
}
type Benchmarks struct {
	PassmarkMultiThreads int `json:"passmark_multithread_rating"`
	PassmarkSingleThread int `json:"passmark_singlethread_rating"`
}
type Specs struct {
	Cores                     int     `json:"cores"`
	Threads                   int     `json:"threads"`
	CacheMB                   float64 `json:"cache_mb"`
	BaseFrequencyGHz          float64 `json:"base_freq_ghz"`
	BoostFrequencyGHz         float64 `json:"boost_freq_ghz"`
	TDPWatts                  int     `json:"tdp_watts"`
	Socket                    string  `json:"socket"`
	Architecture              string  `json:"architecture"`
	IntegratedGPU             string  `json:"integrated_gpu"`
	MaximumSupportedMemoryGB  int     `json:"max_mem_supported_gb"`
	MaximumSupportedMemoryMHz int     `json:"max_mem_freq_mhz"`
}

func compareHigher(val1, val2 float64) (string, string) {
	if val1 > val2 {
		return ColorGreen, ColorRed
	} else if val2 > val1 {
		return ColorRed, ColorGreen
	}
	return ColorYellow, ColorYellow
}
func compareLower(val1, val2 float64) (string, string) {
	if val1 < val2 {
		return ColorGreen, ColorRed
	} else if val2 < val1 {
		return ColorRed, ColorGreen
	}
	return ColorYellow, ColorYellow
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
	fmt.Printf("│  ├─Maximum supported memory: %dGB\n", cpus.CPUs[id].Specs.MaximumSupportedMemoryGB)
	fmt.Printf("│  └─Maximum supported memory frequency: %dMT/s\n", cpus.CPUs[id].Specs.MaximumSupportedMemoryMHz)
	fmt.Println("├─Benchmarks Scores:")
	if cpus.CPUs[id].Benchmarks.PassmarkMultiThreads != 0 {
		fmt.Printf("│  ├─Passmark Multithread Rating: %d\n", cpus.CPUs[id].Benchmarks.PassmarkMultiThreads)
	} else {
		fmt.Printf("│  ├─Passmark Multithread Rating: %s\n", "N\\A")
	}
	if cpus.CPUs[id].Benchmarks.PassmarkSingleThread != 0 {
		fmt.Printf("│  ├─Passmark Singlethread Rating: %d\n", cpus.CPUs[id].Benchmarks.PassmarkSingleThread)
	} else {
		fmt.Printf("│  ├─Passmark Singlethread Rating: %s\n", "N\\A")
	}

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
		fmt.Printf("│  ├─Maximum supported memory: %dGB\n", cpus.CPUs[i].Specs.MaximumSupportedMemoryGB)
		fmt.Printf("│  └─Maximum supported memory frequency: %dMT/s\n", cpus.CPUs[i].Specs.MaximumSupportedMemoryMHz)
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
func compareCpus(id1, id2 int, cpus CPUs) {
	cpu1, cpu2 := cpus.CPUs[id1], cpus.CPUs[id2]
	fmt.Println()
	title := fmt.Sprintf("CPU Comparison: %s VS %s\n", cpu1.Name, cpu2.Name)
	fmt.Println(title)
	fmt.Println("═══════════════════════════════════════════════════════════════════")
	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "Spec", cpu1.Name, cpu2.Name)
	fmt.Println("───────────────────────────────────────────────────────────────────")
	c1, c2 := compareHigher(float64(cpu1.Specs.Cores), float64(cpu2.Specs.Cores))
	fmt.Printf("  %-16s │ %s%-24d%s │ %s%-24d%s\n", "Cores", c1, cpu1.Specs.Cores, ColorReset, c2, cpu2.Specs.Cores, ColorReset)

	t1, t2 := compareHigher(float64(cpu1.Specs.Threads), float64(cpu2.Specs.Threads))
	fmt.Printf("  %-16s │ %s%-24d%s │ %s%-24d%s\n", "Threads", t1, cpu1.Specs.Threads, ColorReset, t2, cpu2.Specs.Threads, ColorReset)

	ca1, ca2 := compareHigher(cpu1.Specs.CacheMB, cpu2.Specs.CacheMB)
	fmt.Printf("  %-16s │ %s%-24.2f%s │ %s%-24.2f%s\n", "Cache (MB)", ca1, cpu1.Specs.CacheMB, ColorReset, ca2, cpu2.Specs.CacheMB, ColorReset)

	b1, b2 := compareHigher(cpu1.Specs.BaseFrequencyGHz, cpu2.Specs.BaseFrequencyGHz)
	fmt.Printf("  %-16s │ %s%-24.2f%s │ %s%-24.2f%s\n", "Base (GHz)", b1, cpu1.Specs.BaseFrequencyGHz, ColorReset, b2, cpu2.Specs.BaseFrequencyGHz, ColorReset)

	bo1, bo2 := compareHigher(cpu1.Specs.BoostFrequencyGHz, cpu2.Specs.BoostFrequencyGHz)
	fmt.Printf("  %-16s │ %s%-24.2f%s │ %s%-24.2f%s\n", "Boost (GHz)", bo1, cpu1.Specs.BoostFrequencyGHz, ColorReset, bo2, cpu2.Specs.BoostFrequencyGHz, ColorReset)

	tdp1, tdp2 := compareLower(float64(cpu1.Specs.TDPWatts), float64(cpu2.Specs.TDPWatts))
	fmt.Printf("  %-16s │ %s%-24d%s │ %s%-24d%s\n", "TDP (Watts)", tdp1, cpu1.Specs.TDPWatts, ColorReset, tdp2, cpu2.Specs.TDPWatts, ColorReset)
	mm1, mm2 := compareHigher(float64(cpu1.Specs.MaximumSupportedMemoryGB), float64(cpu2.Specs.MaximumSupportedMemoryGB))
	fmt.Printf("  %-16s │ %s%-24d%s │ %s%-24d%s\n", "Max Mem (GB)", mm1, cpu1.Specs.MaximumSupportedMemoryGB, ColorReset, mm2, cpu2.Specs.MaximumSupportedMemoryGB, ColorReset)

	mh1, mh2 := compareHigher(float64(cpu1.Specs.MaximumSupportedMemoryMHz), float64(cpu2.Specs.MaximumSupportedMemoryMHz))
	fmt.Printf("  %-16s │ %s%-24d%s │ %s%-24d%s\n", "Max Mem Freq", mh1, cpu1.Specs.MaximumSupportedMemoryMHz, ColorReset, mh2, cpu2.Specs.MaximumSupportedMemoryMHz, ColorReset)

	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "Socket", cpu1.Specs.Socket, cpu2.Specs.Socket)
	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "Architecture", cpu1.Specs.Architecture, cpu2.Specs.Architecture)
	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "IGPU", cpu1.Specs.IntegratedGPU, cpu2.Specs.IntegratedGPU)
	fmt.Println("═══════════════════════════════════════════════════════════════════")
	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "General Info", "-", "-")
	fmt.Println("───────────────────────────────────────────────────────────────────")
	gen1, gen2 := ColorYellow, ColorYellow
	if cpu1.Brand == cpu2.Brand {
		if cpu1.Generation > cpu2.Generation {
			gen1, gen2 = ColorGreen, ColorRed
		} else if cpu2.Generation > cpu1.Generation {
			gen1, gen2 = ColorRed, ColorGreen
		}
	}
	fmt.Printf("  %-16s │ %s%-24d%s │ %s%-24d%s\n", "Generation", gen1, cpu1.Generation, ColorReset, gen2, cpu2.Generation, ColorReset)

	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "Brand", cpu1.Brand, cpu2.Brand)
	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "Series", cpu1.Series, cpu2.Series)
	fmt.Printf("  %-16s │ %-24t │ %-24t\n", "Overclockable?", cpu1.Overclockable, cpu2.Overclockable)
	fmt.Printf("  %-16s │ %-24s │ %-24s\n", "Release date", cpu1.ReleaseDate, cpu2.ReleaseDate)

}

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
			var numericalIDs []int
			for i := range cpus.CPUs {
				if cpus.CPUs[i].ID == id1 || cpus.CPUs[i].ID == id2 {
					numericalIDs = append(numericalIDs, i)
				}
			}
			if len(numericalIDs) == 2 {
				compareCpus(numericalIDs[0], numericalIDs[1], cpus)

			} else {
				fmt.Println("Please provide two VALID CPUs ids to compare.")
				if len(numericalIDs) == 1 {
					fmt.Printf("%s is a valid ID\n", cpus.CPUs[numericalIDs[0]].ID)

				}

			}

		} else {
			fmt.Println("Please provide two CPUs ids to compare.")
		}
	}
	if *dumpAllFlag {
		dumpAllCpus(cpus)
	}
	//dumpAllCpus(cpus)
}
