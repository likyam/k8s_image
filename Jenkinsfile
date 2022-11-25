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
        ORDERSERVICE_DOCKERFILE = "$WORKSPACE/OrderService/Dockerfile"
        USERSERVICE_DOCKERFILE = "$WORKSPACE/UserService/Dockerfile"

        //yaml 文件地址
        YAML_FILE = "$WORKSPACE/service/permission/devops/dev/dev.yaml"

        //镜像名称
        ORDERAPI_DOCKER_IMAGE_NAME = "order-api:$DEV_IMAGE_VERSION"
        ORDERSERVICE_DOCKER_IMAGE_NAME = "order-service:$DEV_IMAGE_VERSION"
        USERSERVICE_DOCKER_IMAGE_NAME = "user-service:$DEV_IMAGE_VERSION"

        //镜像名称
        ORDERAPI_REGISTRY_IMAGE_NAME = "registry.cn-guangzhou.aliyuncs.com/likyam_docker/order-api:$DOCKER_IMAGE_NAME"
        ORDERSERVICE_REGISTRY_IMAGE_NAME = "registry.cn-guangzhou.aliyuncs.com/likyam_docker/order-server:$DOCKER_IMAGE_NAME"
        USERSERVICE_REGISTRY_IMAGE_NAME = "registry.cn-guangzhou.aliyuncs.com/likyam_docker/user-server:$DOCKER_IMAGE_NAME"

    }

    stages{
        //拉取项目代码
        stage('Pull Code'){
            steps{

                echo '>>>>>>>>>>>>>>>>>>START PULL CODE<<<<<<<<<<<<<<<<<<<<'

                checkout([$class: 'GitSCM', branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'ssh://git@github.com:likyam/k8s_image.git']]])

                echo '>>>>>>>>>>>>>>>>>>END PULL CODE <<<<<<<<<<<<<<<<<<<<'

            }
        }
        //构建镜像
        stage('Build Image'){
            steps{

                echo '>>>>>>>>>>>>>>>>>>START BUILD IMAGE<<<<<<<<<<<<<<<<<<<<'

                sh "docker build -t $ORDERAPI_DOCKER_IMAGE_NAME -f $ORDERAPI_DOCKERFILE ."
                sh "docker build -t $ORDERSERVICE_DOCKER_IMAGE_NAME -f $ORDERSERVICE_DOCKERFILE ."
                sh "docker build -t $USERSERVICE_DOCKER_IMAGE_NAME -f $USERSERVICE_DOCKERFILE ."

                echo '>>>>>>>>>>>>>>>>>>START BUILD IMAGE<<<<<<<<<<<<<<<<<<<<'
            }
        }
        //推送镜像
        stage('Push Image'){
            steps{

                echo '>>>>>>>>>>>>>>>>>>START PUSH IMAGE<<<<<<<<<<<<<<<<<<<<'

                //账号密码脱敏
                //withCredentials([usernamePassword(credentialsId: 'harbor_secret', passwordVariable: 'password', usernameVariable: 'username')]) {
                //    sh "echo $password | docker login https://$REGISTRY_URL -u $username --password-stdin"
                //}

                sh "docker login --username=qq13044226657 --password=qq87708291. registry.cn-guangzhou.aliyuncs.com"

                sh "docker tag $ORDERAPI_DOCKER_IMAGE_NAME $ORDERAPI_REGISTRY_IMAGE_NAME"
                sh "docker push $ORDERAPI_REGISTRY_IMAGE_NAME"

                sh "docker tag $ORDERSERVICE_DOCKER_IMAGE_NAME $ORDERSERVICE_REGISTRY_IMAGE_NAME"
                sh "docker push $ORDERSERVICE_REGISTRY_IMAGE_NAME"

                sh "docker tag $USERSERVICE_DOCKER_IMAGE_NAME $USERSERVICE_REGISTRY_IMAGE_NAME"
                sh "docker push $USERSERVICE_REGISTRY_IMAGE_NAME"

                echo '>>>>>>>>>>>>>>>>>>END PUSH IMAGE<<<<<<<<<<<<<<<<<<<<'
            }
        }
        //删除镜像
        stage('Delete Image'){
            steps{

                echo '>>>>>>>>>>>>>>>>>>START DELETE IMAGE<<<<<<<<<<<<<<<<<<<<'

                //删除本地镜像
                //sh "docker rmi janrs-permission-api:$DEV_IMAGE_VERSION"
                //sh "docker rmi 172.16.222.250:8443/janrs-io/janrs-permission-api:$DEV_IMAGE_VERSION"
                sh "docker rmi $ORDERAPI_DOCKER_IMAGE_NAME"
                sh "docker rmi $USERSERVICE_DOCKER_IMAGE_NAME"
                sh "docker rmi $USERSERVICE_DOCKER_IMAGE_NAME"

                //删除缓存的中间镜像。如果是第一次部署项目，可先将该命令注释。
                //后面再打开该参数。不删除的话，会一直增加中间镜像，占用磁盘空间。
                sh "docker image prune -f --filter \"until=10m\""

                echo '>>>>>>>>>>>>>>>>>>END DELETE IMAGE<<<<<<<<<<<<<<<<<<<<'

            }
        }
        //部署到 kubernetes
//         stage('Deploy to k8s'){
//
//             steps{
//
//                 echo '>>>>>>>>>>>>>>>>>>START DEPLOY<<<<<<<<<<<<<<<<<<<<'
//
//                 sh "sed -i 's/DEV_IMAGE_VERSION/$DEV_IMAGE_VERSION/g' $YAML_FILE"
//                 sh "kubectl apply -f $YAML_FILE --record"
//
//                 echo '>>>>>>>>>>>>>>>>>>END DEPLOY<<<<<<<<<<<<<<<<<<<<'
//
//             }
//         }
    }
}
