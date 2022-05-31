# todo
a todo list web application

# 接口
## 注册
```
curl -v -X POST -H "Content-Type:application/json" -d '{"user_name":"robin_zhang", "password": "123456"}' localhost:6061/api/v1/user/register
```
## 登陆
```
```

# 遇到的问题
## 1、c.ShouldBind()函数的机制是什么样的？
当我使用以下命令验证`user/register`接口时报错：
```
curl -X POST -d '{"user_name":"robin_zhang", "password": "123456"}' localhost:6061/api/v1/user/register | python3 -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   357  100   308  100    49   150k  24500 --:--:-- --:--:-- --:--:--  174k
{
    "status": 40001,
    "data": null,
    "msg": "Field.%!s(func() string=0x9b6860)Tag.Valid.%!s(func() string=0x9b68c0)",
    "error": "Key: 'UserService.UserName' Error:Field validation for 'UserName' failed on the 'required' tag\nKey: 'UserService.Password' Error:Field validation for 'Password' failed on the 'required' tag"
}
```
但是奇怪的是，这里发送的参数也是按照接口定义的UserService中定义来传递的，为什么一直失败呢？

最终排查出来原因是curl命令没有指定`Content-Type`

## 2、jwt是怎么使用的？
在登陆的时候登陆成功后在body中返回token，由后续请求将token做为header提交
