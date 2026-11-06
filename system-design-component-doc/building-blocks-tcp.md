# IP:
The IP stands for Internet Protocol. This is protocol for network for which sending data or packet from one device to another device is possible. This sending and receiving workflow happens step by step.

1. Breaking data into packets
 -- Large files , web pages or videos cannot be sent accross the network all at once.
 -- The sending device breaks the data down into smaller chunks called packets.
2. Adding IP headers:
  -- Each packet is wrapped with an IP header
  -- This header contains vital control information ,  most importantly the source IP address and destination IP address.
  -- Devices use either IPV4 or IPV6.

3. Routing Across Network
  -- The packet leave your device and pass through local hardware()
  -- Routers read the destination IP address on each packet and look to the routing tables to get the next route which will help to reach the destination
  -- Packets may take different physical paths accross the global networks thats why the packet may arrive out of order

 4. Reassembly and Transport Hand-Off
 IP is connectionless and "best-effort"- It does not garauntee all the packets arrives or in order.

 So there are some limitations of IP which introduces us with Transmission Control Protocol that gives us some garauntee.

 # TCP:

 Transmission Control Protocol is a communication protocol for networks that focuses on making connection betweend devices , services , applications. TCP establishes connection between sender and receiver and translate data into sequence of packets also ensure the receiver can track sequence of packet's order.

 TCP connection:

 All TCP connections begin with a three way handshake. Before the client and server can exchange data. They must agree on starting packet sequence numbers, as well as a number of other connection specific variables.