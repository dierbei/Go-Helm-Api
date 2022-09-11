package helmservice

import (
	"fmt"
	"strings"

	"githup.com/dierbei/go-helm-api/internal/pkg/helmclient"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
)

func (s *service) GetChartInfo(name, info, version string) (interface{}, error) {
	client := action.NewShow(action.ShowAll)
	client.Version = version
	if info == string(action.ShowChart) {
		client.OutputFormat = action.ShowChart
	} else if info == string(action.ShowReadme) {
		client.OutputFormat = action.ShowReadme
	} else if info == string(action.ShowValues) {
		client.OutputFormat = action.ShowValues
	} else if info == string(action.ShowAll) {
		client.OutputFormat = action.ShowAll
	} else {
		return nil, fmt.Errorf("bad info %s, chart info only support readme/values/chart", info)
	}

	settings := helmclient.GetHelmSettings().GetSettings()
	cp, err := client.ChartPathOptions.LocateChart(name, settings)
	if err != nil {
		return nil, err
	}

	chrt, err := loader.Load(cp)
	if err != nil {
		return nil, err
	}

	if client.OutputFormat == action.ShowChart {
		return chrt.Metadata, nil
	}

	if client.OutputFormat == action.ShowValues {
		var values string
		for _, v := range chrt.Raw {
			if v.Name == chartutil.ValuesfileName {
				values = string(v.Data)
				break
			}
		}
		return values, nil
	}

	if client.OutputFormat == action.ShowReadme {
		return findReadme(chrt.Files).Data, nil
	}

	// client.OutputFormat == action.ShowAll
	values := make([]*file, 0, len(chrt.Raw))
	for _, v := range chrt.Raw {
		values = append(values, &file{
			Name: v.Name,
			Data: string(v.Data),
		})
	}
	return values, nil
}

var readmeFileNames = []string{"readme.md", "readme.txt", "readme"}

func findReadme(files []*chart.File) (file *chart.File) {
	for _, file := range files {
		for _, n := range readmeFileNames {
			if strings.EqualFold(file.Name, n) {
				return file
			}
		}
	}
	return nil
}

type file struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
