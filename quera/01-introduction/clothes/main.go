package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	OptionShirt        = "SHIRT"
	OptionPants        = "PANTS"
	OptionCap          = "CAP"
	OptionCoat         = "COAT"
	OptionJacket       = "JACKET"
	OptionJacketOrCoat = "JACKET_COAT"

	SeasonSpring = "SPRING"
	SeasonSummer = "SUMMER"
	SeasonFall   = "FALL"
	SeasonWinter = "WINTER"
)

var PrintOrder = []string{OptionCoat, OptionShirt, OptionPants, OptionCap, OptionJacket}

func main() {
	options, season := getInputs()
	removeOptionsBasedOnSeason(options, season)
	printResults(options, season)
}

func getInputs() (options map[string][]string, season string) {
	options = make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := scanner.Text()
		inputs := strings.Split(line, ":")
		options[inputs[0]] = strings.Split(strings.TrimSpace(inputs[1]), " ")
	}
	scanner.Scan()
	season = scanner.Text()
	return
}

/*
- [x] پیراهن و شلوار در همه فصل‌ها حتما پوشیده می‌شوند.
- [x] کلاه آفتابی در تابستان حتما پوشیده می‌شود و در بهار و پاییز ممکن است پوشیده شود.
- [x] کت در فصل‌های بهار و پاییز و زمستان ممکن است پوشیده شود.
- [x] ژاکت فقط در زمستان ممکن است پوشیده شود،
- [ ] در زمستان بین ژاکت و کت دقیقا یکی پوشیده می‌شود.
- [x] در پاییز کت زرد یا نارنجی پوشیده نمی‌شود. (yellow و orange)
*/

func removeOptionsBasedOnSeason(options map[string][]string, season string) {
	// cap logics
	if season == SeasonWinter {
		delete(options, OptionCap)
	} else if season != SeasonSummer {
		options[OptionCap] = append(options[OptionCap], "")
	}

	// coat logics
	if season == SeasonSummer {
		delete(options, OptionCoat)
	} else if season != SeasonWinter {
		options[OptionCoat] = append(options[OptionCoat], "")
	}

	// jacket logics
	if season == SeasonFall {
		for i, color := range options[OptionJacket] {
			if color == "orange" {
				options[OptionJacket] = append(options[OptionJacket][:i], options[OptionJacket][i+1:]...)
				break
			}
		}

		for i, color := range options[OptionJacket] {
			if color == "yellow" {
				options[OptionJacket] = append(options[OptionJacket][:i], options[OptionJacket][i+1:]...)
				break
			}
		}
	}

	if season != SeasonWinter {
		delete(options, OptionJacket)
	}
}

type clothesOption struct {
	clothes string
	color   string
}

func printResults(options map[string][]string, season string) {
	clothesOptions := make(map[string][]*clothesOption)
	for option, items := range options {
		clothesOptions[option] = make([]*clothesOption, 0)
		for _, item := range items {
			if item == "" {
				clothesOptions[option] = append(clothesOptions[option], nil)
			}
			clothesOptions[option] = append(clothesOptions[option], &clothesOption{option, item})
		}
	}
	_ = options

	// jacket and coat logics
	if season == SeasonWinter {
		clothesOptions[OptionJacketOrCoat] = append(clothesOptions[OptionCoat], clothesOptions[OptionJacket]...)
		delete(clothesOptions, OptionCoat)
		delete(clothesOptions, OptionJacket)
	}

	mapping := make(map[int]string)
	buckets := []int{}
	count := 1
	for option, items := range clothesOptions {
		// fmt.Println(option, len(items), items)
		if len(items) == 0 {
			continue
		}
		mapping[count] = option
		buckets = append(buckets, count)
		count *= len(items)
	}

	// fmt.Println(clothesOptions)

	for i := 0; i < count; i++ {
		val := i
		combination := make(map[string]string)
		for j := len(buckets) - 1; j >= 0; j-- {
			bucket := buckets[j]
			option := clothesOptions[mapping[bucket]][val/bucket]
			combination[option.clothes] = option.color
			val %= bucket
		}

		result := ""
		for _, cloth := range PrintOrder {
			if color, found := combination[cloth]; found {
				result += cloth + ": " + color + " "
			}
		}
		fmt.Println(strings.TrimSpace(result))
	}
}
