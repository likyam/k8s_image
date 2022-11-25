pipeline{

    agent   any

    //选项
    options {
        //设置终端显示颜色
        ansiColor('gnome-terminal')
    }

    //全局变量
    environment {
        //设置镜像版本
        //DEV_IMAGE_VERSION = "dev_${sh(script:'head -c 6 /dev/random | base64' , returnStdout: true)}"
        DEV_IMAGE_VERSION = "v1.2"

        //镜像中心地址
        REGISTRY_URL = "registry.cn-guangzhou.aliyuncs.com"

        //Dockerfile 文件地址
        ORDERAPI_DOCKERFILE = "$WORKSPACE/OrderApi/Dockerfile"
        ORDERSERVER_DOCKERFILE = "$WORKSPACE/OrderServer/Dockerfile"
        USERSERVER_DOCKERFILE = "$WORKSPACE/UserServer/Dockerfile"

        //yaml 文件地址
        YAML_FILE = "$WORKSPACE/service/permission/devops/dev/dev.yaml"

        //镜像名称
        ORDERAPI_DOCKER_IMAGE_NAME = "order-api:$DEV_IMAGE_VERSION"
        ORDERSERVER_DOCKER_IMAGE_NAME = "order-server:$DEV_IMAGE_VERSION"
        USERSERVER_DOCKER_IMAGE_NAME = "user-server:$DEV_IMAGE_VERSION"

        //镜像名称
        ORDERAPI_REGISTRY_IMAGE_NAME = "registry.cn-guangzhou.aliyuncs.com/likyam_docker/order-api:$DEV_IMAGE_VERSION"
        ORDERSERVER_REGISTRY_IMAGE_NAME = "registry.cn-guangzhou.aliyuncs.com/likyam_docker/order-server:$DEV_IMAGE_VERSION"
        USERSERVER_REGISTRY_IMAGE_NAME = "registry.cn-guangzhou.aliyuncs.com/likyam_docker/user-server:$DEV_IMAGE_VERSION"

    }

    stages{
        //构建镜像
        stage('Build Image'){
            steps{

                echo '>>>>>>>>>>>>>>>>>>START BUILD IMAGE<<<<<<<<<<<<<<<<<<<<'

                sh "docker build -t $ORDERAPI_DOCKER_IMAGE_NAME -f $ORDERAPI_DOCKERFILE ."
                sh "docker build -t $ORDERSERVER_DOCKER_IMAGE_NAME -f $ORDERSERVER_DOCKERFILE ."
                sh "docker build -t $USERSERVER_DOCKER_IMAGE_NAME -f $USERSERVER_DOCKERFILE ."

                echo '>>>>>>>>>>>>>>>>>>START BUILD IMAGE<<<<<<<<<<<<<<<<<<<<'
            }
        }
        //推送镜像
        stage('Push Image'){
            steps{

                echo '>>>>>>>>>>>>>>>>>>START PUSH IMAGE<<<<<<<<<<<<<<<<<<<<'

                sh "docker login --username=qq13044226657 --password=qq87708291. registry.cn-guangzhou.aliyuncs.com"

                sh "docker tag $ORDERAPI_DOCKER_IMAGE_NAME $ORDERAPI_REGISTRY_IMAGE_NAME"
                sh "docker push $ORDERAPI_REGISTRY_IMAGE_NAME"

                sh "docker tag $ORDERSERVER_DOCKER_IMAGE_NAME $ORDERSERVER_REGISTRY_IMAGE_NAME"
                sh "docker push $ORDERSERVER_REGISTRY_IMAGE_NAME"

                sh "docker tag $USERSERVER_DOCKER_IMAGE_NAME $USERSERVER_REGISTRY_IMAGE_NAME"
                sh "docker push $USERSERVER_REGISTRY_IMAGE_NAME"

                echo '>>>>>>>>>>>>>>>>>>END PUSH IMAGE<<<<<<<<<<<<<<<<<<<<'
            }
        }
        //删除镜像
        stage('Delete Image'){
            steps{

                echo '>>>>>>>>>>>>>>>>>>START DELETE IMAGE<<<<<<<<<<<<<<<<<<<<'

                //删除本地镜像
                sh "docker rmi $ORDERAPI_DOCKER_IMAGE_NAME"
                sh "docker rmi $ORDERSERVER_DOCKER_IMAGE_NAME"
                sh "docker rmi $USERSERVER_DOCKER_IMAGE_NAME"

                //删除缓存的中间镜像。如果是第一次部署项目，可先将该命令注释。
                //后面再打开该参数。不删除的话，会一直增加中间镜像，占用磁盘空间。
                sh "docker image prune -f --filter \"until=10m\""

                echo '>>>>>>>>>>>>>>>>>>END DELETE IMAGE<<<<<<<<<<<<<<<<<<<<'

            }
        }
    }
}
