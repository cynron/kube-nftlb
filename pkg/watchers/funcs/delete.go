package funcs

import (
	"fmt"

	defaults "github.com/zevenet/kube-nftlb/pkg/defaults"
	json "github.com/zevenet/kube-nftlb/pkg/json"
	request "github.com/zevenet/kube-nftlb/pkg/request"
	types "github.com/zevenet/kube-nftlb/pkg/types"
	v1 "k8s.io/api/core/v1"
)

// DeleteNftlbFarm deletes any nftlb farm given a Service object.
func DeleteNftlbFarm(service *v1.Service) {
	farmName := service.ObjectMeta.Name
	response := deleteNftlbRequest(farmName)
	// Prints info
	printDeleted("Farm", farmName, "", response)
}

// DeleteNftlbBackends deletes all nftlb backends from a farm given a Endpoints object.
func DeleteNftlbBackends(endpoints *v1.Endpoints) {
	farmName := endpoints.ObjectMeta.Name
	for json.GetBackendID(farmName) > 0 {
		// Makes the full path for the request
		backendName := fmt.Sprintf("%s%d", farmName, json.GetBackendID(farmName))
		fullPath := fmt.Sprintf("%s/backends/%s", farmName, backendName)
		response := deleteNftlbRequest(fullPath)
		// Prints info
		printDeleted("Backend", farmName, backendName, response)
		// Decreases backend ID by 1
		json.DecreaseBackendID(farmName)
	}
}

func deleteNftlbRequest(name string) string {
	// Makes the farm path
	farmPath := fmt.Sprintf("/%s", name)
	// Makes the URL and its Header
	farmURL := defaults.SetNftlbURL(farmPath)
	nftlbKey := defaults.SetNftlbKey()
	// Fills the request
	rq := &types.Request{
		Header: nftlbKey,
		Action: types.DELETE,
		URL:    farmURL,
	}
	// Returns the response
	return request.GetResponse(rq)
}

func printDeleted(object string, farmName string, backendName string, response string) {
	var message string
	switch object {
	case "Farm":
		message = fmt.Sprintf("\nDeleted %s name: %s\n%s", object, farmName, response)
	case "Backend":
		message = fmt.Sprintf("\nDeleted %s:\nFarm: %s, Backend:%s\n%s", object, farmName, backendName, response)
	default:
		err := fmt.Sprintf("Unknown deleted object of type %s", object)
		panic(err)
	}
	fmt.Println(message)
}
