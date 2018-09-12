package urdf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testURDF = `
<robot name="myrobot">
	<link name="link1">
	   <inertial>
		 <origin xyz="0 0 0.5" rpy="0 0 0"/>
		 <mass value="1"/>
		 <inertia ixx="100"  ixy="0"  ixz="0" iyy="100" iyz="0" izz="100" />
	   </inertial>

	   <visual>
		 <origin xyz="0 0 0" rpy="0 0 0" />
		  <geometry>
			<box size="1 1 1" />
		  </geometry>
		  <material name="Cyan">
			<color rgba="0 1.0 1.0 1.0"/>
		  </material>
		</visual>

		<collision>
		  <origin xyz="0 0 0" rpy="0 0 0"/>
		  <geometry>
			<cylinder radius="1" length="0.5"/>
		  </geometry>
		</collision>
	</link>

	<joint name="joint" type="floating">
		<origin xyz="0 0 1" rpy="0 0 3.1416"/>
		<parent link="link1"/>
		<child link="link2"/>

		<calibration rising="0.0"/>
		<dynamics damping="0.0" friction="0.0"/>
		<limit effort="30" velocity="1.0" lower="-2.2" upper="0.7" />
		<safety_controller k_velocity="10" k_position="15" soft_lower_limit="-2.0" soft_upper_limit="0.5" />
	</joint>

	<link name="link2">
	   <inertial>
		 <origin xyz="0.5 0 0.5" rpy="0 0 0"/>
		 <mass value="1"/>
		 <inertia ixx="100"  ixy="0"  ixz="0" iyy="100" iyz="0" izz="100" />
	   </inertial>

	   <visual>
		 <origin xyz="0 0 0" rpy="0 0 0" />
		  <geometry>
			<box size="1 1 1" />
		  </geometry>
		  <material name="Cyan">
			<color rgba="0 1.0 1.0 1.0"/>
		  </material>
		</visual>

		<collision>
		  <origin xyz="0 0 0" rpy="0 0 0"/>
		  <geometry>
			<cylinder radius="1" length="0.5"/>
		  </geometry>
		</collision>
	</link>
</robot>
`

func TestReadShortURDF(t *testing.T) {
	robot, err := LoadFromString(testURDF)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, robot.Name, "myrobot", "Robot name didn't match")

	assert.Equal(t, len(robot.Links), 2, "Number of links didn't match")

	assert.Equal(t, len(robot.Joints), 1, "Number of joints didn't match")

	link1 := robot.Links[robot.FindLink("link1")]
	assert.Equal(t, link1.Name, "link1", "Link 1 name")

	link2 := robot.Links[robot.FindLink("link2")]
	assert.Equal(t, link2.Name, "link2", "Link 2 name")

	joint := robot.Joints[robot.FindJoint("joint")]
	assert.Equal(t, joint.Name, "joint", "Joint name")
}
