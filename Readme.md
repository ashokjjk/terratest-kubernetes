## Test code for Kubernetes deployment
<b> Test your IaaC </b>
<p>Terratest is a tool from Gruntwork for Kubernetes manifest automation testing with Terraform. For more details about the tool <a href="https://terratest.gruntwork.io/">Check here</a> </p>

<b> Prerequisite </b>
1. Go compiler 1.14 or higher
2. Terraform 0.13
3. Docker for desktop
4. Kubernetes cluster enabled on Docker desktop

<b>How it works </b>
* Code directory contains Kubernetes manifest for s simple nginx pod deployment
* The deployment is expose to a service as LoadBalancer on port 8080
* Test directory contains go code to  perform integration testing on the k8s resource
* Notice the <b>name</b> of the go file. Should end with <code>_test.go</code>
* Once all the files are placed open shell or cmd navigate to test directory.
```
go mod init test
go test -v
```
### Expected results 
```
PS F:\terratest\k8s\test> go test -v                                                                                                                                           === RUN   TestK8sDeployment
=== PAUSE TestK8sDeployment
=== CONT  TestK8sDeployment
TestK8sDeployment 2020-10-18T23:37:00+05:30 logger.go:66: Running command kubectl with args [--namespace default apply -f ../code/deployment.yaml]
TestK8sDeployment 2020-10-18T23:37:00+05:30 logger.go:66: deployment.apps/ngapp created
TestK8sDeployment 2020-10-18T23:37:02+05:30 logger.go:66: service/ng-out created
TestK8sDeployment 2020-10-18T23:37:02+05:30 retry.go:72: Wait for service ng-out to be provisioned.
TestK8sDeployment 2020-10-18T23:37:02+05:30 client.go:33: Configuring kubectl using config file C:\Users\CFSDSAHP\.kube\config with context
TestK8sDeployment 2020-10-18T23:37:02+05:30 node.go:33: Getting list of nodes from Kubernetes
TestK8sDeployment 2020-10-18T23:37:02+05:30 client.go:33: Configuring kubectl using config file C:\Users\CFSDSAHP\.kube\config with context
TestK8sDeployment 2020-10-18T23:37:03+05:30 retry.go:84: Wait for service ng-out to be provisioned. returned an error: Service ng-out is not available. Sleeping for 5s and will try again.
TestK8sDeployment 2020-10-18T23:37:08+05:30 retry.go:72: Wait for service ng-out to be provisioned.
TestK8sDeployment 2020-10-18T23:37:08+05:30 client.go:33: Configuring kubectl using config file C:\Users\CFSDSAHP\.kube\config with context
TestK8sDeployment 2020-10-18T23:37:08+05:30 node.go:33: Getting list of nodes from Kubernetes
TestK8sDeployment 2020-10-18T23:37:08+05:30 client.go:33: Configuring kubectl using config file C:\Users\CFSDSAHP\.kube\config with context
TestK8sDeployment 2020-10-18T23:37:08+05:30 service.go:86: Service is now available
TestK8sDeployment 2020-10-18T23:37:08+05:30 client.go:33: Configuring kubectl using config file C:\Users\CFSDSAHP\.kube\config with context
TestK8sDeployment 2020-10-18T23:37:08+05:30 node.go:33: Getting list of nodes from Kubernetes
TestK8sDeployment 2020-10-18T23:37:08+05:30 client.go:33: Configuring kubectl using config file C:\Users\CFSDSAHP\.kube\config with context
TestK8sDeployment 2020-10-18T23:37:08+05:30 retry.go:72: HTTP GET to URL http://localhost:8080
TestK8sDeployment 2020-10-18T23:37:08+05:30 http_helper.go:32: Making an HTTP GET call to URL http://localhost:8080
TestK8sDeployment 2020-10-18T23:37:09+05:30 retry.go:84: HTTP GET to URL http://localhost:8080 returned an error: Get "http://localhost:8080": EOF. Sleeping for 5s and will try again.
TestK8sDeployment 2020-10-18T23:37:14+05:30 retry.go:72: HTTP GET to URL http://localhost:8080
TestK8sDeployment 2020-10-18T23:37:14+05:30 http_helper.go:32: Making an HTTP GET call to URL http://localhost:8080
TestK8sDeployment 2020-10-18T23:37:14+05:30 logger.go:66: Running command kubectl with args [--namespace default delete -f ../code/deployment.yaml]
TestK8sDeployment 2020-10-18T23:37:14+05:30 logger.go:66: deployment.apps "ngapp" deleted
TestK8sDeployment 2020-10-18T23:37:14+05:30 logger.go:66: service "ng-out" deleted
--- PASS: TestK8sDeployment (15.24s)
PASS
ok      test    15.307s
```

<p> Automation testing for K8S using Helm charts check <a href="https://github.com/gruntwork-io/terratest-helm-testing-example">here</a></p>
