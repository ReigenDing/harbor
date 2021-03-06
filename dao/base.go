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
	"log"
	"net"

	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const NON_EXIST_USER_ID = 0

func isIllegalLength(s string, min int, max int) bool {
	if min == -1 {
		return (len(s) > max)
	}
	if max == -1 {
		return (len(s) <= min)
	}
	return (len(s) < min || len(s) > max)
}

func isContainIllegalChar(s string, illegalChar []string) bool {
	for _, c := range illegalChar {
		if strings.Index(s, c) >= 0 {
			return true
		}
	}
	return false
}

func GenerateRandomString() (string, error) {
	o := orm.NewOrm()
	var uuid string
	err := o.Raw(`select uuid() as uuid`).QueryRow(&uuid)
	if err != nil {
		return "", err
	}
	return uuid, nil

}

func InitDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	addr := os.Getenv("MYSQL_HOST")
	if len(addr) == 0 {
		addr = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	}

	port := os.Getenv("MYSQL_PORT_3306_TCP_PORT")
	username := os.Getenv("MYSQL_USER")

	password := os.Getenv("MYSQL_ENV_MYSQL_ROOT_PASSWORD")
	if len(password) == 0 {
		password = os.Getenv("MYSQL_PWD")
	}

	var flag bool = true
	if addr == "" {
		beego.Error("Unset env of MYSQL_HOST")
		flag = false
	} else if port == "" {
		beego.Error("Unset env of MYSQL_PORT_3306_TCP_PORT")
		flag = false
	} else if username == "" {
		beego.Error("Unset env of MYSQL_USR")
		flag = false
	} else if password == "" {
		beego.Error("Unset env of MYSQL_PWD")
		flag = false
	}

	if !flag {
		os.Exit(1)
	}

	db_str := username + ":" + password + "@tcp(" + addr + ":" + port + ")/registry"
	ch := make(chan int, 1)
	go func() {
		var err error
		var c net.Conn
		for {
			c, err = net.Dial("tcp", "127.0.0.1:3306")
			if err == nil {
				c.Close()
				ch <- 1
			} else {
				log.Printf("failed to connect to db, retry after 2 seconds...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	select {
	case <-ch:
	case <-time.After(60 * time.Second):
		panic("Failed to connect to DB after 60 seconds")
	}
	err := orm.RegisterDataBase("default", "mysql", db_str)
	if err != nil {
		panic(err)
	}
}
