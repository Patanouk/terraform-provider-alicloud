package ecs

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

// AttachKeyPair invokes the ecs.AttachKeyPair API synchronously
func (client *Client) AttachKeyPair(request *AttachKeyPairRequest) (response *AttachKeyPairResponse, err error) {
	response = CreateAttachKeyPairResponse()
	err = client.DoAction(request, response)
	return
}

// AttachKeyPairWithChan invokes the ecs.AttachKeyPair API asynchronously
func (client *Client) AttachKeyPairWithChan(request *AttachKeyPairRequest) (<-chan *AttachKeyPairResponse, <-chan error) {
	responseChan := make(chan *AttachKeyPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachKeyPair(request)
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

// AttachKeyPairWithCallback invokes the ecs.AttachKeyPair API asynchronously
func (client *Client) AttachKeyPairWithCallback(request *AttachKeyPairRequest, callback func(response *AttachKeyPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachKeyPairResponse
		var err error
		defer close(result)
		response, err = client.AttachKeyPair(request)
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

// AttachKeyPairRequest is the request struct for api AttachKeyPair
type AttachKeyPairRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	KeyPairName          string           `position:"Query" name:"KeyPairName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceIds          string           `position:"Query" name:"InstanceIds"`
}

// AttachKeyPairResponse is the response struct for api AttachKeyPair
type AttachKeyPairResponse struct {
	*responses.BaseResponse
	RequestId   string                 `json:"RequestId" xml:"RequestId"`
	TotalCount  string                 `json:"TotalCount" xml:"TotalCount"`
	FailCount   string                 `json:"FailCount" xml:"FailCount"`
	KeyPairName string                 `json:"KeyPairName" xml:"KeyPairName"`
	Results     ResultsInAttachKeyPair `json:"Results" xml:"Results"`
}

// CreateAttachKeyPairRequest creates a request to invoke AttachKeyPair API
func CreateAttachKeyPairRequest() (request *AttachKeyPairRequest) {
	request = &AttachKeyPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AttachKeyPair", "", "")
	request.Method = requests.POST
	return
}

// CreateAttachKeyPairResponse creates a response to parse from AttachKeyPair response
func CreateAttachKeyPairResponse() (response *AttachKeyPairResponse) {
	response = &AttachKeyPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
