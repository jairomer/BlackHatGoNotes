*** TCP Scanners, and proxies ***

Chapter Objective:
- Learn and understand how to leverage the TCP protocol in Go.
- Build a concurrent port scanner.
- Build a TCP proxy that can be used for port forwarding.
- Recreate Netcat's "gaping security hole" feature.

Open: 3-Way handsahke takes place.
1. Syn      -->
2. Syn-Ack  <--
3. Ack      -->

Closed: Sent if the port is closed.
1. Syn      -->
2. Rst      <--

Filtered: By a firewall, then client will not receive a response.
1. Syn      -->
2. (timeout)


These responses are important to understand when writting network based tools.

We need to correlate the output of the tools to the low-level packet flows to
validate that a network connection has been established.

*** Bypassing firewals with port forwarding ***
- The purpose of a firewall is to prevent a client from connecting to certain
  servers and ports, while allowing others.

- Restrictions can be circumbent using an intermediary system to proxy the
  connection around or through a firewall. This technique is called port
  forwarding.

- This technique will work only with blacklist-powered firewalls.


