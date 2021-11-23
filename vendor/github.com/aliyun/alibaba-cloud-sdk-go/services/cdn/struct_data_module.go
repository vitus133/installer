package cdn

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

// DataModule is a nested struct in cdn response
type DataModule struct {
	PeakTime              string  `json:"PeakTime" xml:"PeakTime"`
	Acc                   int64   `json:"Acc" xml:"Acc"`
	Ipv6Qps               float64 `json:"Ipv6Qps" xml:"Ipv6Qps"`
	HttpsAccOverseasValue string  `json:"HttpsAccOverseasValue" xml:"HttpsAccOverseasValue"`
	HttpsOverseasValue    string  `json:"HttpsOverseasValue" xml:"HttpsOverseasValue"`
	TotalValue            string  `json:"TotalValue" xml:"TotalValue"`
	TimeStamp             string  `json:"TimeStamp" xml:"TimeStamp"`
	Ipv6Acc               int64   `json:"Ipv6Acc" xml:"Ipv6Acc"`
	HttpsValue            string  `json:"HttpsValue" xml:"HttpsValue"`
	Traf                  int64   `json:"Traf" xml:"Traf"`
	Ipv6Bps               float64 `json:"Ipv6Bps" xml:"Ipv6Bps"`
	OverseasValue         string  `json:"OverseasValue" xml:"OverseasValue"`
	SpecialValue          string  `json:"SpecialValue" xml:"SpecialValue"`
	DomainName            string  `json:"DomainName" xml:"DomainName"`
	Ipv6Traf              int64   `json:"Ipv6Traf" xml:"Ipv6Traf"`
	Bps                   float64 `json:"Bps" xml:"Bps"`
	DomesticValue         string  `json:"DomesticValue" xml:"DomesticValue"`
	AccValue              string  `json:"AccValue" xml:"AccValue"`
	Value                 string  `json:"Value" xml:"Value"`
	Qps                   float64 `json:"Qps" xml:"Qps"`
	HttpCode              string  `json:"HttpCode" xml:"HttpCode"`
	AccDomesticValue      string  `json:"AccDomesticValue" xml:"AccDomesticValue"`
	TrafficValue          string  `json:"TrafficValue" xml:"TrafficValue"`
	HttpsDomesticValue    string  `json:"HttpsDomesticValue" xml:"HttpsDomesticValue"`
	HttpsAccValue         string  `json:"HttpsAccValue" xml:"HttpsAccValue"`
	AccOverseasValue      string  `json:"AccOverseasValue" xml:"AccOverseasValue"`
	HttpsAccDomesticValue string  `json:"HttpsAccDomesticValue" xml:"HttpsAccDomesticValue"`
}