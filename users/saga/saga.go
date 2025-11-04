package saga

import (
	"fmt"
	"log"
)

// SagaStep represents a unit of work in the Saga.
type SagaStep struct {
	Name         string
	Original     func() (string, error) // forward action
	Compensation func(id string) error  // rollback action
}

// Saga manages and executes multiple SagaSteps.
type Saga struct {
	Steps []SagaStep
}

// Run executes all steps sequentially and compensates if any fail.
func (s *Saga) Run() error {
	var completed []SagaStep
	var ids []string

	for i, step := range s.Steps {
		log.Printf("Running step %d: %s", i+1, step.Name)
		id, err := step.Original()
		if err != nil {
			log.Printf("Step failed: %s, starting compensation...", step.Name)

			// Rollback completed steps
			for j := len(completed) - 1; j >= 0; j-- {
				cStep := completed[j]
				cID := ids[j]
				log.Printf("Compensating step: %s", cStep.Name)
				if err := cStep.Compensation(cID); err != nil {
					log.Printf("Compensation failed for %s: %v", cStep.Name, err)
				}
			}

			return fmt.Errorf("saga failed at step %s: %v", step.Name, err)
		}

		completed = append(completed, step)
		ids = append(ids, id)
	}

	log.Println("Saga completed successfully.")
	return nil
}
