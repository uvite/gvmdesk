package snowflake

import (
	"errors"

	"github.com/uvite/gvmdesk/pkg/platform"
	"github.com/uvite/gvmdesk/pkg/utils/snowflake"

	rand2 "github.com/uvite/gvmdesk/pkg/rand"
	"math/rand"
	"sync"
	"time"




)

var seededRand *rand.Rand

func init() {
	lockedSource :=  rand2.NewLockedSourceFromSeed(time.Now().UnixNano())
	seededRand = rand.New(lockedSource)
	SetGlobalMachineID(seededRand.Intn(1023))
}

var globalmachineID struct {
	id  int
	set bool
	sync.RWMutex
}

// ErrGlobalIDBadVal means that the global machine id value wasn't properly set.
var ErrGlobalIDBadVal = errors.New("globalID must be a number between (inclusive) 0 and 1023")

// SetGlobalMachineID returns the global machine id.  This number is limited to a number between 0 and 1023 inclusive.
func SetGlobalMachineID(id int) error {
	if id > 1023 || id < 0 {
		return ErrGlobalIDBadVal
	}
	globalmachineID.Lock()
	globalmachineID.id = id
	globalmachineID.set = true
	globalmachineID.Unlock()
	return nil
}

// GlobalMachineID returns the global machine id.  This number is limited to a number between 0 and 1023 inclusive.
func GlobalMachineID() int {
	var id int
	globalmachineID.RLock()
	id = int(globalmachineID.id)
	globalmachineID.RUnlock()
	return id
}

// NewDefaultIDGenerator returns an *IDGenerator that uses the currently set global machine ID.
// If you change the global machine id, it will not change the id in any generators that have already been created.
func NewDefaultIDGenerator() *IDGenerator {
	globalmachineID.RLock()
	defer globalmachineID.RUnlock()
	if globalmachineID.set {
		return NewIDGenerator(WithMachineID(globalmachineID.id))
	}
	return NewIDGenerator()
}

// IDGenerator holds the ID generator.
type IDGenerator struct {
	Generator *snowflake.Generator
}

// IDGeneratorOp is an option for an IDGenerator.
type IDGeneratorOp func(*IDGenerator)

// WithMachineID uses the low 12 bits of machineID to set the machine ID for the snowflake ID.
func WithMachineID(machineID int) IDGeneratorOp {
	return func(g *IDGenerator) {
		g.Generator = snowflake.New(machineID & 1023)
	}
}

// NewIDGenerator returns a new IDGenerator.  Optionally you can use an IDGeneratorOp.
// to use a specific Generator
func NewIDGenerator(opts ...IDGeneratorOp) *IDGenerator {
	gen := &IDGenerator{}
	for _, f := range opts {
		f(gen)
	}
	if gen.Generator == nil {
		machineId := seededRand.Intn(1023)
		gen.Generator = snowflake.New(machineId)
	}
	return gen
}

// ID returns the next platform.ID from an IDGenerator.
func (g *IDGenerator) ID() platform.ID {
	var id platform.ID
	for !id.Valid() {
		id = platform.ID(g.Generator.Next())
	}
	return id
}
