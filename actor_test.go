package metadata

import (
	"reflect"
	"testing"

	collection "github.com/gtyrin/go-collection"
)

func TestCompareCollections(t *testing.T) {
	actors := NewActorCollection()
	otherActors := NewActorCollection()
	if actors.Compare(otherActors) != 0. {
		t.Fail()
	}
	actors.AddActorEntry("John Doe")
	otherActors.AddActorEntry("John Doe")
	if actors.Compare(otherActors) != 1. {
		t.FailNow()
	}
}

func TestActor(t *testing.T) {
	actors := NewActorCollection()
	if actors.Actor("Nemo") != nil {
		t.FailNow()
	}
}

func TestActorIndex(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	if actors.ActorIndexByName("John Doe") != 0 {
		t.Fail()
	}
	if actors.ActorIndexByName("Nemo") != -1 {
		t.FailNow()
	}
}

func TestActorRoles(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	if !reflect.DeepEqual(actors.Roles("John Doe"), []string{"performer"}) {
		t.Fail()
	}
	if actors.Roles("Nemo") != nil {
		t.FailNow()
	}
}

func TestAddActor(t *testing.T) {
	actors := NewActorCollection()
	actors.AddActorEntry("John Doe")
	actors.AddActorEntry("John Doe")
	if len(*actors) != 1 {
		t.FailNow()
	}
}

func TestAddRole(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	if (*actors)[0].Name != "John Doe" {
		t.Fail()
	}
	actor := actors.Actor("John Doe")
	if len(actor.Roles) != 1 && actor.Roles[0] != "performer" {
		t.FailNow()
	}
}

func TestDeleteRole(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	if err := actors.DeleteRole("John Doe", "performer"); err != nil {
		t.Fail()
	}
	actor := actors.Actor("John Doe")
	if len(actor.Roles) != 0 {
		t.Fail()
	}
	if err := actors.DeleteRole("Nemo", "performer"); err == nil {
		t.FailNow()
	}
}

func TestDeleteActor(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	actors.DeleteActor("John Doe")
	if len(*actors) != 0 {
		t.FailNow()
	}
}

func TestFilterActors(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	if len(*actors.Filter(IsPerformer)) != 1 {
		t.FailNow()
	}
}

func TestActorByID(t *testing.T) {
	actors := NewActorCollection()
	actor := actors.AddActorEntry("John Doe")
	actor.IDs = map[string]string{"discogs": "12345"}
	if actors.ActorByID("discogs", "12345") == nil {
		t.Fail()
	}
	if actors.ActorByID("musicbrainz", "12345") != nil {
		t.FailNow()
	}
}

func TestActorsCount(t *testing.T) {
	actors := NewActorCollection()
	actors.AddActorEntry("John Doe")
	if len(*actors) != 1 {
		t.FailNow()
	}
}

func TestActorKeys(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	air := actors.ActorByIndex(0)
	if !collection.ContainsStr("performer", air.Roles) {
		t.FailNow()
	}
}

func TestMerge(t *testing.T) {
	actors := NewActorCollection()
	actors.AddRole("John Doe", "performer")
	air := NewActorInRoles("John Doe")
	air.Roles = []string{"performer", "producer"}
	air.Notes = "some notes"
	air.IDs["discogs"] = "12345"
	actors.Merge(air)
	actor := actors.Actor("John Doe")
	if !reflect.DeepEqual(actor.Roles, []string{"performer", "producer"}) {
		t.Fail()
	}
	if !actor.IDs.Exists("discogs") {
		t.Fail()
	}
	if actor.Notes != "some notes" {
		t.Fail()
	}
	actors.Merge(NewActorInRoles("Nemo"))
	if actors.Actor("Nemo") == nil {
		t.FailNow()
	}
}
