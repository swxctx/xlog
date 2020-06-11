# xlog
[![Build Status](https://travis-ci.org/swxctx/xlog.svg?branch=master)](https://travis-ci.org/swxctx/xlog)
[![Go Report Card](https://goreportcard.com/badge/github.com/swxctx/xlog)](https://goreportcard.com/report/github.com/swxctx/xlog)
[![GoDoc](http://godoc.org/github.com/swxctx/xlog?status.svg)](http://godoc.org/github.com/swxctx/xlog)

Golang日志工具。

## 功能
1. 支持日志级别设置  
2. 多级别输出  

## 使用
```
package main

import(
    "github.com/swxctx/xlog"
)

funca main(){
    // 设置日志级别
    xlog.Default.Level = xlog.DebugLevel
	xlog.Infof("infof test->%d", xlog.DebugLevel)
	xlog.Debugf("debug test->%d", xlog.DebugLevel)
	xlog.Warnf("warnf test->%d", xlog.DebugLevel)
	xlog.Errorf("error test->%d", xlog.DebugLevel)
	xlog.Fatalf("fatalf test->%d", xlog.DebugLevel)
}
```  
  
效果图:
![uzthoozlogg](https://github.com/swxctx/xlog/blob/master/img/golang.png)
