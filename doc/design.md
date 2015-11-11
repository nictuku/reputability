## Summary
Reputability is a peer to peer system that provides distributed information about the reliability of *service sharing* nodes and gives them *incentives to share resources*.

When a node A *provides a service* to a node B, it gains *reputation* that is acknowledged by the network of nodes that trust B.

A can can then *use* or *transfer* the benefits of its reputation to other nodes.


## Derivatives

Higher-level services can be built on top of infrastructure services, using the same reputation currency.

Infrastructure services receive *premium reputation* which incentivizes providers to prioritize their reliability. Conversely, high-level services should only give a fractional return of reputation credit compared to base ones with the same resource expenses.

## Example use cases
- network of computers storing reputation history.
- network of peers that provide distributed encrypted blob storage. Credits are given when the data is stored and periodically when it's proven to remain stored.
- network of raw computing power providers


## Basic operation
- initial node layout A - B - C - E
- E needs service.
- E searches the network for a *provider* and finds A
- A
- A stores file for E
- E now trusts A.
- E reliability score for A = bytes * seconds
- A reliability score for E = 0
- C

## Challenges
- how to **persist reputation data** despite node storage being non-permanent
- how to **keep information accurate** despite node location being highly ephemeral
-

## Definitions

- node: routing element. Does not have intrinsic reliability or reputation
- peer: service element (e.g: storage server). Has reputation.

## Procedures
- Find most reliable peer with [search criteria]
- Find a node's total reliability score, according to direct nodes.

Upkeep:
- Recalculate neighbor's reliability


Example procedure use cases:
 - Find my own reputation and display it in the UI.

Basic principles:

- node A keeps track of other nodes, with two attributes: -
node should not get penalized from opening new connections (should have an
incentive, maybe use sum of reputation)
