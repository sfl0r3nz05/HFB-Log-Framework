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
#   args = ['c371cc1f-c164-4226-bc21-6e5bdf003b9a']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org2.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)

#   #   #   #Query a chaincode
#   args = ['06ddb22d-fbe9-43a2-a9ff-de8000c865e0']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org2.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)
#   
#   #   #   #Query a chaincode
#   args = ['79e9148b-d18c-4b00-a0d0-05a6d096bdcc']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org2.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)
#      
#   #   #   #   #Query a chaincode
#   args = ['bfa51510-a512-4117-b373-ae634b26347d']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org2.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)
#   
#   #   #   #Query a chaincode
#   args = ['5cff99fc-9a76-46d5-9514-0d846ab9ea6b']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org2.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)

#   #   #   #Query a chaincode
#   args = ['1df9e263-46c7-4e88-bee1-28b2901b041b']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='modbuschannel',
#       peers=['peer0.org2.example.com'],
#       args=args,
#       cc_name='base_cc',
#       fcn="get"
#   ))
#   print("response", response)

#   #   #Query a chaincode
args = ['8a6f5eb6-97e5-460d-85e7-2e4a2fb919d5']
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