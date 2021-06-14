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

package ns

/**
* Statistics for simple ACL resource.
*/

type Nssimpleaclstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Sacltothits uint64 `json:"sacltothits,omitempty"`
	/**
	* Packets matching a SimpleACL.
	*/
	Saclhitsrate int32 `json:"saclhitsrate,omitempty"`
	Sacltotmisses uint64 `json:"sacltotmisses,omitempty"`
	/**
	* Packets not matching any SimpleACL.
	*/
	Saclmissesrate int32 `json:"saclmissesrate,omitempty"`
	Saclscount uint64 `json:"saclscount,omitempty"`
	Sacltotpktsallowed uint64 `json:"sacltotpktsallowed,omitempty"`
	/**
	* Total packets that matched a SimpleACL with action ALLOW and got consumed by Citrix ADC.
	*/
	Saclpktsallowedrate int32 `json:"saclpktsallowedrate,omitempty"`
	Sacltotpktsbridged uint64 `json:"sacltotpktsbridged,omitempty"`
	/**
	* Total packets that matched a SimpleACL with action BRIDGE and got bridged by Citrix ADC.
	*/
	Saclpktsbridgedrate int32 `json:"saclpktsbridgedrate,omitempty"`
	Sacltotpktsdenied uint64 `json:"sacltotpktsdenied,omitempty"`
	/**
	* Packets dropped because they match SimpleACL (Access Control List) with processing mode set to DENY.
	*/
	Saclpktsdeniedrate int32 `json:"saclpktsdeniedrate,omitempty"`

}