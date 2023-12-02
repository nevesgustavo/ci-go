# ci-go
This is a very small golang appliation and the ideia is to have a proof of concept on how to use Github Actions to:
- Test the application
- Validate code coverage using SonarQube
- Build a Docker image
- Push a Docker image to the registry

## Step by Step
**Generate a DockerHub token**
- Go to dockerhub and click on your account -> settings -> security -> create token
![image](https://github.com/nevesgustavo/ci-go/assets/10534124/0cceb3c8-1c9f-4a81-be90-a9328e404780)

**After generate the token copy the value and create secrets on github**
- Enter the repository and go to Settings -> Secrets and Variables -> Actions
- Create secrets `DOCKERHUB_TOKEN` and `DOCKERHUB_USERNAME`



  


