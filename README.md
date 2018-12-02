# Go Micoservice Template


### Description:
Minimal working template for creating and deploying a go micro-service in both local and lambda http / json


### Structure:
* **cmd**
  * commands which run the micro-service using either local http or lambda http networking
* **networking**
  * transport layer which receives a request and write an appropriate response, as specified by the business logic in the service; this request is treated as either a http request or a generic apit gateway request, depending on whether or not the http or lambda transport layer is used; in both cases the response is marshalled via an json encoder
* **service**
  * mock service which returns an appropriate message given a name string in the url query params


### Requirements:
Install both **go** and **dep**
(Example: ` brew install golang, dep `)


### Local Usage:
1. Remember to import this repo into your gopath: https://golang.org/doc/code.html
2. Run ` make build ` to build the binary from the go project
3. Start the server: `./bin/run_http_server ` and test with a local client GET request
4. (Example: http://localhost:8080/process_name/gary)


### Serverless Deployment:
1. Install serverless if necessary: ` npm install serverless -g `
2. Login to serverless: ` sls login `
3. Add an application through the serverless website
4. Remember to add your app and tenant name to *./serverless.yml*: https://serverless.com/framework/docs/providers/aws/guide/deploying/
5. Setup an appropriate AWS profile: https://serverless.com/framework/docs/providers/aws/guide/credentials/
6. Set "PROFILE_NAME" in "./Makefile"
7. Run ` make deploy `
8. (Bonus: Print logging to console: ` sls logs -f lambdahttpserver `)
9. (Bonus: Remove serverless deployment: ` make remove `)
