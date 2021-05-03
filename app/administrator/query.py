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
args = ['b']
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
#   args = ['e7a989ba-0d36-4434-8f2d-e21408304739']
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
#   args = ['0704898c-8886-4e72-ac7d-010f4cc4fbf9']
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
#   args = ['9c9c553c-ba38-44d6-8897-885d775e53be']
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
#   args = ['2bf1caa6-9b5d-4487-8466-f0f476e025bb']
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