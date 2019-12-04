package intcode

// Program is an Intcode program.
type Program []int

// Memory is the Intcode machine memory
type Memory []int

type opcode int

const (
	opAdd  opcode = 1
	opMult        = 2
	opHalt        = 99
)

// Machine is an Intcode machine.
type Machine struct {
	ip  uint
	mem Memory
}

// New returns a new Intcode machine.
func New() *Machine {
	return &Machine{
		ip:  0,
		mem: nil,
	}
}

// Load loads a program into the machine memory.
func (c *Machine) Load(p Program) {
	c.ip = 0
	c.mem = make(Memory, len(p))
	copy(c.mem, Memory(p))
}

// Run runs an Intcode program. Changes are reflected in the machines memory.
func (c *Machine) Run() {
	for {
		switch opcode(c.mem[c.ip]) {
		case opAdd:
			c.executeAdd()
		case opMult:
			c.executeMultiply()
		case opHalt:
			return
		}
	}
}

// Memory returns the contents of the machine memory.
func (c *Machine) Memory() Memory {
	return c.mem
}

func (c *Machine) executeAdd() {
	c.mem[c.mem[c.ip+3]] = c.mem[c.mem[c.ip+1]] + c.mem[c.mem[c.ip+2]]
	c.ip += 4
}

func (c *Machine) executeMultiply() {
	c.mem[c.mem[c.ip+3]] = c.mem[c.mem[c.ip+1]] * c.mem[c.mem[c.ip+2]]
	c.ip += 4
}
