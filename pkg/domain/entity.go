package domain

import "errors"

type Entity struct {
	id         string // Unique identifier for the entity
	entityType string // Type of the entity (e.g., "restaurant", "city")
}

func (e *Entity) GetID() string {
	if e == nil {
		return ""
	}
	return e.id
}

func (e *Entity) SetID(id string) error {
	if e == nil {
		return errors.New("cannot set ID on nil Entity")
	}
	e.id = id
	return nil
}

func (e *Entity) GetType() string {
	if e == nil {
		return ""
	}
	return e.entityType
}

func (e *Entity) SetType(entityType string) error {
	if e == nil {
		return errors.New("cannot set type on nil Entity")
	}
	e.entityType = entityType
	return nil
}

type EntityHierarchy struct {
	parentEntity *Entity
	childEntity  *Entity
}

func (eh *EntityHierarchy) GetParentEntity() *Entity {
	if eh == nil || eh.parentEntity == nil {
		return nil
	}
	return eh.parentEntity
}

func (eh *EntityHierarchy) SetParentEntity(parent *Entity) {
	if eh == nil {
		return
	}
	eh.parentEntity = parent
}

func (eh *EntityHierarchy) GetChildEntity() *Entity {
	if eh == nil || eh.childEntity == nil {
		return nil
	}
	return eh.childEntity
}

func (eh *EntityHierarchy) SetChildEntity(child *Entity) {
	if eh == nil {
		return
	}
	eh.childEntity = child
}
