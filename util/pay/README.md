# 支付

## 介绍

已经支持的部分

```
	Pay_Wechat_Mini  PayType = 11 // 小程序支付
	Pay_Wechat_App   PayType = 12 // app支付
	Pay_Wechat_JsApi PayType = 13 // jsapi 支付

	Pay_Ali_Wap  PayType = 21 // 手机网站支付
	Pay_Ali_App  PayType = 22 // app支付
	Pay_Ali_Page PayType = 22 // 电脑网站支付
```



本集成支持的方法

```
下单付款/支付
退款
支付查询
```



其他

```
如果需要使用支付宝和微信更多功能，参考：https://github.com/go-pay/gopay，自行扩展

本次集成，只集成公共和常用的部分，即：支付，退款，查询
```



## 使用

使用参考 test下的各个测试方法



## 如何扩展

如果要扩展支付宝或者微信的方法，在client模块下进行方法的追加

比如增加了 查询微信的退款方法，在 client/wechat里增加方法，如何进行调用

