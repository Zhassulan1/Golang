package searchengine

import (
	"strings"
	findfilm "tsis_1/pkg/find-film"
)

// Implementing Levenshtein distance method using wagner-fischer algorithm
func wagnerFischer(s1, s2 string) float64 {
	len1, len2 := len(s1), len(s2)

	if len1 > len2 {
		s1, s2 = s2, s1
		len1, len2 = len2, len1
	}

	current_row := make([]int, len1+1)
	for i := 1; i <= len1; i++ {
		current_row[i] = i
	}

	for i := 1; i <= len2; i++ {
		prev_row := current_row
		current_row = make([]int, len1+1)
		current_row[0] = i

		for j := 1; j <= len1; j++ {
			// getting the cost of add, delete and replace operations
			add := prev_row[j] + 1
			delete := current_row[j-1] + 1
			replace := prev_row[j-1]

			if s1[j-1] != s2[i-1] {
				replace++
			}
			// the bottom right corner of the matrix is equal to the minimum of the three
			current_row[j] = min(add, delete, replace)
		}
	}

	return float64(current_row[len1])
}

func prepareString(text string) []string {
	text = strings.ToLower(text)
	text = strings.Replace(text, ".", " ", -1)
	text = strings.Replace(text, ",", " ", -1)
	list := strings.Split(text, " ")
	for i, v := range list {
		list[i] = strings.Trim(v, " ")
	}

	return list
}

func searchFilms(s1, s2 []string, textLen int) bool {
	textLenf := float64(textLen)
	for i := range s1 {
		for j := range s2 {
			distance := wagnerFischer(s1[i], s2[j])
			if distance <= 3 && textLenf >= 5 || distance/textLenf >= 0.2 {
				return true
			} else if distance/textLenf <= 0.28 {
				return true
			} else if distance <= 2 && textLen <= 4 {
				return true
			}
		}
	}

	return false
}

func exactMatch(text string) (findfilm.Film, bool) {
	for i := range findfilm.Films {
		if findfilm.Films[i].Name == text {
			return findfilm.Films[i], true
		}
	}
	return findfilm.Films[0], false
}

func Search(text string) []findfilm.Film {
	var films []findfilm.Film = findfilm.Films
	var Result []findfilm.Film

	if text[0] == '"' && text[len(text)-1] == '"' {
		film, ok := exactMatch(text[1 : len(text)-1])
		if ok {
			Result = append(Result, film)
		}
	}

	for i := range films {
		s1 := prepareString(films[i].Name)
		s2 := prepareString(text)
		textlen := len(text)
		if searchFilms(s1, s2, textlen) {
			Result = append(Result, films[i])
		}
	}
	return Result
}
