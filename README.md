# jc_mall
this is a online mall project writing by go!

## 公告
模仿京东商城完成一个电商的项目

# 实现的接口
I、商品的接口
1.商品列表
2.商品的增加
3.商品的编辑
4.商品的删除
5.商品的详情

II、商品分类
1.获取商品所有分类
2.获取商品子分类
3.创建商品你分类
4.删除商品分类
5.更新商品分类
6.品牌列表
7.通过分类获取品牌信息
8.创建商品品牌
9.更新商品品牌信息
10.删除商品品牌信息

III、订单的接口
1.创建订单
2.修改订单
3.查看订单
4.取消订单

IV、购物车的接口
1.加入购物车
2.修改购物车
3.查看购物车
4.清空购物车

V、用户的接口
1.用户的注册
2.检查用户密码
3.通过ID查找用户
4.通过Mobile查找用户
5.更新用户信息
6.获取用户列表


# 文件目录
jc_srv
    ------goods_srv
        --- config  定义相关的配置信息
        --- proto  定义相关的数据格式
        --- handler 定义相关的接口方法
        --- model  模型文件，定义相关的数据表
        --- global 全局文件
        --- tests 测试  
    ------inventory_srv
        --- config  定义库存相关的配置信息
        --- proto  定义库存相关的数据格式
        --- handler 定义库存相关的接口方法
        --- model  模型文件，定义相关的数据表
        --- global 全局文件
        --- tests 测试 
    ------order_srv
        --- config 定义订单相关的配置信息
        --- proto  定义订单相关的数据格式
        --- handler 定义相关的接口方法
        --- model  订单模型文件，定义相关的数据表
        --- global 全局文件
        --- tests 测试 
    ------user_srv
        --- config  定义相关的配置信息
        --- proto  定义用户相关的数据格式
        --- handler 定义用户相关的接口方法
        --- model  模型文件，定义用户相关的数据表
        --- global 全局文件
        --- tests 测试 
    ------userop_srv
        --- config  定义相关的配置信息
        --- proto  定义用户配置相关的数据格式
        --- handler 定义用户配置相关的接口方法
        --- model  模型文件，定义用户配置相关的数据表
        --- global 全局文件
        --- tests 测试 
    ------go.mod   需要安装的go模块，类似于python中的requirements.txt 文件
    ------go.sum
