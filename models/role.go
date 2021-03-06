/*
    Copyright (c) 2016 VMware, Inc. All Rights Reserved.
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
        
        http://www.apache.org/licenses/LICENSE-2.0
        
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
package models

const (
	SYSADMIN     = 1
	PROJECTADMIN = 2
	DEVELOPER    = 3
	GUEST        = 4
)

type Role struct {
	RoleId   int    `json:"role_id"`
	RoleCode string `json:"role_code"`
	Name     string `json:"role_name"`
}
