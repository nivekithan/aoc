package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	source_start      int
	destination_start int
	length            int
}

func (r *Range) has(num int) (bool, int) {
	if r.source_start > num {
		return false, 0
	}

	if r.source_start+r.length-1 < num {
		return false, 0
	}

	return true, num - r.source_start + r.destination_start
}

type Input struct {
	seeds                   []int
	seed_to_soil            []Range
	soil_to_fertilizer      []Range
	fertilizer_to_water     []Range
	water_to_light          []Range
	light_to_temperature    []Range
	temperature_to_humidity []Range
	humidity_to_location    []Range
}

type ParsingKind int

const (
	SeedToSoil ParsingKind = iota
	SoilToFertilizer
	FertilizerToWater
	WaterToLight
	LightToTemperature
	TemperatureToHumidity
	HumidityToLocation
)

func read_input() Input {
	file_name := "d5.data"

	file, err := os.Open(file_name)

	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := Input{}

	var parsingKind ParsingKind = -1

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seed_nums := strings.Split(strings.Split(line, "seeds: ")[1], " ")

			for _, nums_in_str := range seed_nums {
				num, err := strconv.Atoi(nums_in_str)

				if err != nil {
					log.Panic(err)
				}

				input.seeds = append(input.seeds, num)
			}

			continue
		}

		if strings.HasPrefix(line, "seed-to-soil map:") {
			parsingKind = SeedToSoil
			continue
		} else if strings.HasPrefix(line, "soil-to-fertilizer map:") {
			parsingKind = SoilToFertilizer
			continue
		} else if strings.HasPrefix(line, "fertilizer-to-water map:") {
			parsingKind = FertilizerToWater
			continue
		} else if strings.HasPrefix(line, "water-to-light map:") {
			parsingKind = WaterToLight
			continue
		} else if strings.HasPrefix(line, "light-to-temperature map:") {
			parsingKind = LightToTemperature
			continue
		} else if strings.HasPrefix(line, "temperature-to-humidity map:") {
			parsingKind = TemperatureToHumidity
			continue
		} else if strings.HasPrefix(line, "humidity-to-location map:") {
			parsingKind = HumidityToLocation
			continue
		}

		range_values := strings.Split(line, " ")

		if line == "" {
			continue
		}

		range_input := Range{}

		// log.Println(line)
		for i, range_in_str := range range_values {
			num, err := strconv.Atoi(range_in_str)

			if err != nil {
				log.Panic(err)
			}

			if i == 0 {
				range_input.destination_start = num
			} else if i == 1 {
				range_input.source_start = num
			} else if i == 2 {
				range_input.length = num
			}
		}

		if parsingKind == SeedToSoil {
			input.seed_to_soil = append(input.seed_to_soil, range_input)
		} else if parsingKind == SoilToFertilizer {
			input.soil_to_fertilizer = append(input.soil_to_fertilizer, range_input)
		} else if parsingKind == FertilizerToWater {
			input.fertilizer_to_water = append(input.fertilizer_to_water, range_input)
		} else if parsingKind == WaterToLight {
			input.water_to_light = append(input.water_to_light, range_input)
		} else if parsingKind == LightToTemperature {
			input.light_to_temperature = append(input.light_to_temperature, range_input)
		} else if parsingKind == TemperatureToHumidity {
			input.temperature_to_humidity = append(input.temperature_to_humidity, range_input)
		} else if parsingKind == HumidityToLocation {
			input.humidity_to_location = append(input.humidity_to_location, range_input)
		}
	}

	return input

}

func d5p1() {
	input := read_input()

	c := make(chan int, len(input.seeds))

	var wg sync.WaitGroup
	for _, seed := range input.seeds {
		wg.Add(1)
		log.Printf("Seed: %v", seed)
		go find_location_of_seed(&input, seed, &wg, c)
	}

	wg.Wait()
	close(c)

	min_location := 0
	for location := range c {
		if min_location == 0 {
			min_location = location
		} else {
			min_location = min(min_location, location)
		}
	}

	fmt.Println(min_location)
}

func d5p2() {
	input := read_input()

	c := make(chan int, len(input.seeds))

	var wg sync.WaitGroup
	for _, seed := range input.seeds {
		wg.Add(1)
		log.Printf("Seed: %v", seed)
		go find_location_of_seed(&input, seed, &wg, c)
	}

	wg.Wait()
	close(c)

	min_location := 0
	for location := range c {
		if min_location == 0 {
			min_location = location
		} else {
			min_location = min(min_location, location)
		}
	}

	fmt.Println(min_location)
}

func find_location_of_seed(input *Input, seed int, wg *sync.WaitGroup, c chan int) int {
	defer wg.Done()
	soil := get_seed_to_soil(input, seed)
	log.Printf("soil %v", soil)
	fertilizer := get_soil_to_fertilizer(input, soil)
	log.Printf("fertilizer %v", fertilizer)
	water := get_fertilizer_to_water(input, fertilizer)
	log.Printf("water %v", water)
	light := get_water_to_light(input, water)
	log.Printf("light %v", light)
	temperature := get_light_to_temperature(input, light)
	log.Printf("temperature %v", temperature)
	humidity := get_temperature_to_humidity(input, temperature)
	log.Printf("humiditiy %v", humidity)
	location := get_humidity_to_location(input, humidity)

	c <- location

	return location

}

func get_seed_to_soil(input *Input, seed int) int {
	for _, range_input := range input.seed_to_soil {
		is_present, soil := range_input.has(seed)

		if is_present {
			return soil
		}
	}

	return seed
}

func get_soil_to_fertilizer(input *Input, soil int) int {
	for _, range_input := range input.soil_to_fertilizer {
		is_present, fertilizer := range_input.has(soil)

		if is_present {
			return fertilizer
		}
	}

	return soil
}

func get_fertilizer_to_water(input *Input, fertilizer int) int {

	for _, range_input := range input.fertilizer_to_water {
		is_present, water := range_input.has(fertilizer)

		if is_present {
			return water
		}
	}

	return fertilizer
}

func get_water_to_light(input *Input, water int) int {
	for _, range_input := range input.water_to_light {
		is_present, light := range_input.has(water)

		if is_present {
			return light
		}
	}

	return water
}

func get_light_to_temperature(input *Input, light int) int {
	for _, range_input := range input.light_to_temperature {
		is_present, temp := range_input.has(light)

		if is_present {
			return temp
		}
	}

	return light
}

func get_temperature_to_humidity(input *Input, temp int) int {
	for _, range_input := range input.temperature_to_humidity {
		is_present, humdity := range_input.has(temp)

		if is_present {
			return humdity
		}
	}

	return temp
}

func get_humidity_to_location(input *Input, humidity int) int {
	for _, range_input := range input.humidity_to_location {
		is_present, location := range_input.has(humidity)

		if is_present {
			return location
		}
	}

	return humidity
}

func main() {
	d5p1()
}
