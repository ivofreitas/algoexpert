package main

import "fmt"

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

type OrgInfo struct {
	lowestCommonManager *OrgChart
	directReportCount   int
}

func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	// Write your code here.
	return getOrgInfo(org, reportOne, reportTwo).lowestCommonManager
}

func getOrgInfo(manager, reportOne, reportTwo *OrgChart) *OrgInfo {

	directReportCount := 0
	var lowestCommonManager *OrgChart
	for _, directReport := range manager.DirectReports {
		orgInfo := getOrgInfo(directReport, reportOne, reportTwo)
		if orgInfo.lowestCommonManager != nil {
			return orgInfo
		}
		directReportCount += orgInfo.directReportCount
	}
	if manager == reportOne || manager == reportTwo {
		directReportCount += 1
	}

	if directReportCount == 2 {
		lowestCommonManager = manager
	}

	return &OrgInfo{lowestCommonManager, directReportCount}
}

func getOrgCharts() map[rune]*OrgChart {
	orgCharts := map[rune]*OrgChart{}
	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		orgCharts[r] = &OrgChart{
			Name:          string(r),
			DirectReports: []*OrgChart{},
		}
	}
	return orgCharts
}

func (chart *OrgChart) addDirectReports(reports ...*OrgChart) {
	chart.DirectReports = append(chart.DirectReports, reports...)
}

func main() {
	orgCharts := getOrgCharts()
	orgCharts['A'].addDirectReports(orgCharts['B'], orgCharts['C'])
	orgCharts['B'].addDirectReports(orgCharts['D'], orgCharts['E'])
	orgCharts['C'].addDirectReports(orgCharts['F'], orgCharts['G'])
	orgCharts['D'].addDirectReports(orgCharts['H'], orgCharts['I'])

	fmt.Println(GetLowestCommonManager(orgCharts['A'], orgCharts['E'], orgCharts['I']))
}
