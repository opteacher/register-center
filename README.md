# register-center
Kratos微服务的注册中心，提供服务注册、注销、以及Kong网关相关的HTTP绑定

## 启动流程
1. `docker-compose up -d --build`，将会依次启动
    1. `kong-database`：Kong数据库
    2. `discovery`：Bl的服务注册中心
    3. `kong-migrations`：Kong启动时执行的某些脚本
    4. `kong`：Kong网关服务
    5. `konga`：Kong网关的GUI站点
    6. `register-center`：微服务注册中心（也是个微服务）
2. 暴露如下端口供使用：
    * `1337`：`konga`的站点端口，通过`localhost:1337`访问
    * `8001`：`kong`的管理员端口，可依照[Kong文档](https://docs.konghq.com/1.2.x/admin-api/)使用Postman操作
    * `7171`：Bl的服务发现端口，可通过`http://localhost:7171/discovery/fetch/all`访问注册中的所有微服务（以后可做GUI界面管理所有GRPC服务）
    * `9093`：注册中心微服务的GRPC端口，不过微服务的使用是通过appID调用的，一般也用不到这个端口。必要时可用过直连的方式查看注册中心是否可用
    > 直连微服务：`conn, err := cli.Dial(context.Background(), "direct://default/X.X.X.X:9093")`