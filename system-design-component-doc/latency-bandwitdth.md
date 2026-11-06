# Latency

Latency is the time it takes for a message or a packet to travel from source to destination. This is a simple definition but it hides the undergoing maths or details that is doing on this delay.

There are some major components that are contributing this delay or duration.

Propagation delay:

Amount of time required for a message to travel from sender to receiver, which is a function of distance over speed with which the signal propagates.

Transmission Delay:
Amount of time required to push all the packet's bits into the link, which is a function of the packet's length and data rate of the link


Processing delay:
Amount of time required to process the packet header , check for bit-level errors and determine the packet destination/

Queuing delay

Amount of time the packet is waiting in the queue until it can be processed.

The total latency between client and server is the summation of all the delays just listed.

# confusion part:

Then whats the difference between propagation delay and latency both sounds same.

Transmission delay says amount of time it requires to push the packet's bits into the link. but why we have to push the packet's bits into the link

processing delay: what it does after processing the packet header. what it determines by checking the bit level errors so does it mean it starts when is about a packet is going to be traveled.

The speed of light in fiber is around 200000000 = 2*10^8 meters per second

Need to verify the statement once the 300 millisecond delay threshold is exceeded the interaction is often reported as sluggish.

Keyword: Last-Mile Latency