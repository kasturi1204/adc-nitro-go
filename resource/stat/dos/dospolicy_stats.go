/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package dos

/**
* Statistics for DoS policy resource.
*/

type Dospolicystats struct {
	/**
	* The name of the DoS protection policy whose statistics must be displayed. If a name is not provided, statistics of all the DoS protection policies are displayed.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Doscurrentcltdetectrate float64 `json:"doscurrentcltdetectrate,omitempty"`
	Dosphysicalserviceip string `json:"dosphysicalserviceip,omitempty"`
	Dosphysicalserviceport int32 `json:"dosphysicalserviceport,omitempty"`
	Doscurrentqueuesize uint32 `json:"doscurrentqueuesize,omitempty"`
	/**
	* Current queue size of the server to which this policy is bound.
	*/
	Doscurrentqueuesizerate int32 `json:"doscurrentqueuesizerate,omitempty"`
	Dostotjssent uint64 `json:"dostotjssent,omitempty"`
	/**
	* Total number of DoS JavaScript transactions performed for this policy.
	*/
	Dosjssentrate int32 `json:"dosjssentrate,omitempty"`
	Dostotjsrefused uint64 `json:"dostotjsrefused,omitempty"`
	/**
	* Number of times the DoS JavaScript was not sent because the set JavaScript rate was not met for this policy.
	*/
	Dosjsrefusedrate int32 `json:"dosjsrefusedrate,omitempty"`
	Dostotvalidclients uint64 `json:"dostotvalidclients,omitempty"`
	/**
	* Total number of valid DoS cookies received for this policy.
	*/
	Dosvalidclientsrate int32 `json:"dosvalidclientsrate,omitempty"`
	Dostotjsbytessent uint64 `json:"dostotjsbytessent,omitempty"`
	/**
	* Total number of DoS JavaScript bytes sent for this policy.
	*/
	Dosjsbytessentrate int32 `json:"dosjsbytessentrate,omitempty"`
	Dostotnongetpostrequests uint64 `json:"dostotnongetpostrequests,omitempty"`
	/**
	* Number of non-GET and non-POST requests for which DOS JavaScript was sent.
	*/
	Dosnongetpostrequestsrate int32 `json:"dosnongetpostrequestsrate,omitempty"`

}