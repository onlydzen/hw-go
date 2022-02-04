package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const quantityOfWordsFromTop = 10

type WordFrequency struct {
	Word      string
	Frequency int
}

type AllWordsFrequency struct {
	WordFrequencyList []*WordFrequency
}

var (
	OnlyEnglishCyrillicAndDigits = regexp.MustCompile(`[^\wа-яА-я-]`)
	HyphenNotPartOfWord          = regexp.MustCompile(`[^\wа-яА-я-]-|-[^\wа-яА-я-]`)
)

func Top10(data string) []string {
	if len(data) == 0 {
		return nil
	}

	mapWordFrequency := getWordsFrequencyMap(parseData(data))

	awf := AllWordsFrequency{}
	awf.ConvertFromMap(mapWordFrequency)
	awf.SortLexicographical()

	return awf.GetTopWords(quantityOfWordsFromTop)
}

func (awf *AllWordsFrequency) SortLexicographical() {
	wf := awf.WordFrequencyList
	sort.Slice(wf, func(i, j int) bool {
		if wf[i].Frequency > wf[j].Frequency {
			return true
		}

		if wf[i].Frequency == wf[j].Frequency {
			return wf[i].Word < wf[j].Word
		}

		return false
	})
}

func (awf *AllWordsFrequency) ConvertFromMap(countMap map[string]int) {
	for k, v := range countMap {
		awf.WordFrequencyList = append(awf.WordFrequencyList, &WordFrequency{k, v})
	}
}

func (awf *AllWordsFrequency) GetTopWords(topWordsCount int) []string {
	if len(awf.WordFrequencyList) == 0 {
		return []string{""}
	}

	offset := topWordsCount
	if offset > len(awf.WordFrequencyList) {
		offset = len(awf.WordFrequencyList)
	}

	result := make([]string, 0)
	for _, v := range awf.WordFrequencyList[0:offset] {
		result = append(result, v.Word)
	}
	return result
}

func parseData(sourceData string) []string {
	// Double replacing because of possible collisions of RegExp matching after one run (incorrect sequence test)
	// TODO Optimise RegExp expression
	removedHyphenOutOfWords := HyphenNotPartOfWord.ReplaceAllString(sourceData, " ")
	removedHyphenOutOfWords = HyphenNotPartOfWord.ReplaceAllString(removedHyphenOutOfWords, " ")

	return OnlyEnglishCyrillicAndDigits.Split(removedHyphenOutOfWords, -1)
}

func getWordsFrequencyMap(cleanData []string) map[string]int {
	mapWordFrequency := map[string]int{}
	for _, value := range cleanData {
		if value == "" {
			continue
		}

		value = strings.ToLower(value)

		if _, ok := mapWordFrequency[value]; ok {
			counter := mapWordFrequency[value]
			mapWordFrequency[value] = counter + 1
			continue
		}

		mapWordFrequency[value] = 1
	}

	return mapWordFrequency
}
