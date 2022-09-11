package helmservice

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

func (s *service) ShowReleaseInfo(name, namespace, info, output string) (interface{}, error) {
	infos := []string{"hooks", "manifest", "notes", "values"}
	infoMap := map[string]bool{}
	for _, i := range infos {
		infoMap[i] = true
	}
	if _, ok := infoMap[info]; !ok {
		return nil, fmt.Errorf("bad info %s, release info only support hooks/manifest/notes/values", info)
	}

	//helmClient, err := helmclient.GetHelmClient(myClusterApiServer, myClusterToken, myClusterCa, namespace)
	//if err != nil {
	//	return nil, err
	//}

	actionConfig := new(action.Configuration)
	if info == "values" {
		//output := ctx.Query("output")
		// get values output format
		if output == "" {
			output = "json"
		}

		if output != "json" && output != "yaml" {
			return nil, fmt.Errorf("invalid format type %s, output only support json/yaml", output)
		}

		client := action.NewGetValues(actionConfig)
		results, err := client.Run(name)
		if err != nil {
			return nil, err
		}
		if output == "yaml" {
			obj, err := yaml.Marshal(results)
			if err != nil {
				return nil, err
			}
			return string(obj), nil
		}
		return results, nil
	}

	client := action.NewGet(actionConfig)
	results, err := client.Run(name)
	if err != nil {
		return nil, err
	}
	// TODO: support all
	if info == "hooks" {
		if len(results.Hooks) < 1 {
			return []*release.Hook{}, nil
		}
		return results.Hooks, nil
	} else if info == "manifest" {
		return results.Manifest, nil

	} else if info == "notes" {
		return results.Info.Notes, nil
	} else {
		return results.Info, nil
	}
}
