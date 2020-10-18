package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestK8sDeployment(t *testing.T) {
	t.Parallel()

	resoursePath := "../code/deployment.yaml"

	// Declaring the namespace
	options := k8s.NewKubectlOptions("", "", "default")
	// Delete the deployment
	defer k8s.KubectlDelete(t, options, resoursePath)

	// Apply the deployment
	k8s.KubectlApply(t, options, resoursePath)

	// Wait for pods & service to come up
	retries := 15
	sleep := 5 * time.Second
	k8s.WaitUntilServiceAvailable(t, options, "ng-out", retries, sleep)
	// Once available get the details of service
	service := k8s.GetService(t, options, "ng-out")

	// Get the url and hit the endpoint
	url := fmt.Sprintf("http://%s", k8s.GetServiceEndpoint(t, options, service, 8080))

	// Get the data from the endpoint and match the details are correct
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		url,
		nil,
		retries,
		sleep,
		func(statusCode int, body string) bool {
			isOk := statusCode == 200
			isNginx := strings.Contains(body, "Welcome to nginx")
			return isOk && isNginx
		},
	)

}
