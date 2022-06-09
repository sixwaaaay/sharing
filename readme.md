# 服务端

- common
  通用的逻辑
  - auth
  jwt 签发及其测试
  - cotypes
  公用的传输类型
  - errorx
  错误类型定义
  - covert
  类型转化辅助函数及测试
  - middleware
  http中间件，处理token，鉴权等及测试
  - secu
  密码加密辅助函数及测试
  - testhelper
  测试辅助函数及其测试
- app
  - basic
  基础的应用逻辑
  - comment
  评论接口相关逻辑
  - feed
  feed流逻辑
  - relation
  关注关系相关接口
  - user
  用户信息相关接口
  - favorite
  点赞相关接口

    以点赞为例子，内部代码组织
    - dal
      数据库操作相关逻辑
    - handler
      处理点赞的handler
    - logic
      点赞相关接口的逻辑

## 数据储存

### 数据库设计

参考deploy目录下的sql文件



## 用户服务

路由前缀：/douyin/user

### 对密码的保护

- 数据库中的储存

  采取加盐值的方式，对密码进行加密，避免密码在数据库中存储的明文形式。

- 传输过程的安全保障

  通过HTTPS自带的安全功能，对传输过程进行加密，防止传输过程中的数据被窃取。

### 登录状态的保持

通过 bearer token 机制。

通过hmac算法，对token进行加密，防止token被篡改。

> 参考：[JWT](https://jwt.io/)

### 注册

请求类型：POST
路由：/register
请求参数如下，属于query参数：
username: 用户名，必须，最长不超过32个字符
password: 密码，必须，最长不超过32个字符

## 错误处理

在 common 目录下的 errorx 包中，定义了包装了错误信息的结构体，并且定义了一些默认的错误码。
同时也提供了自定义错误码的方式。
