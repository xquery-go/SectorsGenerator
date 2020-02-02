package config

import "database/sql"
import "math/rand"

import "github.com/kyleady/SectorsOfInequity/screamingvortex/utilities"

type Perterbation struct {
  SystemId sql.NullInt64 `sql:"system_id"`
  StarClusterId sql.NullInt64 `sql:"star_cluster_id"`
  SystemConfig *System
  StarClusterConfig *StarCluster
  Manager *ConfigManager
  Rand *rand.Rand
}

func CreateEmptyPerterbation(client *utilities.Client, rRand *rand.Rand) *Perterbation {
  perterbation := new(Perterbation)
  if client != nil {
    perterbation.Manager = CreateEmptyManager(client)
  }
  perterbation.Rand = rRand
  perterbation.SystemConfig = CreateEmptySystemConfig()
  perterbation.StarClusterConfig = CreateEmptyStarClusterConfig()
  return perterbation
}

func LoadPerterbationFrom(client utilities.ClientInterface, id int64) *Perterbation {
  perterbation := new(Perterbation)
  client.Fetch(perterbation, "", id)
  if perterbation.SystemId.Valid {
    perterbation.SystemConfig = LoadSystemConfigFrom(client, perterbation.SystemId.Int64)
  } else {
    perterbation.SystemConfig = CreateEmptySystemConfig()
  }

  if perterbation.StarClusterId.Valid {
    perterbation.StarClusterConfig = LoadStarClusterConfigFrom(client, perterbation.StarClusterId.Int64)
  } else {
    perterbation.StarClusterConfig = CreateEmptyStarClusterConfig()
  }

  return perterbation
}

func (perterbation *Perterbation) TableName(perterbationType string) string {
  return "plan_perterbation"
}

func (perterbation *Perterbation) GetId() *int64 {
  panic("GetId() not implemented. Config should not be editted.")
}

func (basePerterbation *Perterbation) AddInspiration(inspirationId int64) (*Inspiration, *Perterbation) {
  inspiration := basePerterbation.Manager.GetInspiration(inspirationId)
  var newPerterbation *Perterbation
  if(inspiration.PerterbationId.Valid) {
    newPerterbation = basePerterbation.AddPerterbation(inspiration.PerterbationId.Int64)
  } else {
    newPerterbation = basePerterbation.Copy()
  }

  return inspiration, newPerterbation
}

func (basePerterbation *Perterbation) Copy() *Perterbation {
  return basePerterbation.addPerterbation(CreateEmptyPerterbation(nil, nil))
}

func (basePerterbation *Perterbation) AddPerterbation(perterbationId int64) *Perterbation {
  modifyingPerterbation := basePerterbation.Manager.GetPerterbation(perterbationId)
  return basePerterbation.addPerterbation(modifyingPerterbation)
}

func (basePerterbation *Perterbation) addPerterbation(modifyingPerterbation *Perterbation) *Perterbation {
  newPerterbation := new(Perterbation)
  newPerterbation.Rand = basePerterbation.Rand
  newPerterbation.Manager = basePerterbation.Manager

  newPerterbation.SystemConfig = basePerterbation.SystemConfig.AddPerterbation(modifyingPerterbation.SystemConfig)
  newPerterbation.StarClusterConfig = basePerterbation.StarClusterConfig.AddPerterbation(modifyingPerterbation.StarClusterConfig)

  return newPerterbation
}
