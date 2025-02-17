CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ignite chain build -o $HOME/Desktop/bbbbbb/

cp $HOME/go/bin/blogd $HOME/Desktop/bbbbbb/keplerd
    
docker buildx build --platform linux/amd64 -t 847647377987.dkr.ecr.ap-southeast-1.amazonaws.com/kepler:latest .


docker tag kepler:latest 847647377987.dkr.ecr.ap-southeast-1.amazonaws.com/kepler:latest
docker push 847647377987.dkr.ecr.ap-southeast-1.amazonaws.com/kepler:latest   

helm upgrade --install kepler ./kepler --namespace ignite   