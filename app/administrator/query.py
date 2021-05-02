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

# Query a chaincode
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

# Query a chaincode
#   args = ['1c8e1071-cd05-4a1f-966d-5d7ef3909b38']
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
#   #   
#   args = ['bceaa089-b393-4ac5-a0de-ef45496fefd7']
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
#   #   
#   args = ['d1dcaa2d-d5f5-4565-9799-3109e7ab63d3']
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