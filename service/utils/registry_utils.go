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
package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func BuildRegistryUrl(segments ...string) string {
	registryURL := os.Getenv("REGISTRY_URL")
	if registryURL == "" {
		registryURL = "http://localhost:5000"
	}
	url := registryURL + "/v2"
	for _, s := range segments {
		if s == "v2" {
			log.Printf("unnecessary v2 in %v", segments)
			continue
		}
		url += "/" + s
	}
	return url
}

func RegistryApiGet(url, username string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return result, nil
	} else if response.StatusCode == http.StatusUnauthorized {
		authenticate := response.Header.Get("WWW-Authenticate")
		str := strings.Split(authenticate, " ")[1]
		log.Println("url: " + url)
		var service string
		var scope string
		strs := strings.Split(str, ",")
		for _, s := range strs {
			if strings.Contains(s, "service") {
				service = s
			} else if strings.Contains(s, "scope") {
				scope = s
			}
		}
		service = strings.Split(service, "\"")[1]
		scope = strings.Split(scope, "\"")[1]
		token, err := GenTokenForUI(username, service, scope)
		if err != nil {
			return nil, err
		}
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Add("Authorization", "Bearer "+token)
		client := &http.Client{}
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			//	log.Printf("via length: %d\n", len(via))
			if len(via) >= 10 {
				return fmt.Errorf("too many redirects")
			}
			for k, v := range via[0].Header {
				if _, ok := req.Header[k]; !ok {
					req.Header[k] = v
				}
			}
			return nil
		}
		response, err = client.Do(request)
		if err != nil {
			return nil, err
		}
		if response.StatusCode != http.StatusOK {
			errMsg := fmt.Sprintf("Unexpected return code from registry: %d", response.StatusCode)
			log.Printf(errMsg)
			return nil, fmt.Errorf(errMsg)
		}
		result, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		return result, nil
	} else {
		return nil, errors.New(string(result))
	}
}
