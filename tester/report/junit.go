package report

import (
	"bufio"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/prometheus/common/log"
)

type TestReport struct {
	Failed  bool
	Name    string
	Message string
	Logs    string
}

type Report struct {
	Name          string
	Id            string
	TestsCount    int
	FailuresCount int
	TestReports   []TestReport
	Stdout        string
	Stderr        string
}

func WriteReport(reportName string, testReports []TestReport, stdout string, stderr string, outputFilePath string) {
	now := time.Now()
	failuresCount := 0
	for _, report := range testReports {
		if report.Failed {
			failuresCount++
		}
	}

	report := Report{
		Name: reportName,
		Id: fmt.Sprintf("%d%02d%02d-%02d%02d%02d",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second()),
		TestsCount:    len(testReports),
		FailuresCount: failuresCount,
		TestReports:   testReports,
		Stdout:        stdout,
		Stderr:        stderr,
	}

	t, err := template.New("junit report").Parse(
		`<?xml version="1.0" encoding="UTF-8" ?> 
      <testsuite id="{{.Id}}" name="{{.Name}}" tests="{{.TestsCount}}" failures="{{.FailuresCount}}" time="0">
         {{range $i, $report := .TestReports}}
         <testcase name="{{$report.Name}}" time="0">
            {{ if $report.Failed }}
            <failure message="{{$report.Message}}" type="ERROR">
              {{$report.Logs}}
            </failure>
            {{ end }}
          </testcase>
          {{end}}
          <system-out><![CDATA[{{.Stdout}}]]></system-out>
          <system-err><![CDATA[{{.Stderr}}]]></system-err>
       </testsuite>`)
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}
	f, err := os.Create(outputFilePath)
	defer f.Close()
	w := bufio.NewWriter(f)
	err = t.Execute(w, report)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
	err = w.Flush()
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
