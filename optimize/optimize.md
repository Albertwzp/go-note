[TOC]
---
# reduce size
go build -ldflags="-s -w" //-s：忽略符号表和调试信息  -w：忽略DWARFv3调试信息，使用该选项后将无法使用gdb进行调试
go build -ldflags="-s -w" -o server main.go && upx -9 server //压缩[1-9]，比例从大到小
go build -gcflags="-d=ssa/check_bce/debug=1" 查看是否进行边界检查
