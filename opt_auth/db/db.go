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
package db

import (
	"fmt"

	"github.com/vmware/harbor/dao"
	"github.com/vmware/harbor/models"
	"github.com/vmware/harbor/opt_auth"
)

type DbAuth struct{}

func (d *DbAuth) Validate(auth models.AuthModel) (*models.User, error) {
	u, err := dao.LoginByDb(auth)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func init() {
	fmt.Println("opt db init")
	opt_auth.Register("db_auth", &DbAuth{})
}
