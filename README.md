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
