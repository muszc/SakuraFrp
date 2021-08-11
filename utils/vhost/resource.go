// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io/ioutil"
	"net/http"

	frpLog "github.com/fatedier/frp/utils/log"
	"github.com/fatedier/frp/utils/version"
)

var (
	ServiceUnavailablePagePath = ""
)

const (
	ServiceUnavailable = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
 <title>GUGU FRP服务错误</title>
<link rel="shortcut icon" href="https://www.musz.cn/wp-content/LOGO/logo.png">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta content="Sample content" name="description" />
  <meta content="sample-author" name="author" />
  <link href="https://frp.musz.cn/pages/css/dmaku.css" rel="stylesheet" type="text/css" />
</head>
<body id="body-main">
  <div class="pi-wrapper">
    <div class="main-wrapper">
      <div id="particles-js" class="canvas-maker1 zindex1"></div>
      <section>
        <div class="spinner">
          <div class="loader">
            <div class="bounce loader1 "></div>
            <div class="bounce loader2 "></div>
          </div>
        </div>
      </section>
      <div class="account-pages"></div>
      <div class="clearfix"></div>
      <div class="wrapper-page relative zindex2">
        <div class="ex-page-content text-center">
          <div class="text-error floating">5 0 3</div>
          <h3 class="text-uppercase font-600 text-white">503 Service Unavailable</h3>
			<p class="text-white">您访问的网站或服务暂时不可用</p>
			<p class="text-white">如果您是隧道所有者，造成无法访问的原因可能有：</p>
				<p class="text-white">您访问的网站使用了内网穿透，但是对应的客户端没有运行。</p>
				<p class="text-white">该网站或隧道已被管理员临时或永久禁止连接。</p>
				<p class="text-white">域名解析更改还未生效或解析错误，请检查设置是否正确。</p>
			<p class="text-white">如果您是普通访问者，您可以：</p>

				<p class="text-white">稍等一段时间后再次尝试访问此站点。</p>
				<p class="text-white">尝试与该网站的所有者取得联系。</p>
				<p class="text-white">刷新您的 DNS 缓存或在其他网络环境访问。</p>
          <div class="text-center">
            <a class="btn  btn-warning btn-rounded  m-t-20 m-b-30" href="https://www.musz.cn"> <i class="fa fa-long-arrow-left"></i> 返回冰糖橙之家 </a>
          </div>
          <p class="text-white">© 2019-2021. All Rights Reserved By: <strong> <a style="color:#fff" href="https://www.musz.cn">GUGU FRP</a></strong></p>
        </div>
      </div>
    </div>
  </div>
  <script  src="https://frp.musz.cn/pages/js/jquery.min.js"></script>
  <script  src="https://frp.musz.cn/pages/js/particles.js"></script>
  <script  src="https://frp.musz.cn/pages/js/particlesapp_bubble.js"></script>
  <script  src="https://frp.musz.cn/pages/js/jquery.app.js"></script>
</body>
</html>
`
)

func getServiceUnavailablePageContent() []byte {
	var (
		buf []byte
		err error
	)
	if ServiceUnavailablePagePath != "" {
		buf, err = ioutil.ReadFile(ServiceUnavailablePagePath)
		if err != nil {
			frpLog.Warn("read custom 503 page error: %v", err)
			buf = []byte(ServiceUnavailable)
		}
	} else {
		buf = []byte(ServiceUnavailable)
	}
	return buf
}

func notFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full()+"-sakurapanel")
	header.Set("Content-Type", "text/html")

	res := &http.Response{
		Status:     "Service Unavailable",
		StatusCode: 503,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     header,
		Body:       ioutil.NopCloser(bytes.NewReader(getServiceUnavailablePageContent())),
	}
	return res
}

func noAuthResponse() *http.Response {
	header := make(map[string][]string)
	header["WWW-Authenticate"] = []string{`Basic realm="Restricted"`}
	res := &http.Response{
		Status:     "401 Not authorized",
		StatusCode: 401,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
	}
	return res
}
