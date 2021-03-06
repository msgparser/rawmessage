# rawmessage
<code>rawmessage</code> is a microservice which parses raw email text message and extracts required fields. This service expects http POST payload in RFC 5322/6532 syntax.
The service is implemented in golang. It uses core golang packages usually shipped alonged with go compiler. 
 
## Directory Structure 
| Directory | Description |
| --- | --- |
| src | golang source, Dockerfiles |
| doc | build, deploy, test documents |
| smallset | Sample raw email text messages | 
| deploy | Kubernetes yamls to add resources |

## Rest API
The current version supports only one API in this microservice
| API | Description | Method | payload | output |
| --- | --- | --- | --- | --- |
| http://Server-name:4000/rawmsg | Consumer's interface | POST | application/text | Json output|

## Testing
All testing is successfull and gave expected results.
* Service is containerized on Linux server & tested successfully by smallset sample data files
* Exected on bash and tested with local port 4000 
* Service is tested successully in EKS with Route 53 subdomain and LB
![AWS LB](https://github.com/msgparser/rawmessage/blob/main/doc/AWS-LB.PNG)

# Build 
https://github.com/msgparser/rawmessage/blob/main/doc/Build%20source.md

