package config

import "database/sql"
import "math/rand"
import "strings"

import "fmt"

import "screamingvortex/utilities"

type Perterbation struct {
  Id int64 `sql:"id"`

  FlagsString sql.NullString `sql:"flags"`
  MutedFlagsString sql.NullString `sql:"muted_flags"`
  RequiredFlagsString sql.NullString `sql:"required_flags"`

  flags []string
  mutedFlags []string
  requiredFlags []string

  Configs []*AssetConfig

  Manager *ConfigManager
  Rand *rand.Rand
}

func CreateEmptyPerterbation(client *utilities.Client, rRand *rand.Rand) *Perterbation {
  perterbation := new(Perterbation)
  if client != nil {
    perterbation.Manager = CreateEmptyManager(client)
  }
  perterbation.Rand = rRand
  return perterbation
}

func LoadPerterbation(manager *ConfigManager, perterbation *Perterbation) {
  if perterbation.FlagsString.String != "" {
    perterbation.flags = strings.Split(perterbation.FlagsString.String, ",")
  } else {
    perterbation.flags = make([]string, 0)
  }

  if perterbation.MutedFlagsString.String != "" {
    perterbation.mutedFlags = strings.Split(perterbation.MutedFlagsString.String, ",")
  } else {
    perterbation.mutedFlags = make([]string, 0)
  }

  if perterbation.RequiredFlagsString.String != "" {
    perterbation.requiredFlags = strings.Split(perterbation.RequiredFlagsString.String, ",")
  } else {
    perterbation.requiredFlags = make([]string, 0)
  }

  perterbation.Configs = FetchManyAssetConfigs(manager, perterbation.Id, perterbation.TableName(""), "configs")
}

func (perterbation *Perterbation) TableName(perterbationType string) string {
  return "plan_perterbation"
}

func (perterbation *Perterbation) GetId() *int64 {
  panic("GetId() not implemented. Config should not be editted.")
}

func (basePerterbation *Perterbation) AddInspiration(inspirationId int64) (*Inspiration, *Perterbation) {
  return basePerterbation.addInspiration(inspirationId)
}

func (basePerterbation *Perterbation) addInspiration(inspirationId int64) (*Inspiration, *Perterbation) {
  inspiration := basePerterbation.Manager.GetInspiration(inspirationId)
  newPerterbation := basePerterbation.Copy()
  for _, perterbationId := range inspiration.PerterbationIds {
    newPerterbation = newPerterbation.AddPerterbation(perterbationId)
  }

  return inspiration, newPerterbation
}

func (basePerterbation *Perterbation) Copy() *Perterbation {
  return basePerterbation.addPerterbation(CreateEmptyPerterbation(nil, nil), false)
}

func (basePerterbation *Perterbation) AddPerterbation(perterbationId int64) *Perterbation {
  modifyingPerterbation := basePerterbation.Manager.GetPerterbation(perterbationId)
  return basePerterbation.addPerterbation(modifyingPerterbation, false)
}

func (basePerterbation *Perterbation) addPerterbation(modifyingPerterbation *Perterbation, isSatellite bool) *Perterbation {
  if !basePerterbation.HasFlags(modifyingPerterbation.requiredFlags) {
    return basePerterbation.Copy()
  }

  newPerterbation := new(Perterbation)
  newPerterbation.Rand = basePerterbation.Rand
  newPerterbation.Manager = basePerterbation.Manager

  newPerterbation.flags = basePerterbation.CombineFlags(modifyingPerterbation)

  newPerterbation.Configs = StackAssetConfigs(basePerterbation.Configs, modifyingPerterbation.Configs)

  return newPerterbation
}

func (perterbation *Perterbation) GetConfig(typeId int64) *AssetConfig {
  for _, config := range perterbation.Configs {
    if config.TypeId == typeId {
      return config
    }
  }

  return CreateEmptyConfigAsset(typeId)
}

func FetchManyPerterbationIds(manager *ConfigManager, parentId int64, tableName string, valueName string) []int64 {
  ids := make([]int64, 0)
  examplePerterbation := new(Perterbation)
  manager.Client.FetchManyToManyChildIds(&ids, parentId, tableName, examplePerterbation.TableName(""), valueName, "", false)
  return ids
}

func (basePerterbation *Perterbation) CombineFlags(perterbation *Perterbation) []string {
  newFlags := make([]string, 0)
  for _, flagToAdd := range append(basePerterbation.flags, perterbation.flags...) {
    addFlag := true
    for _, mutedFlag := range append(basePerterbation.mutedFlags, perterbation.mutedFlags...) {
      if flagToAdd == mutedFlag {
        addFlag = false
        break
      }
    }

    if addFlag {
      newFlags = append(newFlags, flagToAdd)
    }
  }

  return newFlags
}

func (perterbation *Perterbation) HasFlags(requiredFlags []string) bool {
  for _, requiredFlag := range requiredFlags {
    hasFlag := false
    for _, activeFlag := range perterbation.flags {
      if activeFlag == requiredFlag {
        hasFlag = true
        break
      }
    }

    if !hasFlag {
      return false
    }
  }

  return true
}

func (perterbation *Perterbation) Print(indent int) {
  for i := 0; i < indent; i++ {
    fmt.Print(" ")
  }
  fmt.Print("PERTERBATION:\n")
  for i := 0; i < indent; i++ {
    fmt.Print(" ")
  }
  fmt.Printf("{Id:%d, flags:%+v}\n", perterbation.Id, perterbation.flags)

  for _, assetConfig := range perterbation.Configs {
    assetConfig.Print(indent+2)
  }
}
