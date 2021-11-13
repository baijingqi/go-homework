第九周作业
1、总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用
    fix length : 固定不超过缓冲区的长度，比如缓冲区每次读1M的数据，则发包不应超过1M
    delimiter based : 以分隔符为边界，一旦读取到该分隔符，意味着该包数据结束
    length field based : 包头添加长度信息，服务端解析长度参数后决定读取多少字节的数据
    
2、 目录下的a,b文件是我定义的测试文件，main.go启动后，运行client.go会将文件中的信息发送到服务端