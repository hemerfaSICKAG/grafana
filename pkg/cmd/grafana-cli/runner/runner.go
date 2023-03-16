package runner

import (
	"github.com/grafana/grafana/pkg/infra/db"
	"github.com/grafana/grafana/pkg/services/encryption"
	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/services/k8s/authentication"
	"github.com/grafana/grafana/pkg/services/secrets"
	"github.com/grafana/grafana/pkg/services/secrets/manager"
	"github.com/grafana/grafana/pkg/services/user"
	"github.com/grafana/grafana/pkg/setting"
)

type Runner struct {
	Cfg               *setting.Cfg
	SQLStore          db.DB
	SettingsProvider  setting.Provider
	Features          featuremgmt.FeatureToggles
	EncryptionService encryption.Internal
	SecretsService    *manager.SecretsService
	SecretsMigrator   secrets.Migrator
	UserService       user.Service
	K8sAuthnAPI       authentication.K8sAuthnAPI
}

func New(cfg *setting.Cfg, sqlStore db.DB, settingsProvider setting.Provider,
	encryptionService encryption.Internal, features featuremgmt.FeatureToggles,
	secretsService *manager.SecretsService, secretsMigrator secrets.Migrator,
	userService user.Service,
	k8sAuthnAPI authentication.K8sAuthnAPI,
) Runner {
	return Runner{
		Cfg:               cfg,
		SQLStore:          sqlStore,
		SettingsProvider:  settingsProvider,
		EncryptionService: encryptionService,
		SecretsService:    secretsService,
		SecretsMigrator:   secretsMigrator,
		Features:          features,
		UserService:       userService,
		K8sAuthnAPI:       k8sAuthnAPI,
	}
}
