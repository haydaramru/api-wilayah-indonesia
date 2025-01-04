package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Regency struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

type District struct {
	ID        string `json:"id"`
	RegencyID string `json:"regency_id"`
	Name      string `json:"name"`
}

type Village struct {
	ID         string `json:"id"`
	DistrictID string `json:"district_id"`
	Name       string `json:"name"`
}

func main() {
	provinces := readProvinces("data/provinces.csv")
	regencies := readRegencies("data/regencies.csv")
	districts := readDistricts("data/districts.csv")
	villages := readVillages("data/villages.csv")

	regenciesByProvince := make(map[string][]Regency)
	for _, r := range regencies {
		regenciesByProvince[r.ProvinceID] = append(regenciesByProvince[r.ProvinceID], r)
	}

	districtsByRegency := make(map[string][]District)
	for _, d := range districts {
		districtsByRegency[d.RegencyID] = append(districtsByRegency[d.RegencyID], d)
	}

	villagesByDistrict := make(map[string][]Village)
	for _, v := range villages {
		villagesByDistrict[v.DistrictID] = append(villagesByDistrict[v.DistrictID], v)
	}

	createOutputDirs()

	writeJSON(filepath.Join("static", "api", "provinces.json"), provinces)

	for _, prov := range provinces {
		provinceRegencies := regenciesByProvince[prov.ID]
		outputFile := filepath.Join("static", "api", "regencies", prov.ID+".json")
		writeJSON(outputFile, provinceRegencies)
	}

	for _, r := range regencies {
		regencyDistricts := districtsByRegency[r.ID]
		outputFile := filepath.Join("static", "api", "districts", r.ID+".json")
		writeJSON(outputFile, regencyDistricts)
	}

	for _, d := range districts {
		districtVillages := villagesByDistrict[d.ID]
		outputFile := filepath.Join("static", "api", "villages", d.ID+".json")
		writeJSON(outputFile, districtVillages)
	}

	fmt.Println("All JSON files generated successfully!")
}

func createOutputDirs() {
	os.MkdirAll("static/api", 0755)
	os.MkdirAll("static/api/regencies", 0755)
	os.MkdirAll("static/api/districts", 0755)
	os.MkdirAll("static/api/villages", 0755)
}

func writeJSON(filePath string, data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal data to JSON: %v", err)
	}

	err = os.WriteFile(filePath, bytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write file %s: %v", filePath, err)
	}
}

func readProvinces(filePath string) []Province {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open provinces CSV: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		log.Fatalf("Failed to read header for provinces.csv: %v", err)
	}

	var results []Province
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read provinces CSV record: %v", err)
		}

		// record[0] = id
		// record[1] = name
		p := Province{
			ID:   strings.TrimSpace(record[0]),
			Name: strings.TrimSpace(record[1]),
		}
		results = append(results, p)
	}
	return results
}

func readRegencies(filePath string) []Regency {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open regencies CSV: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		log.Fatalf("Failed to read header for regencies.csv: %v", err)
	}

	var results []Regency
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read regencies CSV record: %v", err)
		}

		// record[0] = id
		// record[1] = province_id
		// record[2] = name
		re := Regency{
			ID:         strings.TrimSpace(record[0]),
			ProvinceID: strings.TrimSpace(record[1]),
			Name:       strings.TrimSpace(record[2]),
		}
		results = append(results, re)
	}
	return results
}

func readDistricts(filePath string) []District {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open districts CSV: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		log.Fatalf("Failed to read header for districts.csv: %v", err)
	}

	var results []District
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read districts CSV record: %v", err)
		}

		// record[0] = id
		// record[1] = regency_id
		// record[2] = name
		d := District{
			ID:        strings.TrimSpace(record[0]),
			RegencyID: strings.TrimSpace(record[1]),
			Name:      strings.TrimSpace(record[2]),
		}
		results = append(results, d)
	}
	return results
}

func readVillages(filePath string) []Village {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open villages CSV: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		log.Fatalf("Failed to read header for villages.csv: %v", err)
	}

	var results []Village
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read villages CSV record: %v", err)
		}

		// record[0] = id
		// record[1] = district_id
		// record[2] = name
		v := Village{
			ID:         strings.TrimSpace(record[0]),
			DistrictID: strings.TrimSpace(record[1]),
			Name:       strings.TrimSpace(record[2]),
		}
		results = append(results, v)
	}
	return results
}
