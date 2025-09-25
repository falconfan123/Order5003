@echo off
REM 餐厅点餐系统启动脚本（Windows）
echo 正在启动餐厅点餐系统...
echo.
start order5003.exe
echo 系统已启动！
echo.
echo 请在浏览器中访问以下地址：
echo 顾客端：http://localhost:8080
echo 商家端：http://localhost:8080/merchant
echo.
echo 商家登录账号：admin
echo 商家登录密码：password
echo.
echo 按任意键关闭此窗口...
pause > nul