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
package controllers

import (
	"net/http"
	"net/url"
	"os"

	"github.com/vmware/harbor/dao"
	"github.com/vmware/harbor/models"

	"github.com/astaxie/beego"
)

type ItemDetailController struct {
	BaseController
}

func (idc *ItemDetailController) Get() {

	projectId, _ := idc.GetInt64("project_id")
	if projectId <= 0 {
		beego.Error("Invalid project id:", projectId)
		idc.Redirect("/signIn", http.StatusFound)
		return
	}

	project, err := dao.GetProjectById(projectId)

	if err != nil {
		beego.Error("Error occurred in GetProjectById:", err)
		idc.CustomAbort(http.StatusInternalServerError, "Internal error.")
	}

	if project == nil {
		idc.Redirect("/signIn", http.StatusFound)
		return
	}

	sessionUserId := idc.GetSession("userId")

	if project.Public != 1 && sessionUserId == nil {
		idc.Redirect("signIn?uri="+url.QueryEscape(idc.Ctx.Input.URI()), http.StatusFound)
		return
	}

	if sessionUserId != nil {
		idc.Data["Username"] = idc.GetSession("username")
		idc.Data["UserId"] = sessionUserId.(int)

		roleList, err := dao.GetUserProjectRoles(models.User{UserId: sessionUserId.(int)}, projectId)
		if err != nil {
			beego.Error("Error occurred in GetUserProjectRoles:", err)
			idc.CustomAbort(http.StatusInternalServerError, "Internal error.")
		}
		if project.Public == 0 && len(roleList) == 0 {
			idc.Redirect("registry/project", http.StatusFound)
			return
		}
		if len(roleList) > 0 {
			idc.Data["RoleId"] = roleList[0].RoleId
		}
	}

	idc.Data["ProjectId"] = project.ProjectId
	idc.Data["ProjectName"] = project.Name
	idc.Data["OwnerName"] = project.OwnerName
	idc.Data["OwnerId"] = project.OwnerId

	idc.Data["HarborRegUrl"] = os.Getenv("HARBOR_REG_URL")
	idc.Data["RepoName"] = idc.GetString("repo_name")

	idc.ForwardTo("page_title_item_details", "item-detail")

}
