package domain

type IEntityRegistry interface {
	RegisterEntity(entity Entity) error
	RegisterEntityHierarchy(parentID, childID string) error
	GetParentsForEntityInOrder(childEntity *Entity) ([]*Entity, error)
}

type IFeatureToggleManager interface {
	SetFeatureToggle(toggle FeatureToggle) error
	IsFeatureEnabled(ftoggle FeatureToggle) (bool, error)
}
