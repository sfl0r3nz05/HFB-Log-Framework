# HFB-network

## SET ENVIRONMENTAL VARIABLES

1. export FABRIC_VERSION=1.4.6
2. export FABRIC_CA_VERSION=1.4.6

## CASE SOLO 2ORG WITH GOLEVELDB

3. docker network inspect NETWORK 2org_2peer_solo_goleveldb_default
   - copy Gateway: e.g.:"Gateway": "172.18.0.1"
   - paste ip address (url): ~/HFB-network-main/connection-profile/network.json

4. Deploy HFB network

   - cd networks/2org_2peer_solo_goleveldb
   - docker-compose up -d

---

| chaincode          | channel-join | deploy-chaincode | invoke | query  |
| ------------------ | ------------ | ---------------- | ------ | ------ |
| usecase_cc         | v1           | v1               | v1     | v1     |
| ------------------ | ------------ | ---------------- | ------ | ------ |

---

5. If error "gopath not found":
   find -name gopath
      The container id is returned: 4300e823d5268d558c19370bda0ff2fa61e54c4c743133b5be58634adfbefc3b

6. go get github.com/google/uuid
   sudo cp -rf /home/ubuntu/go/src/github.com/google /var/lib/docker/overlay2/4300e823d5268d558c19370bda0ff2fa61e54c4c743133b5be58634adfbefc3b/diff/opt/gopath/src/github.com/

7. git clone https://github.com/sirupsen/logrus.git
   sudo cp -rf /home/ubuntu/go/src/github.com/sirupsen /var/lib/docker/overlay2/4300e823d5268d558c19370bda0ff2fa61e54c4c743133b5be58634adfbefc3b/diff/opt/gopath/src/github.com/

8. go get golang.org/x/sys/unix
   sudo cp -rf /home/ubuntu/go/src/golang.org /var/lib/docker/overlay2/4300e823d5268d558c19370bda0ff2fa61e54c4c743133b5be58634adfbefc3b/diff/opt/gopath/src/

9. 
   sudo cp -rf /home/ubuntu/HFB-network-main/chaincode/src/github.com/usecase_cc/log /var/lib/docker/overlay2/4300e823d5268d558c19370bda0ff2fa61e54c4c743133b5be58634adfbefc3b/diff/opt/gopath/src/github.com/