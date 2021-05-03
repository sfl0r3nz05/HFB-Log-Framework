import os
import asyncio
from hfc.fabric import Client

loop = asyncio.get_event_loop()

cli = Client(net_profile="../connection-profile/network.json")
org1_admin = cli.get_user('org1.example.com', 'Admin')
org2_admin = cli.get_user('org2.example.com', 'Admin')

# Make the client know there is a channel in the network
cli.new_channel('modbuschannel')

# Install Example Chaincode to Peers
# GOPATH setting is only needed to use the example chaincode inside sdk
gopath_bak = os.environ.get('GOPATH', '')
gopath = os.path.normpath(os.path.join(
    os.path.dirname(os.path.realpath('__file__')),
    '../chaincode'
))
os.environ['GOPATH'] = os.path.abspath(gopath)

#   #   #Query a chaincode
#   args = ['b']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org1.example.com'],
#       args=args,
#       cc_name='usecase_cc',
#       fcn="get"
#   ))
#   print("response", response)

#Query a chaincode
args = ['e6f8849a-9077-4f0f-ab94-51da5cb94ef2']
# The response should be true if succeed
response = loop.run_until_complete(cli.chaincode_query(
    requestor=org1_admin,
    channel_name='modbuschannel',
    peers=['peer0.org1.example.com'],
    args=args,
    cc_name='base_cc',
    fcn="get"
))
print("response", response)

#   #   #Query a chaincode
args = ['04e3fe3d-de42-4c06-87fb-a5e172d8ad16']
# The response should be true if succeed
response = loop.run_until_complete(cli.chaincode_query(
    requestor=org1_admin,
    channel_name='modbuschannel',
    peers=['peer0.org1.example.com'],
    args=args,
    cc_name='base_cc',
    fcn="get"
))
print("response", response)
   
#   #   #Query a chaincode
args = ['35f684e0-982d-4503-8e1f-4095ff856114']
# The response should be true if succeed
response = loop.run_until_complete(cli.chaincode_query(
    requestor=org1_admin,
    channel_name='modbuschannel',
    peers=['peer0.org1.example.com'],
    args=args,
    cc_name='base_cc',
    fcn="get"
))
print("response", response)
   
#   #   #   #Query a chaincode
args = ['dd03c36e-be8c-4e28-9bd0-478fe62f3ee4']
# The response should be true if succeed
response = loop.run_until_complete(cli.chaincode_query(
    requestor=org1_admin,
    channel_name='modbuschannel',
    peers=['peer0.org1.example.com'],
    args=args,
    cc_name='base_cc',
    fcn="get"
))
print("response", response)