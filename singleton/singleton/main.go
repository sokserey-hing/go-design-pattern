package main

// Singleton is a struct that holds the instance of the singleton object. It is a private variable. OR A component which is instantiated only once.
// For some components it only makes sense to have one instance in the system. For example, a single thread pool or a single file system., database connection, etc.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once, init() -- thread safety
//lazy loading

var once sync.Once // somthing get called only once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase { // but this might not be thread safe.
	once.Do(func() {
		// caps, err := readData(".\\capitals.txt")
		caps, err := readData("C:\\APIDevelopment\\Golang\\go-design-pattern\\singleton\\singleton\\capitals.txt")
		if err != nil {
			panic(err)
		}
		// fmt.Println(caps)
		db := singletonDatabase{caps}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {

	file, err := os.Open(path)
	// fmt.Println("File path:", path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	} // since the database is singleton(GetSingletonDatabase()), it will return the same instance of the database which become DIP violation.
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	} // this will represent the real life scenario where the database is injected.
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

func main() {
	fmt.Println("----------Singleton Pattern---------")
	fmt.Println("----------with DIP violation---------")
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Tokyo")
	totalPopulation := GetTotalPopulation([]string{"Tokyo", "Delhi"})

	expectedPop := 33200000 + 14300000
	ok := totalPopulation == (33200000 + 14300000)

	fmt.Println(ok)                                                     // true
	fmt.Println("Population of Tokyo is ", pop)                         // 33200000
	fmt.Println("Total Population of Tokyo and Delhi is ", expectedPop) // 47500000

	fmt.Println("----------without DIP violation---------")
	cities := []string{"Tokyo", "Delhi"}
	tp := GetTotalPopulationEx(GetSingletonDatabase(), cities)
	fmt.Println("Total Population of Tokyo and Delhi is ", tp)

	names := []string{"alpha", "beta"}
	tp = GetTotalPopulationEx(&DummyDatabase{}, names) // expected 3
	fmt.Println("Total Population of alpha and beta is ", tp)

}
