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

#Query a chaincode
args = ['a']
# The response should be true if succeed
response = loop.run_until_complete(cli.chaincode_query(
    requestor=org1_admin,
    channel_name='modbuschannel',
    peers=['peer0.org1.example.com'],
    args=args,
    cc_name='usecase_cc',
    fcn="get"
))
print("response", response)

#   #Query a chaincode
#   args = ['5bb9b970-4c38-4eab-b644-fe21742a0cac']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org1.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)
#   
#   #   #   #Query a chaincode
#   args = ['d91fc8a9-1a68-4299-a57a-c0f4e28d0397']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org1.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)
#      
#   #   #   #Query a chaincode
#   args = ['7e87517c-0daf-4704-885c-f571d07f3510']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org1.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)
#      
#   #   #   #   #Query a chaincode
#   args = ['e77a32c9-36ea-433b-9400-11e2ea82dc49']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org1.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)
#   
#   #   #   #   #Query a chaincode
#   args = ['1253f817-1ad1-4342-ad4e-efcc5ca9c9d5']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org1.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)