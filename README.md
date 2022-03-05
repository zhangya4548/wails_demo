### 文档
### 命令行窗口必须先确定当前node版本与npm版本
```
 "node": ">=15.0.0",
 "npm": ">=7.0.0"
```    
### 安装前端插件先到JS文件夹
```
   cd ~/www/go/wails_demo/frontend/JS
☁  npm install element-plus --save
```
### 本地启动程序并监听(go文件有改动会自动重启)
```
    wails dev
```