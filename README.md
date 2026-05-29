# flock

<div align="center">
  <img src="./.github/assets/flock.png" alt="Flock Mascot" />
</div>

Peer-to-peer workload orchestrator.

Single binary, no external dependencies. Every node runs the same binary, joins via gossip, and can accept workload submissions. No node stores an allocation table. Instead, every node calls a deterministic pure function that computes allocations from the current set of workloads and alive nodes. CRDTs guarantee all nodes converge to the same inputs, so the same function produces the same outputs everywhere. This eliminates the need for a leader, a consensus store, or reconciliation loops.

## Architecture

```mermaid
graph TD
    subgraph "internal/domain/"
        SCH[scheduler.go<br/>Schedule]
        SUP[supervisor.go<br/>Transition]
        WK[workload.go]
        ND[node.go]
        AL[allocation.go]
    end

    subgraph "internal/service/"
        REC[reconciler.go<br/>gossip → schedule → diff → act]
    end

    subgraph "internal/handler/"
        HTTP[http/<br/>workload CRUD]
    end

    subgraph "internal/client/ (ports)"
        RT[runtime/<br/>process lifecycle]
        DIS[discovery/<br/>DNS records]
        TEL[telemetry/<br/>wide events]
    end

    subgraph "meld (library dependency)"
        MEMB[membership/swim]
        CRDT[crdt/orset]
        GOSSIP[gossip/udp]
        HASH[hash/rendezvous]
    end

    REC --> SCH
    REC --> SUP
    REC --> RT
    REC --> DIS
    REC --> TEL
    REC --> MEMB
    REC --> CRDT
    REC --> GOSSIP
    SCH --> HASH
    HTTP --> REC
```

## Scheduling During Partitions

```mermaid
graph TD
    subgraph "Before Partition"
        ALL[All nodes: same CRDT state<br/>Schedule → same allocations]
    end

    subgraph "During Partition"
        direction LR
        subgraph "Side A"
            A1[Node 1 + Node 2<br/>Schedule with 2 nodes]
        end
        subgraph "Side B"
            B1[Node 3<br/>Schedule with 1 node]
        end
    end

    subgraph "After Heal"
        MERGE[CRDT merge → unified state<br/>Schedule → same allocations<br/>Drain duplicates]
    end

    ALL -->|network partition| A1
    ALL -->|network partition| B1
    A1 -->|partition heals| MERGE
    B1 -->|partition heals| MERGE
```

## Dependencies

- **meld**: CRDT types, version vectors, gossip transport, SWIM membership. Must be complete.
