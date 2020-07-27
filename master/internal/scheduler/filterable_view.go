package scheduler

import (
	"fmt"
	"github.com/determined-ai/determined/master/pkg/actor"
)

// FilterableView keeps track of tasks and agents that pass the task and agent filters.
// The `TaskSummary`s and `AgentSummary` should not be modified because a reference to
// this struct is contained in another goroutine.
type FilterableView struct {
	tasks       map[TaskID]*TaskSummary
	agents      map[*actor.Ref]*AgentSummary
	taskFilter  func(*Task) bool
	agentFilter func(*agentState) bool
}

// Return a view of the scheduler state that is relevant to the provisioner. Specifically, the
// provisioner cares about (1) idle agents (2) pending tasks.
func newProvisionerView(provisionerSlotsPerInstance int) *FilterableView {
	return &FilterableView{
		tasks:       make(map[TaskID]*TaskSummary),
		agents:      make(map[*actor.Ref]*AgentSummary),
		taskFilter:  schedulableTaskFilter(provisionerSlotsPerInstance),
		agentFilter: idleAgentFilter,
		//agentFilter: noAgentFilter,
	}
}

func schedulableTaskFilter(provisionerSlotsPerInstance int) func(*Task) bool {
	// We only tell the provisioner about pending tasks that are compatible with the
	// provisioner's configured instance type.
	return func(task *Task) bool {
		slotsNeeded := task.SlotsNeeded()

		switch {
		case task.state != taskPending:
			return false
		case slotsNeeded == 0 || slotsNeeded == 1:
			return true
		case slotsNeeded%provisionerSlotsPerInstance == 0:
			return true
		default:
			return false
		}
	}
}

func noAgentFilter(agent *agentState) bool {
	return true
}

func idleAgentFilter(agent *agentState) bool {
	// this should actually work because if an agent has not connected, then it should not have any running containers
	fmt.Println("idle agent filter in use")
	fmt.Println(agent.label)
	fmt.Println(agent.containers)
	fmt.Println(agent.devices)
	fmt.Println("EXITING IDLE AGENT FILTER")
	return len(agent.containers) == 0
}

// Update updates the FilterableView with the current state of the cluster.
func (v *FilterableView) Update(rp *DefaultRP) (ViewSnapshot, bool) {
	// We must evaluate v.updateTasks(cluster) and v.updateAgents(cluster)
	// before taking the logical or of the results to ensure that short circuit
	// evaluation of booleans expressions don't prevent the updating of agents.
	fmt.Println("updating tasks and agents")
	tasksUpdateMade := v.updateTasks(rp)
	agentsUpdateMade := v.updateAgents(rp)
	return v.newSnapshot(), tasksUpdateMade || agentsUpdateMade
}

func (v *FilterableView) updateTasks(rp *DefaultRP) bool {
	newTasks := make(map[TaskID]*TaskSummary)

	for iterator := rp.taskList.iterator(); iterator.next(); {
		task := iterator.value()

		if v.taskFilter(task) {
			taskSummary := newTaskSummary(task)
			newTasks[task.ID] = &taskSummary
		}
	}

	updateMade := false
	if len(newTasks) != len(v.tasks) {
		updateMade = true
	} else {
		for _, newTask := range newTasks {
			oldTask, ok := v.tasks[newTask.ID]
			if !ok || !oldTask.equals(newTask) {
				updateMade = true
			}
		}
	}

	v.tasks = newTasks
	return updateMade
}

func (v *FilterableView) updateAgents(rp *DefaultRP) bool {
	newAgents := make(map[*actor.Ref]*AgentSummary)
	//check here to see if the actor ref used as a key can show any information regarding connection status
	fmt.Println("\ninside updateAgents. All Agents below:")
	for actorRef, state := range rp.agents {
		fmt.Println(actorRef.Address)
		fmt.Println("status: ", state)
		if v.agentFilter(state) {
			agentSummary := newAgentSummary(state)
			fmt.Println("agentSummary: ", agentSummary)
			newAgents[actorRef] = &agentSummary
		}
	}

	fmt.Println("agents after filter:", newAgents)

	updateMade := false
	if len(newAgents) != len(v.agents) {
		updateMade = true
	} else {
		for agentRef, newAgent := range newAgents {
			oldAgent, ok := v.agents[agentRef]
			if !ok || !oldAgent.equals(newAgent) {
				updateMade = true
			}
		}
	}

	v.agents = newAgents
	return updateMade
}

func (v *FilterableView) newSnapshot() ViewSnapshot {
	tasks := make([]*TaskSummary, 0, len(v.tasks))
	for _, taskSummary := range v.tasks {
		tasks = append(tasks, taskSummary)
	}
	agents := make([]*AgentSummary, 0, len(v.agents))
	for _, agent := range v.agents {
		agents = append(agents, agent)
	}
	return ViewSnapshot{Tasks: tasks, Agents: agents}
}
