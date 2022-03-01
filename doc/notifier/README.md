# 介绍

`通知器`主要用来在适当的时候通知调用方（或者其它想被通知的目标），而这个适当的时机包括

- 执行出错
- 执行完成

而`通知器`的作用大概有

- 程序内通知，使用方法通知器
- Slack通知
- 企业微信通知
- QQ通知
- 邮件通知
- 短信通知

::: warning 注意

上述通知器，因为不想Gex代码库变得太大，都不作实现，只实现的应用内方法通知，其它通知器以另外的库来作支持
:::

`Gex`内置了常见通知器，包括

- 应用内方法通知`Func`

## 使用

所有的通知器都是使用以`Notifier`结尾的方法来做配置，包括

- `Func``