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
package dao

import (
	"github.com/vmware/harbor/models"

	"github.com/astaxie/beego/orm"
)

func GetUserProjectRoles(userQuery models.User, projectId int64) ([]models.Role, error) {

	o := orm.NewOrm()

	sql := `select distinct r.role_id, r.role_code, r.name 
		from role r 
		left join project_role pr on r.role_id = pr.role_id
		left join user_project_role upr on pr.pr_id = upr.pr_id
		left join user u on u.user_id = upr.user_id
		where u.deleted = 0 
		  and u.user_id = ? `
	queryParam := make([]interface{}, 1)
	queryParam = append(queryParam, userQuery.UserId)

	if projectId > 0 {
		sql += ` and pr.project_id = ? `
		queryParam = append(queryParam, projectId)
	}
	if userQuery.RoleId > 0 {
		sql += ` and r.role_id = ? `
		queryParam = append(queryParam, userQuery.RoleId)
	}

	var roleList []models.Role
	_, err := o.Raw(sql, queryParam).QueryRows(&roleList)

	if err != nil {
		return nil, err
	}
	return roleList, nil
}

func IsAdminRole(userId int) (bool, error) {
	//role_id == 1 means the user is system admin
	userQuery := models.User{UserId: userId, RoleId: models.SYSADMIN}
	adminRoleList, err := GetUserProjectRoles(userQuery, 0)
	if err != nil {
		return false, err
	}
	return len(adminRoleList) > 0, nil
}
