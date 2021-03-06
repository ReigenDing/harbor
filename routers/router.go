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
package routers

import (
	"github.com/vmware/harbor/api"
	"github.com/vmware/harbor/controllers"
	"github.com/vmware/harbor/service"

	"github.com/astaxie/beego"
)

func init() {

	beego.SetStaticPath("registry/static/i18n", "static/i18n")
	beego.SetStaticPath("registry/static/resources", "static/resources")
	beego.SetStaticPath("registry/static/vendors", "static/vendors")

	beego.Router("/login", &controllers.CommonController{}, "post:Login")
	beego.Router("/logout", &controllers.CommonController{}, "get:Logout")
	beego.Router("/language", &controllers.CommonController{}, "get:SwitchLanguage")
	beego.Router("/signUp", &controllers.CommonController{}, "post:SignUp")
	beego.Router("/userExists", &controllers.CommonController{}, "post:UserExists")
	beego.Router("/reset", &controllers.CommonController{}, "post:ResetPassword")
	beego.Router("/sendEmail", &controllers.CommonController{}, "get:SendEmail")
	beego.Router("/updatePassword", &controllers.CommonController{}, "post:UpdatePassword")

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/signIn", &controllers.SignInController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/forgotPassword", &controllers.ForgotPasswordController{})
	beego.Router("/resetPassword", &controllers.ResetPasswordController{})
	beego.Router("/changePassword", &controllers.ChangePasswordController{})

	beego.Router("/registry/project", &controllers.ProjectController{})
	beego.Router("/registry/detail", &controllers.ItemDetailController{})

	beego.Router("/search", &controllers.SearchController{})

	//API:
	beego.Router("/api/search", &api.SearchAPI{})
	beego.Router("/api/projects/:pid/members/?:mid", &api.ProjectMemberAPI{})
	beego.Router("/api/projects/?:id", &api.ProjectAPI{})
	beego.Router("/api/projects/:id/logs/filter", &api.ProjectAPI{}, "post:FilterAccessLog")
	beego.Router("/api/users/?:id", &api.UserAPI{})
	beego.Router("/api/repositories", &api.RepositoryAPI{})
	beego.Router("/api/repositories/tags", &api.RepositoryAPI{}, "get:GetTags")
	beego.Router("/api/repositories/manifests", &api.RepositoryAPI{}, "get:GetManifests")

	//external service that hosted on harbor process:
	beego.Router("/service/notifications", &service.NotificationHandler{})
	beego.Router("/service/token", &service.AuthController{}, "get:Auth")
}
