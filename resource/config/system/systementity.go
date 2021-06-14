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

package system

/**
* Configuration for entity resource.
*/
type Systementity struct {
	/**
	* Specify the entity type.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the source which contains all the stored counter values.
	*/
	Datasource string `json:"datasource,omitempty"`
	/**
	* Specify core ID of the PE in nCore.
	*/
	Core int32 `json:"core,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`

}