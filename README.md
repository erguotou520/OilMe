# 油价提醒
使用`golang`查询每日油价(使用`showapi`接口)，当油价调整时自动给微信发送信息(使用`server酱`提供的服务)

## 环境变量
- API_URL `showapi`的api地址，使用简单模式`secret`发送，形如`https://route.showapi.com/138-46?prov=你所在的城市&showapi_appid=your_id&showapi_sign=your_secret`
- SERVER_URL `server酱`的微信通知地址，形如`https://sc.ftqq.com/YOUR_SCKEY.send`

## 使用方法
- 使用docker部署
  ```bash
  docker run -d -v /data/logs:/go/src/oilme/logs -e API_URL="https://route.showapi.com/138-46?prov=你所在的城市&showapi_appid=your_id&showapi_sign=your_secret" -e SERVER_URL="https://sc.ftqq.com/YOUR_SCKEY.send" --name oilme erguotou/oilme
  ```
- 使用`golang`运行
  ```bash
  # 先clone项目，然后进入项目目录
  API_URL=https://route.showapi.com/138-46?prov=你所在的城市&showapi_appid=your_id&showapi_sign=your_secret SERVER_URL=https://sc.ftqq.com/YOUR_SCKEY.send go run bee.go
  ```
