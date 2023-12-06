package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var DigitRegex = regexp.MustCompile(`\d+`)

type AlmanacMapRecord struct {
	dest int
	src  int
	rng  int
}

type AlmanacMap struct {
	records []AlmanacMapRecord
}

type Almanac struct {
	seeds                 []int
	seedToSoil            AlmanacMap
	soilToFertilizer      AlmanacMap
	fertilizerToWater     AlmanacMap
	waterToLight          AlmanacMap
	lightToTemperature    AlmanacMap
	temperatureToHumidity AlmanacMap
	humidityToLocation    AlmanacMap
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := ""
	for scanner.Scan() {
		line := scanner.Text()
		lines = lines + line + "\n"
	}

	almanac := parse(&lines)

	var minLocation int
	for _, seed := range almanac.seeds {
		soil := findCorrespondent(seed, &almanac.seedToSoil)
		fertilizer := findCorrespondent(soil, &almanac.soilToFertilizer)
		water := findCorrespondent(fertilizer, &almanac.fertilizerToWater)
		light := findCorrespondent(water, &almanac.waterToLight)
		temperature := findCorrespondent(light, &almanac.lightToTemperature)
		humidity := findCorrespondent(temperature, &almanac.temperatureToHumidity)
		location := findCorrespondent(humidity, &almanac.humidityToLocation)

		if minLocation == 0 || location < minLocation {
			minLocation = location
		}
	}

	fmt.Println(minLocation)
}

func parse(lines *string) *Almanac {
	blocks := strings.Split(*lines, "\n\n")

	return &Almanac{
		seeds:                 *parseNumbers(blocks[0]),
		seedToSoil:            *parseMap(blocks[1]),
		soilToFertilizer:      *parseMap(blocks[2]),
		fertilizerToWater:     *parseMap(blocks[3]),
		waterToLight:          *parseMap(blocks[4]),
		lightToTemperature:    *parseMap(blocks[5]),
		temperatureToHumidity: *parseMap(blocks[6]),
		humidityToLocation:    *parseMap(blocks[7]),
	}
}

func parseMap(m string) *AlmanacMap {
	parts := strings.Split(m, ":")

	records := []AlmanacMapRecord{}
	for _, line := range strings.Split(parts[1], "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		numbers := DigitRegex.FindAllString(line, -1)
		records = append(records, AlmanacMapRecord{
			dest: parseInt(numbers[0]),
			src:  parseInt(numbers[1]),
			rng:  parseInt(numbers[2]),
		})
	}

	return &AlmanacMap{records}
}

func parseNumbers(s string) *[]int {
	numbers := []int{}
	for _, n := range DigitRegex.FindAllString(s, -1) {
		numbers = append(numbers, parseInt(n))
	}
	return &numbers
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func findCorrespondent(n int, m *AlmanacMap) int {
	for _, r := range m.records {
		if n >= r.src && n <= r.src+r.rng-1 {
			return r.dest + (n - r.src)
		}
	}

	return n
}
