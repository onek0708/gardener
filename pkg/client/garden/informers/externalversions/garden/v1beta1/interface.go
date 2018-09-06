// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/gardener/gardener/pkg/client/garden/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BackupInfrastructures returns a BackupInfrastructureInformer.
	BackupInfrastructures() BackupInfrastructureInformer
	// CloudProfiles returns a CloudProfileInformer.
	CloudProfiles() CloudProfileInformer
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
	// Quotas returns a QuotaInformer.
	Quotas() QuotaInformer
	// SecretBindings returns a SecretBindingInformer.
	SecretBindings() SecretBindingInformer
	// Seeds returns a SeedInformer.
	Seeds() SeedInformer
	// Shoots returns a ShootInformer.
	Shoots() ShootInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BackupInfrastructures returns a BackupInfrastructureInformer.
func (v *version) BackupInfrastructures() BackupInfrastructureInformer {
	return &backupInfrastructureInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CloudProfiles returns a CloudProfileInformer.
func (v *version) CloudProfiles() CloudProfileInformer {
	return &cloudProfileInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Quotas returns a QuotaInformer.
func (v *version) Quotas() QuotaInformer {
	return &quotaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecretBindings returns a SecretBindingInformer.
func (v *version) SecretBindings() SecretBindingInformer {
	return &secretBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Seeds returns a SeedInformer.
func (v *version) Seeds() SeedInformer {
	return &seedInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Shoots returns a ShootInformer.
func (v *version) Shoots() ShootInformer {
	return &shootInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
