package main

import (
	"fmt"
	"strings"
)

func main() {
	//	content := `#ONESTaskLocationStart
	//    {{if ne .OnesTaskBaseURL "" }}
	//    location ~* (^/taskapp/|/taskapp/api/) {
	//        proxy_set_header Host $host;
	//        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	//        proxy_pass {{.OnesTaskBaseURL}};
	//        proxy_http_version 1.1;
	//        proxy_set_header Upgrade $http_upgrade;
	//        proxy_set_header Connection "upgrade";
	//        proxy_ssl_verify off;
	//        add_header "Cache-Control" $cacheable_types;
	//
	//        set $env production;
	//        access_by_lua_file $lua_scripts_position/api_access.lua;
	//        access_log /data/log/nginx/ones_task.ones.ai_access.log main;
	//        error_log /data/log/nginx/ones_task.ones.ai_error.log;
	//    }
	//    {{ end }}
	//    #ONESTaskLocationEnd`
	//
	//	pa := `#ONESTaskLocationStart(([\s\S]*))#ONESTaskLocationEnd`
	//
	//	s := fmt.Sprintf(`
	//#ONESTaskLocationStart
	//location ~* (^/taskapp/|/taskapp/api/) {
	//        #proxy_set_header X-Real-IP $$remote_addr;
	//        proxy_set_header Host $$host;
	//        proxy_set_header X-Forwarded-For $$proxy_add_x_forwarded_for;
	//        proxy_pass %s;
	//        proxy_http_version 1.1;
	//        proxy_set_header Upgrade $$http_upgrade;
	//        proxy_set_header Connection "upgrade";
	//        proxy_ssl_verify off;
	//        add_header "Cache-Control" $$cacheable_types;
	//
	//        set $$env production;
	//        access_by_lua_file $$lua_scripts_position/api_access.lua;
	//        access_log /data/log/nginx/ones_task.ones.ai_access.log main;
	//        error_log /data/log/nginx/ones_task.ones.ai_error.log;
	//    }
	//#ONESTaskLocationEnd`, "onesTaskBaseURL")
	//
	//	fmt.Println(s)
	//	b := []byte(s)
	//	fmt.Println(string(b))
	//
	//	var re *regexp.Regexp
	//	re, _ = regexp.Compile(pa)
	//
	//	result := re.ReplaceAll([]byte(content), []byte(s))
	//	fmt.Println(string(result))
	//
	//

	str := "http://localhost:8888/"
	fmt.Println(strings.TrimSuffix(str, "/"))
}
