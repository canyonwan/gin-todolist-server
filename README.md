## 服务架构
    routes -> controller层(接收前端过来的参数) -> 交给Dao层(将参数实例化并与数据库交互) 
    -> 完成后交给service层进行一些的具体的逻辑判断并返回响应展示到前端
### Model层
    对数据库表名和表内字段进行模型定义的模块
### Controller层
    负责请求转发,接受前端转过来的参数,传给service层处理,接受到返回值并返回给前端,
    控制器本身不输出任何东西和做任何处理
### Dao层
    基于gorm对数据进行增删查改的模块,该层是数据结构与表结构一一对应的,
    其中封装了CRUD（增加Create、检索Retrieve、更新Update和删除Delete）
    基本操作，建议DAO只做原子操作，增删改查。
    常见封装的方法包括增删查改、分页查询（Paginate）、查询全部（FindAll）、
    按条件查询（WhereQuery）、关联查询（Joins）、预加载（Preloads）等等方法
### Service层
    主要负责调用多个Dao层