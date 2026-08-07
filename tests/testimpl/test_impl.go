package testimpl

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	id := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "id")
	zoneID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "zone_id")
	vpcID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "vpc_id")
	vpcRegion := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "vpc_region")
	hostedZoneName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "hosted_zone_name")
	zoneVpcID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "zone_vpc_id")

	// Verify the ID format is zone_id:vpc_id
	assert.Equal(t, zoneID+":"+vpcID, id, "ID should be in format zone_id:vpc_id")

	// Verify the hosted zone name ends with the expected domain suffix from test.tfvars
	assert.True(t, strings.HasSuffix(hostedZoneName, ".internal"), "Hosted zone name should end with .internal domain suffix")

	// The authorized VPC and zone VPC must be different resources
	assert.NotEqual(t, zoneVpcID, vpcID, "Authorized VPC and zone VPC should be different")

	cfg, err := config.LoadDefaultConfig(context.Background())
	require.NoError(t, err, "Failed to load AWS config")

	client := route53.NewFromConfig(cfg)

	// Verify via AWS API that the VPC association authorization exists
	t.Run("TestAuthorizationExistsViaAPI", func(t *testing.T) {
		result, err := client.ListVPCAssociationAuthorizations(context.Background(), &route53.ListVPCAssociationAuthorizationsInput{
			HostedZoneId: &zoneID,
		})
		require.NoError(t, err, "Failed to list VPC association authorizations")

		found := false
		for _, vpc := range result.VPCs {
			if *vpc.VPCId == vpcID {
				found = true
				assert.Equal(t, vpcRegion, string(vpc.VPCRegion), "VPC region should match")
				break
			}
		}
		require.True(t, found, "VPC association authorization should exist in AWS API response")
	})

	// Verify hosted zone name matches the Terraform output via the API
	t.Run("TestHostedZoneNameMatchesOutput", func(t *testing.T) {
		zone, err := client.GetHostedZone(context.Background(), &route53.GetHostedZoneInput{
			Id: &zoneID,
		})
		require.NoError(t, err, "Failed to get hosted zone")
		require.NotNil(t, zone.HostedZone, "Hosted zone should exist")
		assert.True(t, zone.HostedZone.Config.PrivateZone, "Hosted zone should be private")
		// Route53 appends a trailing dot to zone names
		assert.Equal(t, hostedZoneName+".", *zone.HostedZone.Name, "Hosted zone name from API should match Terraform output")
	})

	// Write operation: exercise the authorization by associating the VPC with
	// the hosted zone, then disassociate to clean up.
	// NOTE: We do not use a subtest here so that the deferred disassociate
	// cleanup runs even if an assertion fails, preventing orphaned associations
	// that would block terraform destroy.
	region := route53types.VPCRegion(vpcRegion)
	assocOut, err := client.AssociateVPCWithHostedZone(context.Background(), &route53.AssociateVPCWithHostedZoneInput{
		HostedZoneId: &zoneID,
		VPC: &route53types.VPC{
			VPCId:     &vpcID,
			VPCRegion: region,
		},
		Comment: strPtr("Functional test: exercising VPC association authorization"),
	})
	require.NoError(t, err, "AssociateVPCWithHostedZone should succeed using the authorization")
	require.NotNil(t, assocOut.ChangeInfo, "AssociateVPCWithHostedZone should return change info")
	require.NotNil(t, assocOut.ChangeInfo.Id, "Associate change should have an ID")

	// Defer disassociate immediately so it runs even if subsequent assertions fail
	defer func() {
		disassocOut, disassocErr := client.DisassociateVPCFromHostedZone(context.Background(), &route53.DisassociateVPCFromHostedZoneInput{
			HostedZoneId: &zoneID,
			VPC: &route53types.VPC{
				VPCId:     &vpcID,
				VPCRegion: region,
			},
			Comment: strPtr("Functional test: cleanup after exercising VPC association authorization"),
		})
		if disassocErr != nil {
			t.Errorf("DisassociateVPCFromHostedZone cleanup failed: %v", disassocErr)
			return
		}
		if disassocOut.ChangeInfo != nil && disassocOut.ChangeInfo.Id != nil {
			waitForRoute53ChangeInSync(t, client, *disassocOut.ChangeInfo.Id)
		}
	}()

	waitForRoute53ChangeInSync(t, client, *assocOut.ChangeInfo.Id)
}

func TestComposableCompleteReadonly(t *testing.T, ctx types.TestContext) {
	id := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "id")
	zoneID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "zone_id")
	vpcID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "vpc_id")
	vpcRegion := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "vpc_region")
	hostedZoneName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "hosted_zone_name")
	zoneVpcID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "zone_vpc_id")

	// Verify the ID format is zone_id:vpc_id
	assert.Equal(t, zoneID+":"+vpcID, id, "ID should be in format zone_id:vpc_id")

	// Verify the hosted zone name ends with the expected domain suffix from test.tfvars
	assert.True(t, strings.HasSuffix(hostedZoneName, ".internal"), "Hosted zone name should end with .internal domain suffix")

	// The authorized VPC and zone VPC must be different resources
	assert.NotEqual(t, zoneVpcID, vpcID, "Authorized VPC and zone VPC should be different")

	cfg, err := config.LoadDefaultConfig(context.Background())
	require.NoError(t, err, "Failed to load AWS config")

	client := route53.NewFromConfig(cfg)

	// Read-only: verify the authorization exists via the list API
	t.Run("TestAuthorizationExistsViaAPI", func(t *testing.T) {
		result, err := client.ListVPCAssociationAuthorizations(context.Background(), &route53.ListVPCAssociationAuthorizationsInput{
			HostedZoneId: &zoneID,
		})
		require.NoError(t, err, "Failed to list VPC association authorizations")

		found := false
		for _, vpc := range result.VPCs {
			if *vpc.VPCId == vpcID {
				found = true
				assert.Equal(t, vpcRegion, string(vpc.VPCRegion), "VPC region should match")
				break
			}
		}
		require.True(t, found, "VPC association authorization should exist in AWS API response")
	})

	// Read-only: verify the hosted zone exists, is private, and name matches output
	t.Run("TestHostedZoneViaAPI", func(t *testing.T) {
		zone, err := client.GetHostedZone(context.Background(), &route53.GetHostedZoneInput{
			Id: &zoneID,
		})
		require.NoError(t, err, "Failed to get hosted zone")
		require.NotNil(t, zone.HostedZone, "Hosted zone should exist")
		assert.True(t, zone.HostedZone.Config.PrivateZone, "Hosted zone should be private")
		// Route53 appends a trailing dot to zone names
		assert.Equal(t, hostedZoneName+".", *zone.HostedZone.Name, "Hosted zone name from API should match Terraform output")
	})
}

func strPtr(s string) *string {
	return &s
}

func waitForRoute53ChangeInSync(t *testing.T, client *route53.Client, changeID string) {
	t.Helper()
	ctx := context.Background()
	deadline := time.Now().Add(5 * time.Minute)
	for time.Now().Before(deadline) {
		out, err := client.GetChange(ctx, &route53.GetChangeInput{Id: &changeID})
		require.NoError(t, err, "GetChange should succeed while waiting for INSYNC")
		require.NotNil(t, out.ChangeInfo, "GetChange should return change info")
		if out.ChangeInfo.Status == route53types.ChangeStatusInsync {
			return
		}
		time.Sleep(2 * time.Second)
	}
	require.Fail(t, "timed out waiting for Route53 change to reach INSYNC", "changeID=%s", changeID)
}
