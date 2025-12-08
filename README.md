# CPUSpecCLI
## Introduction
CPUSpecCLI is a basic command-line tool to fetch info about differnt CPUs. It permit to know the number of cores, threads, maximum supported memory, etc.

It fetch the data from a json database filled with information taken from [Intel's website](https://ark.intel.com).

## Installation
To install CPUSpecCLI you can either grab one of the binaries in the Release tab, or build it yourself :
```bash
git clone https://github.com/hash-ak/cpuspeccli
cd CPUSpecCLI
go build .
```

## Usage
You can invoke the tool with a few options :  
```(binaryname) --dump-all``` : shows all the CPUs with their information from the database  
```(binaryname) --search "searchterm"``` " search the database for CPU(s) matching the search term  
```(binaryname) --id "cpuid"``` : show a single cpu with the specified ID (ID reffering to the ```ID``` field from the database, ex: intel-core-i3-1215UL)  
```(binaryname) --compare id1,id2``` : compare two CPUs by their IDs.
## TODO
- Add PassMark/other benchmark scores
- Add prices
