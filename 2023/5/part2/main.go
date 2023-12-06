package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var DigitRegex = regexp.MustCompile(`\d+`)

type AlmanacMapRecord struct {
	dest int64
	src  int64
	rng  int64
}

type AlmanacMap struct {
	records []AlmanacMapRecord
}

type SeedRange struct {
	start int64
	end   int64
}

type Almanac struct {
	seedRanges            []SeedRange
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

	minLocation := solve(almanac)

	fmt.Println(minLocation)
}

func solve(almanac *Almanac) int64 {
	wg := sync.WaitGroup{}
	lock := sync.RWMutex{}

	var minLocation int64 = math.MaxInt64
	for _, seedRange := range almanac.seedRanges {
		wg.Add(1)

		go func(start, end int64) {
			fmt.Printf("Solving range %d - %d\n", start, end)

			for seed := start; seed <= end; seed++ {
				soil := findCorresponding(seed, &almanac.seedToSoil)
				fertilizer := findCorresponding(soil, &almanac.soilToFertilizer)
				water := findCorresponding(fertilizer, &almanac.fertilizerToWater)
				light := findCorresponding(water, &almanac.waterToLight)
				temperature := findCorresponding(light, &almanac.lightToTemperature)
				humidity := findCorresponding(temperature, &almanac.temperatureToHumidity)
				location := findCorresponding(humidity, &almanac.humidityToLocation)

				lock.RLock()
				minLocation = min(location, minLocation)
				lock.RUnlock()
			}

			fmt.Printf("Range %d to %d done! Min location so far: %d\n", start, end, minLocation)
			wg.Done()
		}(seedRange.start, seedRange.end)
	}
	wg.Wait()

	return minLocation
}

func parse(lines *string) *Almanac {
	blocks := strings.Split(*lines, "\n\n")

	return &Almanac{
		seedRanges:            *parseSeedRanges(blocks[0]),
		seedToSoil:            *parseMap(blocks[1]),
		soilToFertilizer:      *parseMap(blocks[2]),
		fertilizerToWater:     *parseMap(blocks[3]),
		waterToLight:          *parseMap(blocks[4]),
		lightToTemperature:    *parseMap(blocks[5]),
		temperatureToHumidity: *parseMap(blocks[6]),
		humidityToLocation:    *parseMap(blocks[7]),
	}
}

func parseSeedRanges(s string) *[]SeedRange {
	numbers := parseNumbers(s)
	ranges := []SeedRange{}

	for i := 0; i < len(*numbers); i += 2 {
		start := (*numbers)[i]
		end := start + (*numbers)[i+1]
		ranges = append(ranges, SeedRange{start, end})
	}

	return &ranges
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

func parseNumbers(s string) *[]int64 {
	numbers := []int64{}
	for _, n := range DigitRegex.FindAllString(s, -1) {
		numbers = append(numbers, parseInt(n))
	}
	return &numbers
}

func parseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func findCorresponding(n int64, m *AlmanacMap) int64 {
	for _, r := range m.records {
		if n >= r.src && n <= r.src+r.rng-1 {
			return r.dest + (n - r.src)
		}
	}

	return n
}
