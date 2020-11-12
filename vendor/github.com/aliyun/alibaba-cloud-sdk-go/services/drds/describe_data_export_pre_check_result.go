package drds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDataExportPreCheckResult invokes the drds.DescribeDataExportPreCheckResult API synchronously
func (client *Client) DescribeDataExportPreCheckResult(request *DescribeDataExportPreCheckResultRequest) (response *DescribeDataExportPreCheckResultResponse, err error) {
	response = CreateDescribeDataExportPreCheckResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataExportPreCheckResultWithChan invokes the drds.DescribeDataExportPreCheckResult API asynchronously
func (client *Client) DescribeDataExportPreCheckResultWithChan(request *DescribeDataExportPreCheckResultRequest) (<-chan *DescribeDataExportPreCheckResultResponse, <-chan error) {
	responseChan := make(chan *DescribeDataExportPreCheckResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataExportPreCheckResult(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDataExportPreCheckResultWithCallback invokes the drds.DescribeDataExportPreCheckResult API asynchronously
func (client *Client) DescribeDataExportPreCheckResultWithCallback(request *DescribeDataExportPreCheckResultRequest, callback func(response *DescribeDataExportPreCheckResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataExportPreCheckResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataExportPreCheckResult(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDataExportPreCheckResultRequest is the request struct for api DescribeDataExportPreCheckResult
type DescribeDataExportPreCheckResultRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeDataExportPreCheckResultResponse is the response struct for api DescribeDataExportPreCheckResult
type DescribeDataExportPreCheckResultResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	PreCheckResult PreCheckResult `json:"PreCheckResult" xml:"PreCheckResult"`
}

// CreateDescribeDataExportPreCheckResultRequest creates a request to invoke DescribeDataExportPreCheckResult API
func CreateDescribeDataExportPreCheckResultRequest() (request *DescribeDataExportPreCheckResultRequest) {
	request = &DescribeDataExportPreCheckResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDataExportPreCheckResult", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataExportPreCheckResultResponse creates a response to parse from DescribeDataExportPreCheckResult response
func CreateDescribeDataExportPreCheckResultResponse() (response *DescribeDataExportPreCheckResultResponse) {
	response = &DescribeDataExportPreCheckResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
