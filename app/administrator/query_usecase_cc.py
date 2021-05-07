import os
import asyncio
from hfc.fabric import Client

loop = asyncio.get_event_loop()

cli = Client(net_profile="../connection-profile/network.json")
org1_admin = cli.get_user('org1.example.com', 'Admin')
org2_admin = cli.get_user('org2.example.com', 'Admin')

# Make the client know there is a channel in the network
cli.new_channel('channel1')

# Install Example Chaincode to Peers
# GOPATH setting is only needed to use the example chaincode inside sdk
gopath_bak = os.environ.get('GOPATH', '')
gopath = os.path.normpath(os.path.join(
    os.path.dirname(os.path.realpath('__file__')),
    '../chaincode'
))
os.environ['GOPATH'] = os.path.abspath(gopath)

###### FOR UNKNOWN REASON THIS FAIL WHEN THE QUERY IS INVOKED #####################################################
###### TO SOLVE THE ISSUE THE TRANSACTION IS EMMITED AS INVOKE API, BUT USING GET STATE INSTEAD OF PUT STATE ######

#Query a chaincode
args = ['a']
# The response should be true if succeed
response = loop.run_until_complete(cli.chaincode_invoke(
    requestor=org1_admin,
    channel_name='channel1',
    peers=['peer0.org1.example.com'],
    args=args,
    cc_name='usecase_cc',
    transient_map=None,
    fcn='get',
    wait_for_event=True,
    # cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
))

###### THIS WAS THE ORIGINAL CODE ######

#   #Query a chaincode
#   args = ['a']
#   # The response should be true if succeed
#   response = loop.run_until_complete(cli.chaincode_query(
#       requestor=org1_admin,
#       channel_name='channel1',
#       peers=['peer0.org1.example.com'],
#       args=args,
#       cc_name='usecase_cc',
#       fcn="get"
#   ))
#   print("response", response)