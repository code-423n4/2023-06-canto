package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestInitGenesis() {
	// check calculated epochMintProvision at genesis
	epochMintProvision, _ := suite.app.InflationKeeper.GetEpochMintProvision(suite.ctx)
	fmt.Println(suite.app.InflationKeeper.GetParams(suite.ctx))
	expMintProvision := sdk.MustNewDecFromStr("543478266666666666666667.000000000000000000")
	suite.Require().Equal(expMintProvision, epochMintProvision)
}
