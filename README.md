# GoWebServer
簡易的WebServer框架


### version
顯示基本的html，輸入網址會尋找對應的html並顯示  
http://localhost/a.html  
http://localhost/b.html  

可以透過str參數替換html內的內容
http://localhost/a.html?str=123


## 目標 MVC 簡易框架
- [x] 規劃controller model view位置 (`v0.0.0`)  
    資料夾 controller, model, view  

- [x] router連結controller model view (`v0.0.1`)  
    controller對應controller/*.go，action對應其函數  
    在controller/內建立controller並在handle()裡註冊  
    測試: http://localhost/guest/index  
    測試: http://localhost/member/index  
    測試: http://localhost/1234  
  

- [] 參數傳遞  

- [] 測試session是否通用  

- [] 共用header與footer  
