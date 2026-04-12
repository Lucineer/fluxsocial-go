package fluxsocial

type AgentRole int

const (
	RoleNone        AgentRole = iota
	RoleWorker
	RoleCoordinator
	RoleSpecialist
	RoleLeader
	RoleMentor
	RoleLearner
	RoleObserver
)

type RelationType int

const (
	RelPeer        RelationType = iota
	RelMentor
	RelStudent
	RelSubordinate
	RelCollaborator
	RelCompetitor
	RelStranger
)

type Agent struct {
	ID        uint16
	Name      string
	Role      AgentRole
	Reputation float64
}

type Relation struct {
	FromID      uint16
	ToID        uint16
	Type        RelationType
	Weight      float64
	Interactions uint32
}

type Group struct {
	ID       uint16
	Name     string
	LeaderID uint16
	Members  []uint16
}

type SocialGraph struct {
	agents    map[uint16]*Agent
	relations []Relation
	groups    map[uint16]*Group
	nextGroup uint16
}

func NewSocialGraph() *SocialGraph {
	return &SocialGraph{
		agents:    make(map[uint16]*Agent),
		relations: nil,
		groups:    make(map[uint16]*Group),
		nextGroup: 1,
	}
}

func (g *SocialGraph) AddAgent(id uint16, name string, role AgentRole) {
	g.agents[id] = &Agent{ID: id, Name: name, Role: role, Reputation: 0.0}
}

func (g *SocialGraph) FindAgent(id uint16) *Agent {
	return g.agents[id]
}

func (g *SocialGraph) SetRole(id uint16, role AgentRole) {
	if a, ok := g.agents[id]; ok {
		a.Role = role
	}
}

func (g *SocialGraph) AddRelation(from, to uint16, relType RelationType) {
	g.relations = append(g.relations, Relation{
		FromID: from, ToID: to, Type: relType, Weight: 1.0, Interactions: 1,
	})
}

func (g *SocialGraph) FindRelation(from, to uint16) *Relation {
	for i := range g.relations {
		if g.relations[i].FromID == from && g.relations[i].ToID == to {
			return &g.relations[i]
		}
	}
	return nil
}

func (g *SocialGraph) Neighbors(id uint16) []uint16 {
	var out []uint16
	for _, r := range g.relations {
		if r.FromID == id {
			out = append(out, r.ToID)
		} else if r.ToID == id {
			out = append(out, r.FromID)
		}
	}
	return out
}

func (g *SocialGraph) Centrality(id uint16) float64 {
	total := len(g.agents)
	if total <= 1 {
		return 0.0
	}
	deg := len(g.Neighbors(id))
	return float64(deg) / float64(total-1)
}

func (g *SocialGraph) CreateGroup(name string, leader uint16) uint16 {
	id := g.nextGroup
	g.nextGroup++
	g.groups[id] = &Group{ID: id, Name: name, LeaderID: leader, Members: []uint16{leader}}
	return id
}

func (g *SocialGraph) JoinGroup(groupID, agentID uint16) {
	if grp, ok := g.groups[groupID]; ok {
		for _, m := range grp.Members {
			if m == agentID {
				return
			}
		}
		grp.Members = append(grp.Members, agentID)
	}
}

func (g *SocialGraph) GroupMembers(groupID uint16) []*Agent {
	grp, ok := g.groups[groupID]
	if !ok {
		return nil
	}
	var out []*Agent
	for _, id := range grp.Members {
		if a := g.agents[id]; a != nil {
			out = append(out, a)
		}
	}
	return out
}

func (g *SocialGraph) AgentCount() int {
	return len(g.agents)
}
