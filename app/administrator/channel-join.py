import asyncio
from hfc.fabric import Client
from hfc.fabric_network import wallet
from hfc.fabric_ca.caservice import ca_service

loop = asyncio.get_event_loop()

cli = Client(net_profile="../connection-profile/network.json")
org1_admin = cli.get_user(org_name='org1.example.com', name='Admin')
org2_admin = cli.get_user(org_name='org2.example.com', name='Admin')

# Create a New Channel, the response should be true if succeed
response = loop.run_until_complete(cli.channel_create(
    orderer='orderer.example.com',
    channel_name='channel1',
    requestor=org1_admin,
    config_yaml='../crypto-material/config_solo/',
    channel_profile='TwoOrgsChannel'
))
print(response == True)

#   response = loop.run_until_complete(cli.channel_create(
#       orderer='orderer.example.com',
#       channel_name='channel1',
#       requestor=org1_admin,
#       config_yaml='../crypto-material/config_solo/',
#       channel_profile='TwoOrgsChannel'
#   ))
#   print(response == True)

# Join Peers into Channel, the response should be true if succeed
responses = loop.run_until_complete(cli.channel_join(
    requestor=org1_admin,
    channel_name='channel1',
    peers=['peer0.org1.example.com'],
    orderer='orderer.example.com'
))
print(len(responses) == 1)

#   # Join Peers into Channel, the response should be true if succeed
#   responses = loop.run_until_complete(cli.channel_join(
#       requestor=org1_admin,
#       channel_name='channel1',
#       peers=['peer0.org1.example.com'],
#       orderer='orderer.example.com'
#   ))
#   print(len(responses) == 1)

# For operations on peers from org2.example.com, org2_admin is required as requestor
responses = loop.run_until_complete(cli.channel_join(
    requestor=org2_admin,
    channel_name='channel1',
    peers=['peer0.org2.example.com'],
    orderer='orderer.example.com'
))
print(len(responses) == 1)

#   # For operations on peers from org2.example.com, org2_admin is required as requestor
#   responses = loop.run_until_complete(cli.channel_join(
#       requestor=org2_admin,
#       channel_name='channel1',
#       peers=['peer0.org2.example.com'],
#       orderer='orderer.example.com'
#   ))
#   print(len(responses) == 1)