# fluxsocial-go 👥

**Go social graph for agent fleets.** Agents, relationships, roles, groups, centrality. Kubernetes-native, microservice-ready.

```go
g := fluxsocial.NewSocialGraph()
g.AddAgent(1, "worker-a", fluxsocial.RoleWorker)
g.AddAgent(2, "coordinator-b", fluxsocial.RoleCoordinator)
g.AddRelation(1, 2, fluxsocial.RelMentor)

fmt.Println(g.Centrality(1))      // 0.5
fmt.Println(g.Neighbors(1))       // [2]
fmt.Println(g.AgentCount())       // 2
```

## API

```go
// Create graph
g := fluxsocial.NewSocialGraph()

// Agents
g.AddAgent(1, "worker-a", fluxsocial.RoleWorker)
agent := g.FindAgent(1)
g.SetRole(1, fluxsocial.RoleLeader)

// Relationships
g.AddRelation(1, 2, fluxsocial.RelMentor)
rel := g.FindRelation(1, 2)

// Queries
neighbors := g.Neighbors(1)
centrality := g.Centrality(1)

// Groups
groupID := g.CreateGroup("sensor-team", 2)
g.JoinGroup(groupID, 1)
members := g.GroupMembers(groupID)
```

### Roles: `RoleWorker`, `RoleCoordinator`, `RoleSpecialist`, `RoleLeader`, `RoleMentor`, `RoleLearner`, `RoleObserver`

### Relations: `RelPeer`, `RelMentor`, `RelStudent`, `RelSubordinate`, `RelCollaborator`, `RelCompetitor`, `RelStranger`

## Install

```bash
go get github.com/Lucineer/fluxsocial-go
```

## Fleet Context

Part of the Lucineer/Cocapn fleet. Go variant of [flux-social](https://github.com/Lucineer/flux-social) (Rust). Pairs with [fluxtrust-go](https://github.com/Lucineer/fluxtrust-go) for trust-weighted social routing.
