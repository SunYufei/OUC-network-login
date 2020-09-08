OUC实验室网络登录界面（10.81.2.6）更新后不能保存账户密码，使用此程序可实现一键登录，减少操作步骤
# 使用方法

1. 下载可执行文件 [login.exe](https://github.com/SunYufei/OUC-network-login/releases/download/v1.0/login.exe) 和运行脚本 [run.bat](https://github.com/SunYufei/OUC-network-login/releases/download/v1.0/run.bat)
2. 修改 `run.bat` 文件的内容，将 `USERNAME` 和 `PASSWORD` 替换为校园网账户密码
3. 运行 `run.bat` 即可登录

## 注意

目前仅有 64位 Windows 版本可执行程序，其他平台可自行编译 `login.go` 文件并改写 `run.bat` 文件的内容

或者使用 Python 版本

```shell script
pip3 install requests
python3 login.py -u USERNAME -p PASSWORD
```
