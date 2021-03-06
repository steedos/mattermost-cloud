// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
//

package model

import (
	mmv1alpha1 "github.com/mattermost/mattermost-operator/pkg/apis/mattermost/v1alpha1"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

const (
	// InstallationDatabaseMysqlOperator is a database hosted in kubernetes via the operator.
	InstallationDatabaseMysqlOperator = "mysql-operator"
	// InstallationDatabaseSingleTenantRDSMySQL is a MySQL database hosted via
	// Amazon RDS.
	// TODO: update name value to aws-rds-mysql
	InstallationDatabaseSingleTenantRDSMySQL = "aws-rds"
	// InstallationDatabaseSingleTenantRDSPostgres is a PostgreSQL database hosted
	// via Amazon RDS.
	InstallationDatabaseSingleTenantRDSPostgres = "aws-rds-postgres"
	// InstallationDatabaseMultiTenantRDSMySQL is a MySQL multitenant database
	// hosted via Amazon RDS.
	// TODO: update name value to aws-multitenant-rds-mysql
	InstallationDatabaseMultiTenantRDSMySQL = "aws-multitenant-rds"
	// InstallationDatabaseMultiTenantRDSPostgres is a PostgreSQL multitenant
	// database hosted via Amazon RDS.
	InstallationDatabaseMultiTenantRDSPostgres = "aws-multitenant-rds-postgres"

	// DatabaseEngineTypeMySQL is a MySQL database.
	DatabaseEngineTypeMySQL = "mysql"
	// DatabaseEngineTypePostgres is a PostgreSQL database.
	DatabaseEngineTypePostgres = "postgres"
)

// Database is the interface for managing Mattermost databases.
type Database interface {
	Provision(store InstallationDatabaseStoreInterface, logger log.FieldLogger) error
	Teardown(store InstallationDatabaseStoreInterface, keepData bool, logger log.FieldLogger) error
	Snapshot(store InstallationDatabaseStoreInterface, logger log.FieldLogger) error
	GenerateDatabaseSpecAndSecret(store InstallationDatabaseStoreInterface, logger log.FieldLogger) (*mmv1alpha1.Database, *corev1.Secret, error)
}

// InstallationDatabaseStoreInterface is the interface necessary for SQLStore
// functionality to correlate an installation to a cluster for database creation.
// TODO(gsagula): Consider renaming this interface to InstallationDatabaseInterface. For reference,
// https://github.com/mattermost/mattermost-cloud/pull/209#discussion_r424597373
type InstallationDatabaseStoreInterface interface {
	GetClusterInstallations(filter *ClusterInstallationFilter) ([]*ClusterInstallation, error)
	GetMultitenantDatabase(multitenantdatabaseID string) (*MultitenantDatabase, error)
	GetMultitenantDatabases(filter *MultitenantDatabaseFilter) ([]*MultitenantDatabase, error)
	GetMultitenantDatabaseForInstallationID(installationID string) (*MultitenantDatabase, error)
	CreateMultitenantDatabase(multitenantDatabase *MultitenantDatabase) error
	UpdateMultitenantDatabase(multitenantDatabase *MultitenantDatabase) error
	LockMultitenantDatabase(multitenantdatabaseID, lockerID string) (bool, error)
	UnlockMultitenantDatabase(multitenantdatabaseID, lockerID string, force bool) (bool, error)
}

// MysqlOperatorDatabase is a database backed by the MySQL operator.
type MysqlOperatorDatabase struct{}

// NewMysqlOperatorDatabase returns a new MysqlOperatorDatabase interface.
func NewMysqlOperatorDatabase() *MysqlOperatorDatabase {
	return &MysqlOperatorDatabase{}
}

// Provision completes all the steps necessary to provision a MySQL operator database.
func (d *MysqlOperatorDatabase) Provision(store InstallationDatabaseStoreInterface, logger log.FieldLogger) error {
	logger.Info("MySQL operator database requires no pre-provisioning; skipping...")

	return nil
}

// Snapshot is not supported by the operator and it should return an error.
func (d *MysqlOperatorDatabase) Snapshot(store InstallationDatabaseStoreInterface, logger log.FieldLogger) error {
	logger.Error("Snapshotting is not supported by the MySQL operator.")

	return errors.New("not implemented")
}

// Teardown removes all MySQL operator resources for a given installation.
func (d *MysqlOperatorDatabase) Teardown(store InstallationDatabaseStoreInterface, keepData bool, logger log.FieldLogger) error {
	logger.Info("MySQL operator database requires no teardown; skipping...")
	if keepData {
		logger.Warn("Database preservation was requested, but isn't currently possible with the MySQL operator")
	}

	return nil
}

// GenerateDatabaseSpecAndSecret creates the k8s database spec and secret for
// accessing the MySQL operator database.
func (d *MysqlOperatorDatabase) GenerateDatabaseSpecAndSecret(store InstallationDatabaseStoreInterface, logger log.FieldLogger) (*mmv1alpha1.Database, *corev1.Secret, error) {
	return nil, nil, nil
}

// InternalDatabase returns true if the installation's database is internal
// to the kubernetes cluster it is running on.
func (i *Installation) InternalDatabase() bool {
	return i.Database == InstallationDatabaseMysqlOperator
}

// IsSupportedDatabase returns true if the given database string is supported.
func IsSupportedDatabase(database string) bool {
	switch database {
	case InstallationDatabaseSingleTenantRDSMySQL:
	case InstallationDatabaseSingleTenantRDSPostgres:
	case InstallationDatabaseMultiTenantRDSMySQL:
	case InstallationDatabaseMultiTenantRDSPostgres:
	case InstallationDatabaseMysqlOperator:
	default:
		return false
	}

	return true
}
