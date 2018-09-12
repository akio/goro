package goro

import (
	"github.com/akio/goro/urdf"
	"github.com/ungerik/go3d/mat3"
	"github.com/ungerik/go3d/mat4"
	"github.com/ungerik/go3d/quaternion"
	"github.com/ungerik/go3d/vec3"
)

const (
	JointTypeFixed = iota
	JointTypeRevolute
	JointTypePrismatic
	JointTypeFloating
)

type RigidBody struct {
	Name          string
	Position      vec3.T
	Orientation   quaternion.T
	CenterOfMass  vec3.T
	Inertia       mat3.T
	VisualMesh    interface{}
	CollisionMesh interface{}
}

type Joint struct {
	Name string
}

type RobotModel struct {
	Name   string
	Bodies []RigidBody
	Joints []Joint
}

func NewRobotModel(urdfText string) (*RobotModel, error) {
	tree, err := urdf.LoadFromString(urdfText)
	if err != nil {
		return nil, err
	}

	r := new(RobotModel)
	r.Name = tree.Name

	return r, nil
}

func (r *RobotModel) SetJoints(qs []float64) error {

	return nil
}

func (r *RobotModel) GetJacobian() (mat4.T, error) {
	return mat4.Ident, nil
}

func (r *RobotModel) SolveIK(goal mat4.T) ([]float64, error) {

	return []float64{}, nil
}

func (r *RobotModel) InCollision() bool {
	return false
}
