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

package cache

/**
* Binding class showing the cachepolicy that can be bound to cachepolicylabel.
*/
type Cachepolicylabelcachepolicybinding struct {
	/**
	* Name of the cache policy to bind to the policy label.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next-lower priority.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of policy label to invoke: an unnamed label associated with a virtual server, or user-defined policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the policy label to invoke if the current policy rule evaluates to TRUE.
	*/
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	/**
	* Name of the cache policy label to which to bind the policy.
	*/
	Labelname string `json:"labelname,omitempty"`


}