daemon 使程序在后台执行
除了可以 nohup commond > /dev/null 2>&1 & 这种方式启动后台程序外 还可以直接引入下面的 package 使程序直接支持 daemon 模式。

import _ "github.com/genter/x/daemon"
假设编译后程序包为 command ，直接执行 command -d=true 程序就进入后台模式了