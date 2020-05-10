package service

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// fetchMajor downloads a major's website and returns a populated major
func fetchMajor(id int) (*major, error) {
	major := &major{}
	major.externalID = id
	document, err := getDocument(major.getURL())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch major %d; %v", id, err)
	}
	rawName := document.Find("h1").First().Text()
	major.name = strings.Title(strings.TrimSpace(strings.Split(rawName, "(")[0]))
	return major, nil
}

// fetchProgram downloads a major's website and returns a populated program
func fetchProgram(majorID, programID int) (*program, error) {
	p := &program{externalID: programID, major: majorID}
	document, err := getDocument(p.getURL())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch program %d at URL %s; %v", programID, p.getURL(), err)
	}
	parenthesisRegex := regexp.MustCompile(`\(([^\)]+)\)`)
	digitRegex := regexp.MustCompile(`\d+`)
	programRawName := document.Find("h1").First().Text()
	programRawName = parenthesisRegex.FindString(programRawName)
	p.name = strings.Title(programRawName[1 : len(programRawName)-1])
	p.year, err = strconv.Atoi(digitRegex.FindString(p.name))
	if err != nil {
		return nil, fmt.Errorf("failed parsing program year for program %d; %v", programID, err)
	}
	return p, nil
}

// fetchProgramCourses downloads a program's website and returns the list of offered courses.
func fetchProgramCourses(p *program) ([]*course, error) {
	document, err := getDocument(p.getURL())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch program %d at URL %s; %v", p.externalID, p.getURL(), err)
	}

	parenthesisRegex := regexp.MustCompile(`\(([^\)]+)\)`)
	digitRegex := regexp.MustCompile(`\d+`)

	programRawName := document.Find("h1").First().Text()
	programRawName = parenthesisRegex.FindString(programRawName)
	p.name = strings.Title(programRawName[1 : len(programRawName)-1])
	p.year, _ = strconv.Atoi(digitRegex.FindString(p.name))

	uls := document.Find("#info-contenido ul").Last().Find("p,h3,h2")
	semester, mandatory := 0, true
	courseNameRegex := regexp.MustCompile(`, \d+ créditos\.`)

	var courses []*course

	uls.Each(func(i int, s *goquery.Selection) {
		switch {
		case s.Is("h3"):
			semester++
		case s.Is("h2") && i > 0:
			mandatory = false
			semester = 0
		case s.Is("p") && s.Find("a").Length() != 0:
			rawText := strings.TrimSpace(s.Text())
			creditsLocation, cutIndex := courseNameRegex.FindStringIndex(rawText), len(rawText)

			var credits int
			if len(creditsLocation) > 0 {
				cutIndex = creditsLocation[0]
				credits, _ = strconv.Atoi(digitRegex.FindString(rawText[cutIndex:]))
			}

			course := course{
				program:   p.externalID,
				name:      rawText[:cutIndex],
				semester:  uint(semester),
				credits:   uint(credits),
				mandatory: mandatory,
			}

			courseURL, ok := s.Find("a").First().Attr("href")
			if ok {
				seps := strings.Split(courseURL, "/")
				course.externalID, _ = strconv.Atoi(seps[len(seps)-1])
			}
			course.syllabusURL = course.getURL()
			courses = append(courses, &course)
		}
	})

	return courses, nil
}
