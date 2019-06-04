package consensus

import (
	"fmt"
	"github.com/ethereum/eth2-client-tests/tester/report"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"regexp"
	"strings"
)

const artemisFinalizedBlockRootRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastFinalizedBlockRoot>.*?)\""
const artemisFinalizedBlockStateRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastFinalizedBlockState>.*?)\""
const artemisJustifiedBlockRootRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastJustifiedBlockRoot>.*?)\""
const artemisJustifiedBlockStateRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastJustifiedBlockState>.*?)\""


func CheckFinalizedBlockRoot(testReportOutput string, blockchain string, files []string) {
	log.Info("entering check finalized block root")
	failed, testReports, stdout := checkLogs(artemisFinalizedBlockRootRegexp, files, testReportOutput)
	if failed {
		log.Info("Block roots did not match")
	} else {
		log.Info("All block roots match")
	}
	if testReportOutput != "" {
		report.WriteReport(fmt.Sprintf("%s.", blockchain), testReports, stdout, "", testReportOutput)
	}
}



func CheckFinalizedStateRoot(testReportOutput string, blockchain string, files []string) {
	log.Info("entering check finalized block state")
	failed, testReports, stdout := checkLogs(artemisFinalizedBlockStateRegexp, files, testReportOutput)
	if failed {
		log.Info("Block states did not match")
	} else {
		log.Info("All block states match")
	}
	if testReportOutput != "" {
		report.WriteReport(fmt.Sprintf("%s.", blockchain), testReports, stdout, "", testReportOutput)
	}
}

func CheckJustifiedBlockRoot(testReportOutput string, blockchain string, files []string) {
	log.Info("entering check justified block root")
	failed, testReports, stdout := checkLogs(artemisJustifiedBlockRootRegexp, files, testReportOutput)
	if failed {
		log.Info("Block roots did not match")
	} else {
		log.Info("All block roots match")
	}
	if testReportOutput != "" {
		report.WriteReport(fmt.Sprintf("%s.", blockchain), testReports, stdout, "", testReportOutput)
	}
}

func CheckJustifiedStateRoot(testReportOutput string, blockchain string, files []string) {
	log.Info("entering check justified block state")
	failed, testReports, stdout := checkLogs(artemisJustifiedBlockStateRegexp, files, testReportOutput)
	if failed {
		log.Info("Block states did not match")
	} else {
		log.Info("All block states match")
	}
	if testReportOutput != "" {
		report.WriteReport(fmt.Sprintf("%s.", blockchain), testReports, stdout, "", testReportOutput)
	}
}

func checkLogs(exprCaptureRegexp string, files []string, testReportOutput string) (bool, []report.TestReport, string) {
	r := regexp.MustCompile(exprCaptureRegexp)

	totalBlockRoots := [][]string{}
	for _, file := range files {
		counter := 0
		log.Info("Opening file ", file)
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file, err)
		}
		matches := r.FindAllStringSubmatch(string(b), -1)
		for _, lineMatch := range matches {
			if len(totalBlockRoots) <= counter {
				totalBlockRoots = append(totalBlockRoots, []string{})
			}
			totalBlockRoots[counter] = append(totalBlockRoots[counter], lineMatch[2])
			counter++
		}
	}

	testReports := []report.TestReport{}
	stdout := []string{}
	failed := false
	for epochIndex, epochMatch := range totalBlockRoots {
		var firstEntry string
		stdout = append(stdout, strings.Join(epochMatch, ","))
		for nodeIndex, epochEntry := range epochMatch {
			if firstEntry == "" {
				firstEntry = epochEntry
			} else {
				name := fmt.Sprintf("epoch-%d-%d", epochIndex, nodeIndex)
				if firstEntry != epochEntry {
					message := fmt.Sprintf("Epoch didn't match: %s vs %s", firstEntry, epochEntry)

					testReports = append(testReports, report.TestReport{
						Failed : true,
						Name: name,
						Message: message,
					})
					failed = true
				} else {
					testReports = append(testReports, report.TestReport{
						Failed : false,
						Name: name,
					})
				}
			}
		}
	}

	return failed, testReports, strings.Join(stdout, "\n")
}
