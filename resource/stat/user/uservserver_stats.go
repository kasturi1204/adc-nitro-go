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

package user

/**
* Statistics for virtual server resource.
*/

type Uservserverstats struct {
	/**
	* Name of the user defined virtual server. If no name is provided, statistical data of all configured user defined virtual servers is displayed.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	/**
	* use this argument to sort by specific key
	*/
	Sortby string `json:"sortby,omitempty"`
	/**
	* use this argument to specify sort order
	*/
	Sortorder string `json:"sortorder,omitempty"`
	Establishedconn uint32 `json:"establishedconn,omitempty"`
	Inactsvcs uint64 `json:"inactsvcs,omitempty"`
	Vslbhealth uint32 `json:"vslbhealth,omitempty"`
	Primaryipaddress string `json:"primaryipaddress,omitempty"`
	Primaryport int32 `json:"primaryport,omitempty"`
	Protocolname string `json:"protocolname,omitempty"`
	State string `json:"state,omitempty"`
	Actsvcs uint64 `json:"actsvcs,omitempty"`
	Tothits uint64 `json:"tothits,omitempty"`
	/**
	* Total vserver hits
	*/
	Hitsrate int32 `json:"hitsrate,omitempty"`
	Totalrequests uint64 `json:"totalrequests,omitempty"`
	/**
	* Total number of requests received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Requestsrate int32 `json:"requestsrate,omitempty"`
	Totalresponses uint64 `json:"totalresponses,omitempty"`
	/**
	* Number of responses received on this service or virtual server. (This applies to HTTP/SSL services and servers.)
	*/
	Responsesrate int32 `json:"responsesrate,omitempty"`
	Totalrequestbytes uint64 `json:"totalrequestbytes,omitempty"`
	/**
	* Total number of request bytes received on this service or virtual server.
	*/
	Requestbytesrate int32 `json:"requestbytesrate,omitempty"`
	Totalresponsebytes uint64 `json:"totalresponsebytes,omitempty"`
	/**
	* Number of response bytes received by this service or virtual server.
	*/
	Responsebytesrate int32 `json:"responsebytesrate,omitempty"`
	Totalpktsrecvd uint64 `json:"totalpktsrecvd,omitempty"`
	/**
	* Total number of packets received by this service or virtual server.
	*/
	Pktsrecvdrate int32 `json:"pktsrecvdrate,omitempty"`
	Totalpktssent uint64 `json:"totalpktssent,omitempty"`
	/**
	* Total number of packets sent.
	*/
	Pktssentrate int32 `json:"pktssentrate,omitempty"`
	Curclntconnections uint32 `json:"curclntconnections,omitempty"`
	Cursrvrconnections uint32 `json:"cursrvrconnections,omitempty"`
	Invalidrequestresponse uint64 `json:"invalidrequestresponse,omitempty"`
	Invalidrequestresponsedropped uint64 `json:"invalidrequestresponsedropped,omitempty"`

}