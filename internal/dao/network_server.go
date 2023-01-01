// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/sagoo-cloud/sagooiot/internal/dao/internal"
)

// internalNetworkServerDao is internal type for wrapping internal DAO implements.
type internalNetworkServerDao = *internal.NetworkServerDao

// networkServerDao is the data access object for table network_server.
// You can define custom methods on it to extend its functionality as you wish.
type networkServerDao struct {
	internalNetworkServerDao
}

var (
	// NetworkServer is globally public accessible object for table network_server operations.
	NetworkServer = networkServerDao{
		internal.NewNetworkServerDao(),
	}
)

// Fill with you ideas below.
