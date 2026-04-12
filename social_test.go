package fluxsocial

import "testing"

func TestAddAgent(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "Alice", RoleLeader)
	if a := g.FindAgent(1); a == nil || a.Name != "Alice" || a.Role != RoleLeader {
		t.Fatal("agent not found or wrong fields")
	}
}

func TestFindAgentNotFound(t *testing.T) {
	g := NewSocialGraph()
	if g.FindAgent(99) != nil {
		t.Fatal("expected nil for missing agent")
	}
}

func TestSetRole(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "Bob", RoleWorker)
	g.SetRole(1, RoleMentor)
	if g.agents[1].Role != RoleMentor {
		t.Fatal("role not updated")
	}
}

func TestSetRoleMissingAgent(t *testing.T) {
	g := NewSocialGraph()
	g.SetRole(99, RoleLeader) // should not panic
}

func TestAgentCount(t *testing.T) {
	g := NewSocialGraph()
	if g.AgentCount() != 0 {
		t.Fatal("expected 0 agents")
	}
	g.AddAgent(1, "A", RoleWorker)
	g.AddAgent(2, "B", RoleWorker)
	if g.AgentCount() != 2 {
		t.Fatal("expected 2 agents")
	}
}

func TestAddRelation(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleWorker)
	g.AddAgent(2, "B", RoleWorker)
	g.AddRelation(1, 2, RelPeer)
	r := g.FindRelation(1, 2)
	if r == nil || r.Type != RelPeer {
		t.Fatal("relation not found")
	}
}

func TestFindRelationNotFound(t *testing.T) {
	g := NewSocialGraph()
	if g.FindRelation(1, 2) != nil {
		t.Fatal("expected nil")
	}
}

func TestNeighbors(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleWorker)
	g.AddAgent(2, "B", RoleWorker)
	g.AddAgent(3, "C", RoleWorker)
	g.AddRelation(1, 2, RelPeer)
	g.AddRelation(1, 3, RelMentor)
	n := g.Neighbors(1)
	if len(n) != 2 {
		t.Fatalf("expected 2 neighbors, got %d", len(n))
	}
}

func TestNeighborsUndirected(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleWorker)
	g.AddAgent(2, "B", RoleWorker)
	g.AddRelation(1, 2, RelPeer)
	n2 := g.Neighbors(2)
	if len(n2) != 1 || n2[0] != 1 {
		t.Fatal("expected neighbor 1 from node 2")
	}
}

func TestNeighborsEmpty(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleWorker)
	if len(g.Neighbors(1)) != 0 {
		t.Fatal("expected no neighbors")
	}
}

func TestCentrality(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleWorker)
	g.AddAgent(2, "B", RoleWorker)
	g.AddAgent(3, "C", RoleWorker)
	g.AddRelation(1, 2, RelPeer)
	g.AddRelation(1, 3, RelPeer)
	c := g.Centrality(1)
	if c != 1.0 {
		t.Fatalf("expected centrality 1.0, got %f", c)
	}
	c2 := g.Centrality(2)
	if c2 != 0.5 {
		t.Fatalf("expected centrality 0.5, got %f", c2)
	}
}

func TestCentralitySingleNode(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleWorker)
	if g.Centrality(1) != 0.0 {
		t.Fatal("expected 0 centrality for single node")
	}
}

func TestCreateGroup(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "Alice", RoleLeader)
	id := g.CreateGroup("Team", 1)
	if id != 1 {
		t.Fatalf("expected group id 1, got %d", id)
	}
	grp := g.groups[1]
	if grp.Name != "Team" || grp.LeaderID != 1 {
		t.Fatal("wrong group fields")
	}
}

func TestJoinGroup(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleLeader)
	g.AddAgent(2, "B", RoleWorker)
	gid := g.CreateGroup("Team", 1)
	g.JoinGroup(gid, 2)
	members := g.GroupMembers(gid)
	if len(members) != 2 {
		t.Fatalf("expected 2 members, got %d", len(members))
	}
}

func TestJoinGroupNoDupes(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleLeader)
	gid := g.CreateGroup("Team", 1)
	g.JoinGroup(gid, 1) // already member
	if len(g.groups[gid].Members) != 1 {
		t.Fatal("duplicate member added")
	}
}

func TestGroupMembersMissing(t *testing.T) {
	g := NewSocialGraph()
	if g.GroupMembers(99) != nil {
		t.Fatal("expected nil for missing group")
	}
}

func TestRelationDefaults(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleWorker)
	g.AddAgent(2, "B", RoleWorker)
	g.AddRelation(1, 2, RelCollaborator)
	r := g.FindRelation(1, 2)
	if r.Weight != 1.0 || r.Interactions != 1 {
		t.Fatal("wrong default weight/interactions")
	}
}

func TestMultipleGroups(t *testing.T) {
	g := NewSocialGraph()
	g.AddAgent(1, "A", RoleLeader)
	g.AddAgent(2, "B", RoleLeader)
	id1 := g.CreateGroup("G1", 1)
	id2 := g.CreateGroup("G2", 2)
	if id1 == id2 {
		t.Fatal("group ids should differ")
	}
}
