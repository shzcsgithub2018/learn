# 选择基础镜像，这里以Ubuntu为例
FROM ubuntu:24.10

# 安装必要的工具
RUN apt-get update && apt-get install -y curl wget gnupg

# 安装Golang
RUN wget -q -O - https://go.dev/dl/go1.23.5.linux-amd64.tar.gz | tar -C /usr/local -xzf -
ENV PATH=$PATH:/usr/local/go/bin

# 安装Python
RUN apt-get install -y python3 python3-pip

# 安装Java
RUN apt-get install -y default-jdk

# 设置工作目录
WORKDIR /app

# 可以在这里添加复制应用代码等相关操作，比如COPY. /app

# 定义容器启动时的命令，这里只是示例，可根据实际情况修改
CMD ["bash"]