# Go语言常用功能类库
> Welcome!

## 功能列表：

### 数据结构(collections/***)
+ map
    + 链式map(linkedhashmap: collections/map2/linkedhashmap.go)
+ queue
    + 优先级队列(collections/queue/priority_queue.go)
+ set
    + set(collections/set/set.go)
+ lambda&streaming(collections/stream/stream.go)
    + 函数式streaming方式处理切片数据，支持Filter,Map,Foreach,Count,Distinct,Collect等基础操作

### 文件操作(utils/file/fileutil.go)
+ 校验类
    + 校验文件路径是否安全
    + 校验文件是否存在
    + 校验目录是否存在
    + 是否是目录
+ 读取
    + 获取目录中的文件(支持递归获取目录中的所有文件)
    + 读取文件所有行

### 归档（压缩，解压缩）
+ 压缩(utils/archive/compress.go)
    + zip方式压缩文件
    + zip方式压缩目录
    + tar方式归档文件
    + tar方式归档目录
    + gzip方式归档文件
+ 解压
    + 解压zip文件
    + 解压gzip文件
    + 解压tar.gz文件
    + 解归档tar文件

### 配置(utils/config/config.go)
+ 加载配置