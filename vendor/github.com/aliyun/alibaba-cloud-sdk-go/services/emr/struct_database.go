package emr

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

// Database is a nested struct in emr response
type Database struct {
	DataSourceId       string `json:"DataSourceId" xml:"DataSourceId"`
	DataSourceType     string `json:"DataSourceType" xml:"DataSourceType"`
	Owner              string `json:"Owner" xml:"Owner"`
	Location           string `json:"Location" xml:"Location"`
	GmtCreate          int64  `json:"GmtCreate" xml:"GmtCreate"`
	LocationType       string `json:"LocationType" xml:"LocationType"`
	OwnerType          string `json:"OwnerType" xml:"OwnerType"`
	DatabaseType       string `json:"DatabaseType" xml:"DatabaseType"`
	ProjectName        string `json:"ProjectName" xml:"ProjectName"`
	GmtModified        int64  `json:"GmtModified" xml:"GmtModified"`
	ProjectId          string `json:"ProjectId" xml:"ProjectId"`
	ClusterBizId       string `json:"ClusterBizId" xml:"ClusterBizId"`
	ClusterName        string `json:"ClusterName" xml:"ClusterName"`
	Id                 string `json:"Id" xml:"Id"`
	Status             string `json:"Status" xml:"Status"`
	DatabaseParameters string `json:"DatabaseParameters" xml:"DatabaseParameters"`
	DatabaseName       string `json:"DatabaseName" xml:"DatabaseName"`
	DatabaseComment    string `json:"DatabaseComment" xml:"DatabaseComment"`
}