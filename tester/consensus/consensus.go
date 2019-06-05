package consensus

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/ethereum/eth2-client-tests/tester/report"
	"github.com/prometheus/common/log"
)

const artemisFinalizedBlockRootRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastFinalizedBlockRoot>.*?)\""
const artemisFinalizedStateRootRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastFinalizedBlockState>.*?)\""
const artemisJustifiedBlockRootRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastJustifiedBlockRoot>.*?)\""
const artemisJustifiedStateRootRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastJustifiedBlockState>.*?)\""

const prysmFinalizedBlockRootRegexp = "TEST_EPOCH.*\"epoch\":(?P<epoch>\\d+?).*\"block_root\":\"(?P<lastFinalizedBlockRoot>.*?)\""

func CheckFinalizedBlockRoot(testReportOutput string, blockchain string, files []string) {
	log.Info("entering check finalized block root")
	var matchExpression string
	switch blockchain {
	case "artemis":
		matchExpression = artemisFinalizedBlockRootRegexp
	case "prysm":
		matchExpression = ""
	default:
		log.Fatalf("Unsupported blockchain %s", blockchain)
	}
	failed, testReports, stdout := checkLogs(matchExpression, files, testReportOutput)
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
	failed, testReports, stdout := checkLogs(artemisFinalizedStateRootRegexp, files, testReportOutput)
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
	failed, testReports, stdout := checkLogs(artemisJustifiedStateRootRegexp, files, testReportOutput)
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

	valuesCaptured := [][]string{}
	for _, file := range files {
		log.Info("Opening file ", file)
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file, err)
		}
		counter := 0
		matches := r.FindAllStringSubmatch(string(b), -1)
		for _, lineMatch := range matches {
			if len(valuesCaptured) <= counter {
				valuesCaptured = append(valuesCaptured, []string{})
			}
			valuesCaptured[counter] = append(valuesCaptured[counter], lineMatch[2])
			counter++
		}
	}

	testReports := []report.TestReport{}
	stdout := []string{}
	failed := false
	colSizes := []int{}
	for epochIndex, epochMatch := range valuesCaptured {
		var firstEntry string
		stdout = append(stdout, "|" + strings.Join(epochMatch, "|") + "|")
		for nodeIndex, epochEntry := range epochMatch {
			if nodeIndex >= len(colSizes) {
				colSizes = append(colSizes, len(epochEntry))
			} else {
				if len(epochEntry) > colSizes[nodeIndex] {
					colSizes[nodeIndex] = len(epochEntry)
				}
			}

			if firstEntry == "" {
				firstEntry = epochEntry
			} else {
				name := fmt.Sprintf("epoch-%d-%d", epochIndex, nodeIndex)
				if firstEntry != epochEntry {
					message := fmt.Sprintf("Epoch didn't match: %s vs %s", firstEntry, epochEntry)

					testReports = append(testReports, report.TestReport{
						Failed:  true,
						Name:    name,
						Message: message,
					})
					failed = true
				} else {
					testReports = append(testReports, report.TestReport{
						Failed: false,
						Name:   name,
					})
				}
			}
		}
	}
	sum := len(colSizes) + 1
	header := "|"
	for index, colSize := range colSizes {
		sum += colSize
		filename := files[index]
		spaces := colSize - len(filename) -1
		if spaces < 0 {
			spaces = 0
		}
		header += " " + filename + strings.Repeat(" ", spaces) + "|"
	}

	stdout = append([]string{header, strings.Repeat("-", sum)}, stdout...)

	return failed, testReports, strings.Join(stdout, "\n")
}
