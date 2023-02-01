# A-Simple-Tictok-Project
```
| -- constants 	静态常量&&配置信息（硬编码）
| -- errno 		统一返回错误信息
| -- handler 	视图函数
| -- middleware 中间件 
| -- pack		统一返回数据（目前未使用）
| -- repository 表结构信息
| -- router		路由信息
| -- service	业务逻辑
| -- sql
|	| -- only_struct.sql 表结构
|	| -- with_data.sql   表结构+数据
| -- utils		通用函数
```

#### 用户模块

|   描述   |          地址           | 请求方法 | 是否完成 |
| :------: | :---------------------: | :------: | :------: |
| 用户注册 | `/douyin/user/register` |   POST   |    是    |
| 用户登录 |  `/douyin/user/login`   |   POST   |    是    |
| 用户信息 |     `/douyin/user/`     |   GET    |    是    |

#### 视频模块

| 描述       | 地址                     | 请求方法 | 是否完成 |
| ---------- | ------------------------ | -------- | -------- |
| 投稿视频   | `/douyin/publish/action` | POST     | 是       |
| 发布列表   | `/douyin/publish/list`   | GET      | 是       |
| 视频流接口 | `/douyin/feed`           | GET      | 是       |

#### 点赞模块

|   描述    |           地址            | 请求方法 | 是否完成 |
| :-------: | :-----------------------: | :------: | :------: |
| 点赞/点踩 | `/douyin/favorite/action` |   POST   |    是    |
| 喜欢列表  |  `/douyin/favorite/list`  |   GET    |    是    |

#### 评论模块

|   描述    |           地址           | 请求方法 | 是否完成 |
| :-------: | :----------------------: | :------: | :------: |
| 评论/删除 | `/douyin/comment/action` |   POST   |    是    |
| 评论列表  |  `/douyin/comment/list`  |   GET    |    是    |

#### 社交模块

|   描述    |              地址              | 请求方法 | 是否完成 |
| :-------: | :----------------------------: | :------: | :------: |
| 关注/取关 |   `/douyin/relation/action`    |   POST   |    是    |
| 关注列表  | `/douyin/relation/follow/list` |   GET    |    是    |
| 粉丝列表  | `/douyin/relat/follower/list`  |   GET    |    是    |

### 项目运行步骤

1. 下载安装MySQL，创建一个数据库。**注意编码格式和排序格式**

2. 在`repository`中的`db_init`中将`dsn`信息改成自己机器上的数据库信息

3. 按照自己的需求，将`sql`文件夹中的sql文件导入数据库中。**注意必须先创建数据库再导入**

   - `only_struct.sql`只建表，没有数据

   - `with_data.sql`表和数据都会创建。数据是我的一些测试数据。**注意：用户表中的用户密码都是用户名123**

     ```
     username		password
     A-Simple-Tictok-Project			A-Simple-Tictok-Project123
     tank			tank123
     alex			alex123
     eric			eric123
     egon			egon123
     ```

4. 按照正常的编译运行main文件即可
