# LogAuHFBCH: Log Auditing emitted by HFB's Chaincode

![alt text](https://github.com/sfl0r3nz05/LogAuHFBCH/blob/main/img/System%20overview.png)

## CASE SOLO 2ORG WITH GOLEVELDB

1. Deploy HFB network and ELK infrastructure
   - cd networks/2org_2peer_solo_goleveldb
   - docker-compose up -d

2. If error "gopath not found":
   find -name gopath
      The container id is returned: 4300e823d5268d558c19370bda0ff2fa61e54c4c743133b5be58634adfbefc3b

3. go get github.com/google/uuid
   sudo cp -rf /home/ubuntu/go/src/github.com/google /var/lib/docker/overlay2/4300e823d5268d558c19370bda0ff2fa61e54c4c743133b5be58634adfbefc3b/diff/opt/gopath/src/github.com/

4. Once the network is deployed, the chaincode container identifier associated with the use case must be identified.

5. Set execute permissions to the folder associated with the container id from step 4.
   -  E.g. sudo chmod +x -R /var/lib/docker/containers/<containerid>

5. Copy the container id in .env  PATH_TO_CONTAINER
![alt text](https://github.com/sfl0r3nz05/LogAuHFBCH/blob/main/img/Path.png)

6. Delete and run logstash container to start to recive logs.
![alt text](https://github.com/sfl0r3nz05/LogAuHFBCH/blob/main/img/RemoveContainers.png)

7. Logs in logstash can be verified through logs of "logstash container": 
![alt text](https://github.com/sfl0r3nz05/LogAuHFBCH/blob/main/img/Logs.png)

### To Do:
 1. Deploy Hyperledger Explorer
 3. Include Python Script to verify each log
![alt text](https://github.com/sfl0r3nz05/LogAuHFBCH/blob/main/img/System%20overviewII.png)
 4. Include performance evaluation, e.g.:
![alt text](https://github.com/sfl0r3nz05/LogAuHFBCH/blob/main/img/performance.png) 