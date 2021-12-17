package creational

/*

Summary:
Builder pattern is used to create complex objects in steps.

Components:
1. Product: The core complex object (Robot)
2. BuilderInterface: Defines what makes a builder
3. BuilderImplementations: Implements builder interface and makes the
required Product (Tesla and Apple)
4. Director: Director is used for decoupling client from build process.
It decouples the client code from the steps, logic and order of
creating our beloved product. https://stackoverflow.com/a/15317136/4106031.
5. Client: Client uses director to create product.

Example:
Create a Robot, creating parts of robots are complicated operations, so this is
a good example. Also, it has 2 makers to show case how one builder interface
can be used by different products.

   d := NewDirector(builder)
   product = d.Build()

Benefits:
1. Create objects piece by piece. One big constructor is avoided with this.
2. We can use a single build interface to create different products (tesla robot and apple robot)
3. Readability and maintainability.
*/

type RobotMaker int

const (
	Apple RobotMaker = iota
	Tesla
)

// Robot
// Component: Product
type Robot struct {
	head  bool
	torso bool
	legs  bool
	arms  bool
}

// RobotBuilder defines what the robot has
// Component: BuilderInterface
type RobotBuilder interface {
	BuildRobotHead()
	BuildRobotTorso()
	BuildRobotLegs()
	BuildRobotArms()

	Build() Robot
}

func NewRobotBuilder(robotMaker RobotMaker) RobotBuilder {
	switch robotMaker {
	case Apple:
		return &appleRobotBuilder{robot: &Robot{}}
	case Tesla:
		return &teslaRobotBuilder{robot: &Robot{}}
	}

	return &teslaRobotBuilder{robot: &Robot{}}
}

// appleRobotBuilder implement RobotBuilder and creates a robot
// Component: BuilderImplementation
type appleRobotBuilder struct {
	robot *Robot
}

func (b *appleRobotBuilder) BuildRobotHead()  { b.robot.head = true }
func (b *appleRobotBuilder) BuildRobotTorso() { b.robot.torso = true }
func (b *appleRobotBuilder) BuildRobotLegs()  { b.robot.legs = true }
func (b *appleRobotBuilder) BuildRobotArms()  { b.robot.arms = true }
func (b *appleRobotBuilder) Build() Robot {
	return *b.robot
}

// teslaRobotBuilder implement RobotBuilder and creates a robot
// Component: BuilderImplementation
type teslaRobotBuilder struct {
	robot *Robot
}

func (b *teslaRobotBuilder) BuildRobotHead()  { b.robot.head = true }
func (b *teslaRobotBuilder) BuildRobotTorso() { b.robot.torso = true }
func (b *teslaRobotBuilder) BuildRobotLegs()  { b.robot.legs = false } // in car
func (b *teslaRobotBuilder) BuildRobotArms()  { b.robot.arms = true }
func (b *teslaRobotBuilder) Build() Robot {
	return *b.robot
}

// Director encapsulates the robot builder and builds the robot
// by the correct maker based on the currently set builder.
// This decouples the inner creation logics and functions from the client.
type Director struct {
	builder RobotBuilder
}

func NewDirector(m RobotMaker) *Director {
	return &Director{builder: NewRobotBuilder(m)}
}

func (d *Director) SetRobotMaker(m RobotMaker) {
	d.builder = NewRobotBuilder(m)
}

func (d *Director) Build() Robot {
	d.builder.BuildRobotHead()
	d.builder.BuildRobotTorso()
	d.builder.BuildRobotLegs()
	d.builder.BuildRobotArms()

	return d.builder.Build()
}
