/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package generate_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"

	v1alpha1 "github.com/talos-systems/talos/pkg/userdata/v1alpha1"
	udgenv1alpha1 "github.com/talos-systems/talos/pkg/userdata/v1alpha1/generate"
)

var input *udgenv1alpha1.Input

type GenerateSuite struct {
	suite.Suite
}

func TestGenerateSuite(t *testing.T) {
	suite.Run(t, new(GenerateSuite))
}

func (suite *GenerateSuite) SetupSuite() {
	var err error
	input, err = udgenv1alpha1.NewInput("test", []string{"10.0.1.5", "10.0.1.6", "10.0.1.7"})
	suite.Require().NoError(err)
}

func (suite *GenerateSuite) TestGenerateInitSuccess() {
	input.IP = net.ParseIP("10.0.1.5")
	dataString, err := udgenv1alpha1.Userdata(udgenv1alpha1.TypeInit, input)
	suite.Require().NoError(err)
	data := &v1alpha1.NodeConfig{}
	err = yaml.Unmarshal([]byte(dataString), data)
	suite.Require().NoError(err)
}

func (suite *GenerateSuite) TestGenerateControlPlaneSuccess() {
	input.IP = net.ParseIP("10.0.1.6")
	dataString, err := udgenv1alpha1.Userdata(udgenv1alpha1.TypeControlPlane, input)
	suite.Require().NoError(err)
	data := &v1alpha1.NodeConfig{}
	err = yaml.Unmarshal([]byte(dataString), data)
	suite.Require().NoError(err)
}

func (suite *GenerateSuite) TestGenerateWorkerSuccess() {
	dataString, err := udgenv1alpha1.Userdata(udgenv1alpha1.TypeJoin, input)
	suite.Require().NoError(err)
	data := &v1alpha1.NodeConfig{}
	err = yaml.Unmarshal([]byte(dataString), data)
	suite.Require().NoError(err)
}

func (suite *GenerateSuite) TestGenerateTalosconfigSuccess() {
	_, err := udgenv1alpha1.Talosconfig(input)
	suite.Require().NoError(err)
}