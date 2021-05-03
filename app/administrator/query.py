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

#   #   #Query a chaincode
#   args = ['1a02f09f-2113-4c21-b07e-0715fe5f21f2']
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
#   #   #   #   
#   #   #   #Query a chaincode
#   args = ['38c6605b-a2e7-45b6-a4ac-7a4909e13636']
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
#   #   #   #
#   #   #   #Query a chaincode
#   args = ['60149fae-6342-487a-a356-97e4485bd5de']
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
#   #   #   #
#   #   #   #Query a chaincode
#   args = ['f9f68626-0278-4a34-bf43-a5973f1c1e68']
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