package domain

type FeatureToggle struct {
	entity      *Entity // entity used
	featureName string  // Name of the feature
	enabled     bool    // Whether the feature is enabled or not
}

func (ft *FeatureToggle) GetEntity() *Entity {
	if ft == nil || ft.entity == nil {
		return nil
	}
	return ft.entity
}

func (ft *FeatureToggle) SetEntity(entity *Entity) {
	if ft == nil {
		return
	}
	ft.entity = entity
}

func (ft *FeatureToggle) GetFeatureName() string {
	if ft == nil {
		return ""
	}
	return ft.featureName
}

func (ft *FeatureToggle) SetFeatureName(name string) {
	if ft == nil {
		return
	}
	ft.featureName = name
}

func (ft *FeatureToggle) IsEnabled() bool {
	if ft == nil {
		return false
	}
	return ft.enabled
}

func (ft *FeatureToggle) SetEnabled(enabled bool) {
	if ft == nil {
		return
	}
	ft.enabled = enabled
}
