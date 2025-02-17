CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ignite chain build -o $HOME/Desktop/bbbbbb/

cp $HOME/go/bin/blogd $HOME/Desktop/bbbbbb/appd
    
docker buildx build --platform linux/amd64 -t 847647377987.dkr.ecr.ap-southeast-1.amazonaws.com/new-app:latest .


docker tag app:latest 847647377987.dkr.ecr.ap-southeast-1.amazonaws.com/new-app:latest
docker push 847647377987.dkr.ecr.ap-southeast-1.amazonaws.com/new-app:latest   

helm upgrade --install app ./app --namespace ignite   