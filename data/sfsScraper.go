package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/djimenez/iconv-go"
)

func Scrape() []Course {

	res, err := http.Get("https://www.bms-fw.bayern.de/Navigation/Public/lastminute.aspx")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	cookies := res.Header.Get("Set-Cookie")
	session := strings.Split(cookies, ";")[0]
	csvResp, err := getCsv(session)
	defer csvResp.Body.Close()
	//bodyText, _ := io.ReadAll(csvResp.Body)
	//fmt.Printf("%s\n", bodyText)

	// Convert []byte to io.Reader for the CSV parser
	utfBody, err := iconv.NewReader(csvResp.Body, "iso-8859-1", "utf-8")
	csvReader := csv.NewReader(utfBody)
	csvReader.Comma = ';'

	// Read all records
	courseRecords, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	courses := make([]Course, 0)

	// Link zur Kursübersicht, nicht zur Anmeldung
	linkpattern := "https://www.sfsg.de/lehrgaenge/lehrgangsangebot/detailansicht/%s/"

	for i, record := range courseRecords {
		if i == 0 {
			// Skip header row if CSV has headers
			continue
		}
		// Access columns by index
		coursenumber := record[1]
		coursename := record[2]
		pattern := regexp.MustCompile(`^.+ .+ (.+) .+ .+$`)
		ctype := pattern.ReplaceAllString(coursenumber, "$1")
		start := record[3]
		end := record[4]
		place := record[5]
		free := record[6]
		link := fmt.Sprintf(linkpattern, ctype)

		c := Course{coursenumber, coursename, ctype, start, end, free, link, place}
		//fmt.Printf("Lehrgang %d: %s, %s\n", i, coursenumber, coursename)
		courses = append(courses, c)

	}

	return courses
}

func getCsv(sessionId string) (*http.Response, error) {
	client := &http.Client{}
	var data = strings.NewReader("__EVENTTARGET=ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlCoursesList%24CtrlGrid%24ctl09%24ctl09&__EVENTARGUMENT=&__LASTFOCUS=&__SKM_VIEWSTATEID=638969030304292386-011a8c9c-8de6-4834-8094-f2605d227339.vs&__VIEWSTATE=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24ctl00%24CtrlHelpArticleDetails%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24ctl00%24CtrlHelpArticleEdit%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlCoursesList%24CtrlGrid%24ctl01%24HfCurrentlySelectedRow=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlCoursesList%24CtrlGrid%24ctl09%24ctl02=1&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlCoursesList%24CtrlGrid%24ctl09%24ctl06=10&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlCourseDetailsLong%24ctl00%24CtrlHelpArticleDetails%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlCourseDetailsLong%24ctl00%24CtrlHelpArticleEdit%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24HfSelectedTabIndex=0&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24HfSelectedTabId=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24ctl03%24CtrlHelpArticleDetails%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24ctl03%24CtrlHelpArticleEdit%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24ctl05%24CtrlHelpArticleDetails%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24ctl05%24CtrlHelpArticleEdit%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24DtpArrivalDate%24TbDatePicker=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24TbSpecialDietInfo=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24TbComment=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24CtrlSignatureListEdit%24ctl00%24CtrlHelpArticleDetails%24_tbHelpContent=&ctl00%24ctl00%24CM%24CM%24CtrlCourses%24CtrlRegistrationEdit%24CtrlBmsTabs%24CtrlSignatureListEdit%24ctl00%24CtrlHelpArticleEdit%24_tbHelpContent=&__SCROLLPOSITIONX=0&__SCROLLPOSITIONY=0")
	req, err := http.NewRequest("POST", "https://www.bms-fw.bayern.de/Navigation/Public/lastminute.aspx", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://www.bms-fw.bayern.de/Navigation/Public/lastminute.aspx")
	req.Header.Set("Cookie", sessionId)
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}
