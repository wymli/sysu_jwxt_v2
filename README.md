# sysu_jwxt_v2
# 停止维护,手贱试了下无限制的刷新获取课程列表,被拉黑名单了,还得写检讨....
---
- 当前版本
  - 适配当前\(2020.8\)对外网开放的 jwxt.sysu.edu.cn
  - 课程仅查询东校园,仅支持专选和公选
- 内外网兼容
  - 因为有些url还没有兼容,所以通过 portal.sysu.edu.cn 然后访问 jwxt-443.webvpn.sysu.edu.cn 可能会有问题,前端隐去了这一接口
    - 经测试二者不兼容,通过 jwxt.sysu.edu.cn 走 cas.sysu.edu.cn 拿到cookie后访问 portal.sysu.edu.cn 仍提示未登录
      - 虽然理论上作为sso单点登录,应该是可以共用的
    - jwxt-443.webvpn.sysu.edu.cn 似乎没有直接提供登陆接口,必须走 portal.sysu.edu.cn
---
## 页面展示
- 登陆
  - ![](Readme_staticFile/2020-08-15-14-15-30.png)
- 课程列表
  - ![](Readme_staticFile/2020-08-15-15-06-59.png)

  - ![](Readme_staticFile/2020-08-15-14-13-15.png)

  - ![](Readme_staticFile/2020-08-16-00-49-43.png)
---
## TODO:
- [ ] 改成 Lazy Load
  - 目前是鼠标hover到一行后加载那一行的教师信息和照片.
  - [x] fixme:移到下一行才刷新上一行的信息(使用this.$set())
- [x] 选课
- [x] 退课
- [ ] 自动选课
- [ ] 换课
- [ ] 登陆界面,删掉 `记住密码`
- [ ] ...
---
## BUG:
- 前端会请求 `/undefined`,原因可能是 `<img src="">` 空src属性导致[x] (只请求了一次,应该不是这个原因)
- ...
---
## 代码逻辑
- 抢课:
  - jwxt后端未提供查询单个课程的接口
  - 所以这里采用先将课程加入`收藏`,然后不断查询收藏的课程,只要有空位立即选课,然后取消收藏
---
## 运行:
1. 下载本仓库,或git clone  
2. 直接鼠标双击打开 `main.exe` (最好使用shell:`./main`)
3. 打开网页: `localhost:9999`
---
## 编译
> 不清楚会不会遇到一些动态链接出错的问题,可能会缺少运行库,建议自己编译
- 下载安装 `golang `
- `go get` 按照第三方包
  - gin : `go get github.com/gin-gonic/gin`
  - goquery : `github.com/PuerkitoBio/goquery`
- `go build` (在当前文件目录下)
- `./main`

---
### ChangeLog:
- 2020/8/15  repo created
  - 基本界面设计
  - 爬取课程列表
  - 爬取教师信息
- 2020/8/15 教师信息加载: `mounted() -> row @hover` 
- 2020/8/15 更新选课/退课前端后端接口
- 2020/8/16 完成选课/退课,课程信息加载: `mounted -> @tab-click`
- 2020/8/16 修改选课/退课后端接口,修改前端登陆页面,后端添加了课程加入收藏动作
- 2020/8/18 前端添加了定时任务(自动选课)界面
- 2020/8/24 初步完成了定时任务(抢课),目前由于未交学费,接口无法测试
- 下一步:
  - 自动选课
  - 换课
  - chooseCourse(scope.$index, scope.row , 'public') 删去参数'public',直接使用row.courseSelectedType判断
  - 增加加收藏退收藏的状态判断!查询收藏列表即可!
  - 几个按钮重新布局
  - 修复多个定时器问题