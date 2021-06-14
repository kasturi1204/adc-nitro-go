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

package appfw


type Appfwstats struct {
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Appfirewalltotallog uint64 `json:"appfirewalltotallog,omitempty"`
	/**
	* Total number of security check log messages generated by the Application Firewall.
	*/
	Appfirewalllograte int32 `json:"appfirewalllograte,omitempty"`
	Appfirewalltotalviol uint64 `json:"appfirewalltotalviol,omitempty"`
	/**
	* Total number of security check violations seen by the Application Firewall.
	*/
	Appfirewallviolrate int32 `json:"appfirewallviolrate,omitempty"`
	Appfirewallshortavgresptime uint64 `json:"appfirewallshortavgresptime,omitempty"`
	Appfirewalllongavgresptime uint64 `json:"appfirewalllongavgresptime,omitempty"`
	Appfirewallrequests uint64 `json:"appfirewallrequests,omitempty"`
	/**
	* HTTP/HTTPS requests sent to your protected web servers via the Application Firewall.
	*/
	Appfirewallrequestsrate int32 `json:"appfirewallrequestsrate,omitempty"`
	Appfirewallreqbytes uint64 `json:"appfirewallreqbytes,omitempty"`
	/**
	* Number of bytes transfered for requests
	*/
	Appfirewallreqbytesrate int32 `json:"appfirewallreqbytesrate,omitempty"`
	Appfirewallresponses []uint64 `json:"appfirewallresponses,omitempty"`
	/**
	* HTTP/HTTPS responses sent by your protected web servers via the Application Firewall.
	*/
	Appfirewallresponsesrate int32 `json:"appfirewallresponsesrate,omitempty"`
	Appfirewallresbytes uint64 `json:"appfirewallresbytes,omitempty"`
	/**
	* Number of bytes transfered for responses
	*/
	Appfirewallresbytesrate int32 `json:"appfirewallresbytesrate,omitempty"`
	Appfirewallaborts uint64 `json:"appfirewallaborts,omitempty"`
	/**
	* Incomplete HTTP/HTTPS requests aborted by the client before the Application Firewall could finish processing them.
	*/
	Appfirewallabortsrate int32 `json:"appfirewallabortsrate,omitempty"`
	Appfirewallredirects uint64 `json:"appfirewallredirects,omitempty"`
	/**
	* HTTP/HTTPS requests redirected by the Application Firewall to a different Web page or web server. (HTTP 302)
	*/
	Appfirewallredirectsrate int32 `json:"appfirewallredirectsrate,omitempty"`
	Appfirewalltrapsdropped uint64 `json:"appfirewalltrapsdropped,omitempty"`
	Appfirewallviolstarturl uint64 `json:"appfirewallviolstarturl,omitempty"`
	/**
	* Number of Start URL security check violations seen by the Application Firewall.
	*/
	Appfirewallviolstarturlrate int32 `json:"appfirewallviolstarturlrate,omitempty"`
	Appfirewallvioldenyurl uint64 `json:"appfirewallvioldenyurl,omitempty"`
	/**
	* Number of Deny URL security check violations seen by the Application Firewall.
	*/
	Appfirewallvioldenyurlrate int32 `json:"appfirewallvioldenyurlrate,omitempty"`
	Appfirewallviolrefererheader uint64 `json:"appfirewallviolrefererheader,omitempty"`
	/**
	* Number of Referer Header security check violations seen by the Application Firewall.
	*/
	Appfirewallviolrefererheaderrate int32 `json:"appfirewallviolrefererheaderrate,omitempty"`
	Appfirewallviolbufferoverflow uint64 `json:"appfirewallviolbufferoverflow,omitempty"`
	/**
	* Number of Buffer Overflow security check violations seen by the Application Firewall.
	*/
	Appfirewallviolbufferoverflowrate int32 `json:"appfirewallviolbufferoverflowrate,omitempty"`
	Appfirewallpostbodylimitviolations uint64 `json:"appfirewallpostbodylimitviolations,omitempty"`
	/**
	* Number of Post Body Limit security check violations seen by the Application Firewall.
	*/
	Appfirewallpostbodylimitviolationsrate int32 `json:"appfirewallpostbodylimitviolationsrate,omitempty"`
	Appfirewallviolcookie uint64 `json:"appfirewallviolcookie,omitempty"`
	/**
	* Number of Cookie Consistency security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcookierate int32 `json:"appfirewallviolcookierate,omitempty"`
	Appfirewallviolcookiehijack uint64 `json:"appfirewallviolcookiehijack,omitempty"`
	/**
	* Number of Cookie Hijacking security violations seen by the Application Firewall.
	*/
	Appfirewallviolcookiehijackrate int32 `json:"appfirewallviolcookiehijackrate,omitempty"`
	Appfirewallviolcsrftag uint64 `json:"appfirewallviolcsrftag,omitempty"`
	/**
	* Number of Cross Site Request Forgery form tag security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcsrftagrate int32 `json:"appfirewallviolcsrftagrate,omitempty"`
	Appfirewallviolxss uint64 `json:"appfirewallviolxss,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxssrate int32 `json:"appfirewallviolxssrate,omitempty"`
	Appfirewallviolsql uint64 `json:"appfirewallviolsql,omitempty"`
	/**
	* Number of HTML SQL Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallviolsqlrate int32 `json:"appfirewallviolsqlrate,omitempty"`
	Appfirewallviolfieldformat uint64 `json:"appfirewallviolfieldformat,omitempty"`
	/**
	* Number of Field Format security check violations seen by the Application Firewall.
	*/
	Appfirewallviolfieldformatrate int32 `json:"appfirewallviolfieldformatrate,omitempty"`
	Appfirewallviolfieldconsistency uint64 `json:"appfirewallviolfieldconsistency,omitempty"`
	/**
	* Number of Field Consistency security check violations seen by the Application Firewall.
	*/
	Appfirewallviolfieldconsistencyrate int32 `json:"appfirewallviolfieldconsistencyrate,omitempty"`
	Appfirewallviolfileuploadtypes uint64 `json:"appfirewallviolfileuploadtypes,omitempty"`
	/**
	* Number of Field Upload Types security check violations seen by the Application Firewall.
	*/
	Appfirewallviolfileuploadtypesrate int32 `json:"appfirewallviolfileuploadtypesrate,omitempty"`
	Appfirewallviolxmlpayloadcontenttypemismatch uint64 `json:"appfirewallviolxmlpayloadcontenttypemismatch,omitempty"`
	/**
	* Number of Mismatched Content-Type in request with XML Payload security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxmlpayloadcontenttypemismatchrate int32 `json:"appfirewallviolxmlpayloadcontenttypemismatchrate,omitempty"`
	Appfirewallviolcreditcard uint64 `json:"appfirewallviolcreditcard,omitempty"`
	/**
	* Number of Credit Card security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcreditcardrate int32 `json:"appfirewallviolcreditcardrate,omitempty"`
	Appfirewallviolsafeobject uint64 `json:"appfirewallviolsafeobject,omitempty"`
	/**
	* Number of Safe Object security check violations seen by the Application Firewall.
	*/
	Appfirewallviolsafeobjectrate int32 `json:"appfirewallviolsafeobjectrate,omitempty"`
	Appfirewallviolsignature uint64 `json:"appfirewallviolsignature,omitempty"`
	/**
	* Number of Signature violations seen by the Application Firewall.
	*/
	Appfirewallviolsignaturerate int32 `json:"appfirewallviolsignaturerate,omitempty"`
	Appfirewallviolcontenttype uint64 `json:"appfirewallviolcontenttype,omitempty"`
	/**
	* Number of Content type security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcontenttyperate int32 `json:"appfirewallviolcontenttyperate,omitempty"`
	Appfirewallviolcmd uint64 `json:"appfirewallviolcmd,omitempty"`
	/**
	* Number of HTML CMD Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallviolcmdrate int32 `json:"appfirewallviolcmdrate,omitempty"`
	Appfirewallvioljsondos uint64 `json:"appfirewallvioljsondos,omitempty"`
	/**
	* Number of JSON Denial-of-Service security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsondosrate int32 `json:"appfirewallvioljsondosrate,omitempty"`
	Appfirewallvioljsonsql uint64 `json:"appfirewallvioljsonsql,omitempty"`
	/**
	* Number of JSON SQL Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsonsqlrate int32 `json:"appfirewallvioljsonsqlrate,omitempty"`
	Appfirewallvioljsonxss uint64 `json:"appfirewallvioljsonxss,omitempty"`
	/**
	* Number of JSON Cross-Site Scripting (XSS) security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsonxssrate int32 `json:"appfirewallvioljsonxssrate,omitempty"`
	Appfirewallvioljsoncmd uint64 `json:"appfirewallvioljsoncmd,omitempty"`
	/**
	* Number of JSON Command Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallvioljsoncmdrate int32 `json:"appfirewallvioljsoncmdrate,omitempty"`
	Appfirewallviolwellformednessviolations uint64 `json:"appfirewallviolwellformednessviolations,omitempty"`
	/**
	* Number of XML Format security check violations seen by the Application Firewall.
	*/
	Appfirewallviolwellformednessviolationsrate int32 `json:"appfirewallviolwellformednessviolationsrate,omitempty"`
	Appfirewallviolxdosviolations uint64 `json:"appfirewallviolxdosviolations,omitempty"`
	/**
	* Number of XML Denial-of-Service security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxdosviolationsrate int32 `json:"appfirewallviolxdosviolationsrate,omitempty"`
	Appfirewallviolmsgvalviolations uint64 `json:"appfirewallviolmsgvalviolations,omitempty"`
	/**
	* Number of XML Message Validation security check violations seen by the Application Firewall.
	*/
	Appfirewallviolmsgvalviolationsrate int32 `json:"appfirewallviolmsgvalviolationsrate,omitempty"`
	Appfirewallviolwsiviolations uint64 `json:"appfirewallviolwsiviolations,omitempty"`
	/**
	* Number of Web Services Interoperability (WS-I) security check violations seen by the Application Firewall.
	*/
	Appfirewallviolwsiviolationsrate int32 `json:"appfirewallviolwsiviolationsrate,omitempty"`
	Appfirewallviolxmlsqlviolations uint64 `json:"appfirewallviolxmlsqlviolations,omitempty"`
	/**
	* Number of XML SQL Injection security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxmlsqlviolationsrate int32 `json:"appfirewallviolxmlsqlviolationsrate,omitempty"`
	Appfirewallviolxmlxssviolations uint64 `json:"appfirewallviolxmlxssviolations,omitempty"`
	/**
	* Number of XML Cross-Site Scripting (XSS) security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxmlxssviolationsrate int32 `json:"appfirewallviolxmlxssviolationsrate,omitempty"`
	Appfirewallviolxmlattachmentviolations uint64 `json:"appfirewallviolxmlattachmentviolations,omitempty"`
	/**
	* Number of XML Attachment security check violations seen by the Application Firewall.
	*/
	Appfirewallviolxmlattachmentviolationsrate int32 `json:"appfirewallviolxmlattachmentviolationsrate,omitempty"`
	Appfirewallviolxmlsoapfaultviolations uint64 `json:"appfirewallviolxmlsoapfaultviolations,omitempty"`
	/**
	* Number of requests returning soap:fault from the backend server
	*/
	Appfirewallviolxmlsoapfaultviolationsrate int32 `json:"appfirewallviolxmlsoapfaultviolationsrate,omitempty"`
	Appfirewallviolxmlgenviolations uint64 `json:"appfirewallviolxmlgenviolations,omitempty"`
	/**
	* Number of requests returning XML generic error from the backend server
	*/
	Appfirewallviolxmlgenviolationsrate int32 `json:"appfirewallviolxmlgenviolationsrate,omitempty"`
	Appfirewallviolsqlgram uint64 `json:"appfirewallviolsqlgram,omitempty"`
	/**
	* Number of HTML SQL Injection security check violations (using SQL grammar) seen by the Application Firewall.
	*/
	Appfirewallviolsqlgramrate int32 `json:"appfirewallviolsqlgramrate,omitempty"`
	Appfirewallvioljsonsqlgram uint64 `json:"appfirewallvioljsonsqlgram,omitempty"`
	/**
	* Number of JSON SQL Injection security check violations (reported using SQL grammar) seen by the Application Firewall.
	*/
	Appfirewallvioljsonsqlgramrate int32 `json:"appfirewallvioljsonsqlgramrate,omitempty"`
	Appfirewalllogstarturl uint64 `json:"appfirewalllogstarturl,omitempty"`
	/**
	* Number of Start URL security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogstarturlrate int32 `json:"appfirewalllogstarturlrate,omitempty"`
	Appfirewalllogdenyurl uint64 `json:"appfirewalllogdenyurl,omitempty"`
	/**
	* Number of Deny URL security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogdenyurlrate int32 `json:"appfirewalllogdenyurlrate,omitempty"`
	Appfirewalllogrefererheader uint64 `json:"appfirewalllogrefererheader,omitempty"`
	/**
	* Number of Referer Header security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogrefererheaderrate int32 `json:"appfirewalllogrefererheaderrate,omitempty"`
	Appfirewalllogbufferoverflow uint64 `json:"appfirewalllogbufferoverflow,omitempty"`
	/**
	* Number of Buffer Overflow security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogbufferoverflowrate int32 `json:"appfirewalllogbufferoverflowrate,omitempty"`
	Appfirewallpostbodylimitlogs uint64 `json:"appfirewallpostbodylimitlogs,omitempty"`
	/**
	* Number of Post Body Limit security check logs seen by the Application Firewall.
	*/
	Appfirewallpostbodylimitlogsrate int32 `json:"appfirewallpostbodylimitlogsrate,omitempty"`
	Appfirewalllogcookie uint64 `json:"appfirewalllogcookie,omitempty"`
	/**
	* Number of Cookie Consistency security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcookierate int32 `json:"appfirewalllogcookierate,omitempty"`
	Appfirewalllogcookiehijack uint64 `json:"appfirewalllogcookiehijack,omitempty"`
	/**
	* Number of Cookie Hijacking security violation log messages generated by the Application Firewall.
	*/
	Appfirewalllogcookiehijackrate int32 `json:"appfirewalllogcookiehijackrate,omitempty"`
	Appfirewalllogcsrftag uint64 `json:"appfirewalllogcsrftag,omitempty"`
	/**
	* Number of Cross Site Request Forgery form tag security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcsrftagrate int32 `json:"appfirewalllogcsrftagrate,omitempty"`
	Appfirewalllogxss uint64 `json:"appfirewalllogxss,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogxssrate int32 `json:"appfirewalllogxssrate,omitempty"`
	Appfirewalllogtransformxss uint64 `json:"appfirewalllogtransformxss,omitempty"`
	/**
	* Number of HTML Cross-Site Scripting security check transform log messages generated by the Application Firewall.
	*/
	Appfirewalllogtransformxssrate int32 `json:"appfirewalllogtransformxssrate,omitempty"`
	Appfirewalllogsql uint64 `json:"appfirewalllogsql,omitempty"`
	/**
	* Number of HTML SQL Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsqlrate int32 `json:"appfirewalllogsqlrate,omitempty"`
	Appfirewalllogtransformsql uint64 `json:"appfirewalllogtransformsql,omitempty"`
	/**
	* Number of HTML SQL Injection security check transform log messages generated by the Application Firewall.
	*/
	Appfirewalllogtransformsqlrate int32 `json:"appfirewalllogtransformsqlrate,omitempty"`
	Appfirewalllogfieldformat uint64 `json:"appfirewalllogfieldformat,omitempty"`
	/**
	* Number of Field Format security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogfieldformatrate int32 `json:"appfirewalllogfieldformatrate,omitempty"`
	Appfirewalllogfieldconsistency uint64 `json:"appfirewalllogfieldconsistency,omitempty"`
	/**
	* Number of Field Consistency security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogfieldconsistencyrate int32 `json:"appfirewalllogfieldconsistencyrate,omitempty"`
	Appfirewalllogcreditcard uint64 `json:"appfirewalllogcreditcard,omitempty"`
	/**
	* Number of Credit Card security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcreditcardrate int32 `json:"appfirewalllogcreditcardrate,omitempty"`
	Appfirewalllogsafeobject uint64 `json:"appfirewalllogsafeobject,omitempty"`
	/**
	* Number of Safe Object security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsafeobjectrate int32 `json:"appfirewalllogsafeobjectrate,omitempty"`
	Appfirewallsignaturelogs uint64 `json:"appfirewallsignaturelogs,omitempty"`
	/**
	* Number of Signature logs generated by the Application Firewall.
	*/
	Appfirewallsignaturelogsrate int32 `json:"appfirewallsignaturelogsrate,omitempty"`
	Appfirewalllogcontenttype uint64 `json:"appfirewalllogcontenttype,omitempty"`
	/**
	* Number of content type security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcontenttyperate int32 `json:"appfirewalllogcontenttyperate,omitempty"`
	Appfirewalllogsjsondos uint64 `json:"appfirewalllogsjsondos,omitempty"`
	/**
	* Number of JSON Denial-of-Service security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsondosrate int32 `json:"appfirewalllogsjsondosrate,omitempty"`
	Appfirewalllogsjsonsql uint64 `json:"appfirewalllogsjsonsql,omitempty"`
	/**
	* Number of JSON SQL Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsonsqlrate int32 `json:"appfirewalllogsjsonsqlrate,omitempty"`
	Appfirewalllogsjsonxss uint64 `json:"appfirewalllogsjsonxss,omitempty"`
	/**
	* Number of JSON Cross-Site Scripting (XSS) security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsonxssrate int32 `json:"appfirewalllogsjsonxssrate,omitempty"`
	Appfirewalllogsjsoncmd uint64 `json:"appfirewalllogsjsoncmd,omitempty"`
	/**
	* Number of JSON Command Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogsjsoncmdrate int32 `json:"appfirewalllogsjsoncmdrate,omitempty"`
	Appfirewalllogfileuploadtypes uint64 `json:"appfirewalllogfileuploadtypes,omitempty"`
	/**
	* Number of File Upload Types security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogfileuploadtypesrate int32 `json:"appfirewalllogfileuploadtypesrate,omitempty"`
	Appfirewalllogxmlpayloadcontenttypemismatch uint64 `json:"appfirewalllogxmlpayloadcontenttypemismatch,omitempty"`
	/**
	* Number of Mismatched Content-Type in request with XML Payload security check logs seen by the Application Firewall.
	*/
	Appfirewalllogxmlpayloadcontenttypemismatchrate int32 `json:"appfirewalllogxmlpayloadcontenttypemismatchrate,omitempty"`
	Appfirewalllogcmd uint64 `json:"appfirewalllogcmd,omitempty"`
	/**
	* Number of HTML Command Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewalllogcmdrate int32 `json:"appfirewalllogcmdrate,omitempty"`
	Appfirewallwellformednesslogs uint64 `json:"appfirewallwellformednesslogs,omitempty"`
	/**
	* Number of XML Format security check log messages generated by the Application Firewall.
	*/
	Appfirewallwellformednesslogsrate int32 `json:"appfirewallwellformednesslogsrate,omitempty"`
	Appfirewallxdoslogs uint64 `json:"appfirewallxdoslogs,omitempty"`
	/**
	* Number of XML Denial-of-Service security check log messages generated by the Application Firewall.
	*/
	Appfirewallxdoslogsrate int32 `json:"appfirewallxdoslogsrate,omitempty"`
	Appfirewallmsgvallogs uint64 `json:"appfirewallmsgvallogs,omitempty"`
	/**
	* Number of XML Message Validation security check log messages generated by the Application Firewall.
	*/
	Appfirewallmsgvallogsrate int32 `json:"appfirewallmsgvallogsrate,omitempty"`
	Appfirewallwsilogs uint64 `json:"appfirewallwsilogs,omitempty"`
	/**
	* Number of Web Services Interoperability (WS-I) security check log messages generated by the Application Firewall.
	*/
	Appfirewallwsilogsrate int32 `json:"appfirewallwsilogsrate,omitempty"`
	Appfirewallxmlsqllogs uint64 `json:"appfirewallxmlsqllogs,omitempty"`
	/**
	* Number of XML SQL Injection security check log messages generated by the Application Firewall.
	*/
	Appfirewallxmlsqllogsrate int32 `json:"appfirewallxmlsqllogsrate,omitempty"`
	Appfirewallxmlxsslogs uint64 `json:"appfirewallxmlxsslogs,omitempty"`
	/**
	* Number of XML Cross-Site Scripting (XSS) security check log messages generated by the Application Firewall.
	*/
	Appfirewallxmlxsslogsrate int32 `json:"appfirewallxmlxsslogsrate,omitempty"`
	Appfirewallxmlattachmentlogs uint64 `json:"appfirewallxmlattachmentlogs,omitempty"`
	/**
	* Number of XML Attachment security check log messages generated by the Application Firewall.
	*/
	Appfirewallxmlattachmentlogsrate int32 `json:"appfirewallxmlattachmentlogsrate,omitempty"`
	Appfirewallxmlsoapfaultlogs uint64 `json:"appfirewallxmlsoapfaultlogs,omitempty"`
	/**
	* Number of requests generating soap:fault log messages
	*/
	Appfirewallxmlsoapfaultlogsrate int32 `json:"appfirewallxmlsoapfaultlogsrate,omitempty"`
	Appfirewallxmlgenlogs uint64 `json:"appfirewallxmlgenlogs,omitempty"`
	/**
	* Number of requests generating XML generic error log messages
	*/
	Appfirewallxmlgenlogsrate int32 `json:"appfirewallxmlgenlogsrate,omitempty"`
	Appfirewalllogsqlgram uint64 `json:"appfirewalllogsqlgram,omitempty"`
	/**
	* Number of HTML SQL Injection security check log messages (reported by SQL grammar) generated by the Application Firewall.
	*/
	Appfirewalllogsqlgramrate int32 `json:"appfirewalllogsqlgramrate,omitempty"`
	Appfirewalllogsjsonsqlgram uint64 `json:"appfirewalllogsjsonsqlgram,omitempty"`
	/**
	* Number of JSON SQL Injection security check log messages (reported by SQL grammar) generated by the Application Firewall.
	*/
	Appfirewalllogsjsonsqlgramrate int32 `json:"appfirewalllogsjsonsqlgramrate,omitempty"`
	Appfirewallret4xx uint64 `json:"appfirewallret4xx,omitempty"`
	/**
	* Number of requests returning HTTP 4xx from the backend server
	*/
	Appfirewallret4xxrate int32 `json:"appfirewallret4xxrate,omitempty"`
	Appfirewallret5xx uint64 `json:"appfirewallret5xx,omitempty"`
	/**
	* Number of requests returning HTTP 5xx from the backend server
	*/
	Appfirewallret5xxrate int32 `json:"appfirewallret5xxrate,omitempty"`

}