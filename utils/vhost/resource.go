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
<!DOCTYPE html>
<html lang="zh-cn" >

<head>
  <meta charset="UTF-8">
  <title>GUGU FRP服务错误页</title>

  
      <link rel="stylesheet" href="https://frp.musz.cn/pages/style/style.css">

  
</head>

<body>

  
<svg width="380px" height="500px" viewBox="0 0 837 1045" version="1.1" xmlns="https://www.musz.cn/wp-content/LOGO/logo.png" xmlns:xlink="https://www.musz.cn/wp-content/LOGO/logo.png" xmlns:sketch="https://www.musz.cn/wp-content/LOGO/logo.png">
    <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd" sketch:type="MSPage">
        <path d="M353,9 L626.664028,170 L626.664028,487 L353,642 L79.3359724,487 L79.3359724,170 L353,9 Z" id="Polygon-1" stroke="#007FB2" stroke-width="6" sketch:type="MSShapeGroup"></path>
        <path d="M78.5,529 L147,569.186414 L147,648.311216 L78.5,687 L10,648.311216 L10,569.186414 L78.5,529 Z" id="Polygon-2" stroke="#EF4A5B" stroke-width="6" sketch:type="MSShapeGroup"></path>
        <path d="M773,186 L827,217.538705 L827,279.636651 L773,310 L719,279.636651 L719,217.538705 L773,186 Z" id="Polygon-3" stroke="#795D9C" stroke-width="6" sketch:type="MSShapeGroup"></path>
        <path d="M639,529 L773,607.846761 L773,763.091627 L639,839 L505,763.091627 L505,607.846761 L639,529 Z" id="Polygon-4" stroke="#F2773F" stroke-width="6" sketch:type="MSShapeGroup"></path>
        <path d="M281,801 L383,861.025276 L383,979.21169 L281,1037 L179,979.21169 L179,861.025276 L281,801 Z" id="Polygon-5" stroke="#36B455" stroke-width="6" sketch:type="MSShapeGroup"></path>
    </g>
</svg>
<div class="message-box">
        <h1>503</h1>
        <h2>503 Service Unavailable</h2>
        <div class="box">
          <p>您访问的网站或服务暂时不可用</p>
          <p>如果您是隧道所有者，造成无法访问的原因可能有：</p>
          <ul>
            <li>您访问的网站使用了内网穿透，但是对应的客户端没有运行。</li>
            <li>该网站或隧道已被管理员临时或永久禁止连接。</li>
            <li>域名解析更改还未生效或解析错误，请检查设置是否正确。</li>
          </ul>
          <p>如果您是普通访问者，您可以：</p>
          <ul>
            <li>稍等一段时间后再次尝试访问此站点。</li>
            <li>尝试与该网站的所有者取得联系。</li>
            <li>刷新您的 DNS 缓存或在其他网络环境访问。</li>
          </ul>
          <p align="right"><em>Powered by GUGU FRP | Based on Frp</em></p>        
  <div class="buttons-con">
    <div class="action-link-wrap">
      <a onclick="history.back(-1)" class="link-button link-back-button">返回上一页</a>
      <a href="https://www.musz.cn/" class="link-button">返回冰糖橙之家</a>
    </div>
  </div>
</div>
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
