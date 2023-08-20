# StarRandom

## 项目信息

此项目目前运行在:[https://random.geministar.site/](https://random.geministar.site/)

一个简易的随机数生成平台

可以自定义随机数的上限和下限、数量以及是否可重.

## 项目介绍

### 前端

位于[/webapp](https://github.com/Chinese-Gemini-Star/StarRandom/tree/main/webapp)目录下

采用bootstrap5框架和jQuery,通过AJAX技术与后端通信并更新网页.

### 后端

位于[/server](https://github.com/Chinese-Gemini-Star/StarRandom/tree/main/server)目录下

采用iris框架,并尝试着使用了MVC设计模式.

## 更新日志

- v1.1
  - 尝试修复了数值溢出产生的bug
  - 对输入内容不为整数时的情况作出了说明
- v1.2
  - 彻底修复了各种不合法输入导致的问题
  - 服务器端新增表单验证部分,防止api调用时上传异常的请求参数

## 更多功能

- [ ] 管理员可以发布自定义规则的随机数生成界面,以方便各种活动中随机数的生成
- [ ] 随机数存储到数据库中,防止用户篡改前端界面来糊弄他人

